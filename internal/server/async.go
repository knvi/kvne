package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"syscall"
	"time"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core"
)

func readCommandFD(fd int) (*core.Command, error) {
	var buf []byte = make([]byte, 512)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		return nil, err
	}
	return core.ParseCommand(buf[:n])
}

func responseRw(cmd *core.Command, rw io.ReadWriter) {
	err := core.Exec(cmd, rw)
	if err != nil {
		responseErrorRw(err, rw)
	}
}

func responseErrorRw(err error, rw io.ReadWriter) {
	rw.Write([]byte(fmt.Sprintf("-%s%s", err, coder.CRLF)))
}

func RunAsyncTCPServer() error {
	log.Println("starting an asynchronous TCP server on", config.Host, config.Port)

	var events []syscall.EpollEvent = make([]syscall.EpollEvent, config.MaxConnection)
	client_number := 0

	serverFD, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	if err != nil {
		return err
	}
	defer syscall.Close(serverFD)

	if err = syscall.SetNonblock(serverFD, true); err != nil {
		return err
	}

	ip4 := net.ParseIP(config.Host)
	if err = syscall.Bind(serverFD, &syscall.SockaddrInet4{
		Port: config.Port,
		Addr: [4]byte{ip4[0], ip4[1], ip4[2], ip4[3]},
	}); err != nil {
		return err
	}

	if err = syscall.Listen(serverFD, config.MaxConnection); err != nil {
		return err
	}

	epollFD, err := syscall.EpollCreate1(0)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(epollFD)

	var socketServerReadReadyEvent syscall.EpollEvent = syscall.EpollEvent{
		Events: syscall.EPOLLIN,
		Fd:     int32(serverFD),
	}

	// listen to read events on the server
	if err = syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_ADD, serverFD, &socketServerReadReadyEvent); err != nil {
		return err
	}

	for {
		nevents, e := syscall.EpollWait(epollFD, events[:], -1)
		if e != nil {
			continue
		}

		for i := 0; i < nevents; i++ {
			if int(events[i].Fd) == serverFD {
				// accept the incoming connection from a client
				client_number++
				log.Printf("new client: id=%d\n", client_number)
				connFD, _, err := syscall.Accept(serverFD)
				if err != nil {
					log.Println("err", err)
					continue
				}

				if err = syscall.SetNonblock(connFD, true); err != nil {
					return err
				}

				// add the client socket FD to the epoll event loop
				var socketClientEvent syscall.EpollEvent = syscall.EpollEvent{
					Events: syscall.EPOLLIN,
					Fd:     int32(connFD),
				}
				if err := syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_ADD, connFD, &socketClientEvent); err != nil {
					log.Fatal(err)
				}
			} else {
				comm := core.FDCom{Fd: int(events[i].Fd)}
				cmd, err := readCommandFD(comm.Fd)
				if err != nil {
					syscall.Close(int(events[i].Fd))
					client_number--
					log.Println("client quit")
					continue
				}
				responseRw(cmd, comm)
			}
		}

		time.Sleep(1 * time.Millisecond)
	}
}
