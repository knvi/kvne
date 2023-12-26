package core

import "syscall"

type FDCom struct {
	Fd int
}

func (f FDCom) Read(data []byte) (int, error) {
	return syscall.Read(f.Fd, data)
}

func (f FDCom) Write(data []byte) (int, error) {
	return syscall.Write(f.Fd, data)
}