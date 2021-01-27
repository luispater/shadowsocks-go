package plugins

import (
	"fmt"
	ss "github.com/luispater/shadowsocks-go/shadowsocks"
	"io"
	"log"
	"net"
)

type closeWriter interface {
	CloseWrite() error
}

func RunForward(config *ss.Config) {
	listenAddr := fmt.Sprintf("%s:%d", config.LocalAddress, config.ForwordPort)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("starting forward server at %v ...\n", listenAddr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}
		go handleForwardConnection(config, conn)
	}
}

func handleForwardConnection(config *ss.Config, src net.Conn) {
	// defer src.Close()
	addr := net.JoinHostPort(config.Server.(string), fmt.Sprintf("%d", config.ServerPort))
	dst, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	errCh := make(chan error, 2)
	go doForward(dst, src, errCh)
	go doForward(src, dst, errCh)

	for i := 0; i < 2; i++ {
		e := <-errCh
		if e != nil {
			return
		}
	}
	return
}

func doForward(dst io.Writer, src io.Reader, errCh chan error) {
	_, err := io.Copy(dst, src)
	if tcpConn, ok := dst.(closeWriter); ok {
		tcpConn.CloseWrite()
	}
	errCh <- err
}
