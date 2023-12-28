package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core"
)

func readCommand(conn net.Conn) (*core.Command, error) {
	var buf []byte = make([]byte, 512);

	n, err := conn.Read(buf);
	if err != nil {
		return nil, err
	}

	return core.ParseCommand(buf[:n]);
}

func respondError(err error, conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf("-%s%s", err.Error(), coder.CRLF)));
}

func respond(cmd *core.Command, conn net.Conn)  {
	err := core.Exec(cmd, conn);
	if err != nil {
		respondError(err, conn);
	}
}

func StartServer(wg *sync.WaitGroup) {
	defer wg.Done();
	log.Println("Starting server on ", fmt.Sprintf("%s:%d", config.Host, config.Port));

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port));
	if err != nil {
		log.Fatal(err);
	}

	for {
		conn, err := listener.Accept();
		if err != nil {
			log.Fatal(err);
		}

		log.Println("New connection from ", conn.RemoteAddr().String());

		for {
			cmd, err := readCommand(conn);
			if err != nil {
				conn.Close();
				if err == io.EOF {
					log.Println("Connection closed by client");
					break;
				}
				log.Fatal(err);
			}

			log.Println("Command received: ", cmd);
			respond(cmd, conn);
		}	
	}
}