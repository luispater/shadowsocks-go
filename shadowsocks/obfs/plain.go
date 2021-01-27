package obfs

type PlainObfs struct {
	isServer bool
}

func (this PlainObfs) Init(isServer bool) {
	this.isServer = isServer
}

func (this PlainObfs) InitData() []byte {
	return make([]byte, 0)
}

func (this PlainObfs) GetOverhead(direction bool) int64 {
	return 0
}

// 加密前
func (this PlainObfs) PreEncrypt(buf []byte) []byte {
	return buf
}

// 加密
func (this PlainObfs) Encode(buf []byte) []byte {
	return buf
}

// 解密
func (this PlainObfs) Decode(buf []byte) ([]byte, bool) {
	return buf, false
}

// 解密后
func (this PlainObfs) PostDecrypt(buf []byte) []byte {
	return buf
}

// UDP加密前
func (this PlainObfs) UdpPreEncrypt(buf []byte) []byte {
	return buf
}

// UDP解密后
func (this PlainObfs) UdpPostDecrypt(buf []byte) []byte {
	return buf
}

// 处理
func (this PlainObfs) Dispose() {

}
