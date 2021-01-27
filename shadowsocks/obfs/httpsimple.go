package obfs

type HttpSimpleObfs struct {
	isServer      bool
	hasSentHeader bool
	hasRecvHeader bool
	host          string
	port          int
	recvBuffer    string
	userAgent     []string
}

func (this HttpSimpleObfs) Init(isServer bool) {
	this.isServer = isServer
	this.hasSentHeader = false
	this.hasRecvHeader = false
	this.host = ""
	this.port = 0
	this.recvBuffer = ""
	this.userAgent = []string{"Mozilla/5.0 (Windows NT 6.3; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (Windows NT 6.3; WOW64; rv:40.0) Gecko/20100101 Firefox/44.0",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.11 (KHTML, like Gecko) Ubuntu/11.10 Chromium/27.0.1453.93 Chrome/27.0.1453.93 Safari/537.36",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0",
		"Mozilla/5.0 (compatible; WOW64; MSIE 10.0; Windows NT 6.2)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/533.20.25 (KHTML, like Gecko) Version/5.0.4 Safari/533.20.27",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C)",
		"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Linux; Android 4.4; Nexus 5 Build/BuildID) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36",
		"Mozilla/5.0 (iPad; CPU OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
	}
}

func (this HttpSimpleObfs) InitData() []byte {
	return make([]byte, 0)
}

func (this HttpSimpleObfs) GetOverhead(direction bool) int64 {
	return 0
}

// 加密前
func (this HttpSimpleObfs) PreEncrypt(buf []byte) []byte {
	return buf
}

// 加密
func (this HttpSimpleObfs) Encode(buf []byte) []byte {
	return buf
}

// 解密
func (this HttpSimpleObfs) Decode(buf []byte) ([]byte, bool) {
	return buf, false
}

// 解密后
func (this HttpSimpleObfs) PostDecrypt(buf []byte) []byte {
	return buf
}

// UDP加密前
func (this HttpSimpleObfs) UdpPreEncrypt(buf []byte) []byte {
	return buf
}

// UDP解密后
func (this HttpSimpleObfs) UdpPostDecrypt(buf []byte) []byte {
	return buf
}

// 处理
func (this HttpSimpleObfs) Dispose() {

}
