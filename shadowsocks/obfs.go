package shadowsocks

import (
	. "github.com/luispater/shadowsocks-go/shadowsocks/obfs"
)

type Obfs interface {
	Init(isServer bool)
	InitData() []byte
	GetOverhead(direction bool) int64

	PreEncrypt(buf []byte) []byte
	Encode(buf []byte) []byte
	Decode(buf []byte) ([]byte, bool)
	PostDecrypt(buf []byte) []byte

	UdpPreEncrypt(buf []byte) []byte
	UdpPostDecrypt(buf []byte) []byte

	Dispose()
}

func NewObfs(method string) Obfs {
	switch method {
	case "Plain":
		return &PlainObfs{}
	case "HttpSimple":
		return &HttpSimpleObfs{}
	default:
		return &PlainObfs{}
	}
}
