package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/luispater/goproxy"
	"github.com/luispater/goproxy/ext/auth"
	ss "github.com/luispater/shadowsocks-go/shadowsocks"
	"log"
	"net"
	"net/http"
	"os"
)

func shadowSocksDial(_, addr string) (net.Conn, error) {
	cipher, err := ss.NewCipher("aes-128-cfb", "eceasy")
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return nil, err
	}
	serverAddr := net.JoinHostPort("2001:470:18:301::2", "8988")
	rawAddr, err := ss.RawAddr(addr)
	if err != nil {
		return nil, err
	}
	var obfs *ss.Obfs
	conn, err := ss.DialWithRawAddr(rawAddr, serverAddr, cipher.Copy(), obfs)
	if err != nil {
		return nil, errors.New("connect: connection refused")
	}
	return conn, nil
}

func shadowSocksDialContext(_ context.Context, network, addr string) (net.Conn, error) {
	return shadowSocksDial(network, addr)
}

func main() {
	proxy := goproxy.ProxyHttpServer{
		Logger: log.New(os.Stderr, "", log.LstdFlags),
		NonproxyHandler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			http.Error(w, "This is a proxy server. Does not respond to non-proxy requests.", 500)
		}),
	}
	proxy.Tr = &http.Transport{
		DialContext: shadowSocksDialContext,
	}
	proxy.ConnectDial = shadowSocksDial

	proxy.OnRequest().Do(auth.Basic("my_realm", func(user, passwd string) bool {
		return user == "luis" && passwd == "831227"
	}))
	proxy.OnRequest().HandleConnect(auth.BasicConnect("my_realm", func(user, passwd string) bool {
		return user == "luis" && passwd == "831227"
	}))

	log.Fatal(http.ListenAndServe(":8080", &proxy))
}
