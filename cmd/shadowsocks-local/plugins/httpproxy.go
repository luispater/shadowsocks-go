package plugins

import (
	"bytes"
	"github.com/luispater/goproxy"
	"io/ioutil"
	"net/http"
)

var unauthorizedMsg = []byte("407 Proxy Authentication Required")

func HttpProxyAuth(realm string, f func(authType string, authInfo ...string) bool) goproxy.ReqHandler {
	return goproxy.FuncReqHandler(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		status, response, _ := httpProxyAuthBase(req, f)
		if !status {
			if response == nil {
				return nil, Unauthorized(req, realm)
			} else {
				return nil, response
			}
		}
		return req, nil
	})
}

func HttpProxyConnectAuth(realm string, f func(authType string, authInfo ...string) bool) goproxy.HttpsHandler {
	return goproxy.FuncHttpsHandler(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		status, response, action := httpProxyAuthBase(ctx.Req, f)
		if !status {
			if response == nil {
				ctx.Resp = Unauthorized(ctx.Req, realm)
				return goproxy.WaitForAuthConnect, host
			} else {
				ctx.Resp = response
				return action, host
			}
		}
		return goproxy.OkConnect, host
	})
}

func httpProxyAuthBase(_ *http.Request, _ func(authType string, authInfo ...string) bool) (bool, *http.Response, *goproxy.ConnectAction) {
	return true, nil, nil
}

// func httpProxyAuthBase(req *http.Request, f func(authType string, authInfo ...string) bool) (bool, *http.Response, *goproxy.ConnectAction) {
// 	authheader := strings.SplitN(req.Header.Get(proxyAuthorizationHeader), " ", 2)
// 	req.Header.Del(proxyAuthorizationHeader)
// 	if len(authheader) != 2 {
// 		return false, nil, goproxy.RejectConnect
// 	}
// 	if authheader[0] == "Basic" {
// 		rawUserPassword, err := base64.StdEncoding.DecodeString(authheader[1])
// 		if err != nil {
// 			return false, nil, goproxy.RejectConnect
// 		}
// 		userPassword := strings.SplitN(string(rawUserPassword), ":", 2)
// 		if len(userPassword) != 2 {
// 			return false, nil, goproxy.RejectConnect
// 		}
// 		return f(authheader[0], userPassword[0], userPassword[1]), nil, goproxy.RejectConnect
// 	} else if authheader[0] == "NTLM" {
// 		rawProxyAuthPayload, err := base64.StdEncoding.DecodeString(authheader[1])
// 		if err != nil {
// 			return false, nil, goproxy.RejectConnect
// 		}
// 		ntlmMessageType := binary.LittleEndian.Uint32(rawProxyAuthPayload[8:12])
// 		if ntlmMessageType == 1 {
// 			session, err := ntlm.CreateServerSession(ntlm.Version2, ntlm.ConnectionOrientedMode)
// 			if err != nil {
// 				return false, nil, goproxy.RejectConnect
// 			}
// 			challenge, err := session.GenerateChallengeMessage()
// 			if err != nil {
// 				return false, nil, goproxy.RejectConnect
// 			}
// 			Challenges.Store(req.RemoteAddr, challenge)
// 			proxyAuthPayload := base64.StdEncoding.EncodeToString(challenge.Bytes())
//
// 			response := &http.Response{
// 				StatusCode:    407,
// 				ProtoMajor:    1,
// 				ProtoMinor:    1,
// 				Request:       req,
// 				Header:        http.Header{"Proxy-Authenticate": []string{"NTLM " + proxyAuthPayload}},
// 				Body:          ioutil.NopCloser(bytes.NewBuffer(unauthorizedMsg)),
// 				ContentLength: int64(len(unauthorizedMsg)),
// 			}
// 			return false, response, goproxy.WaitForAuthConnect
// 		} else if ntlmMessageType == 3 {
// 			interfaceChallenge, ok := Challenges.Load(req.RemoteAddr)
// 			if !ok || interfaceChallenge == nil {
// 				response := &http.Response{
// 					StatusCode:    407,
// 					ProtoMajor:    1,
// 					ProtoMinor:    1,
// 					Request:       req,
// 					Header:        http.Header{"Proxy-Authenticate": []string{"NTLM"}},
// 					Body:          ioutil.NopCloser(bytes.NewBuffer(unauthorizedMsg)),
// 					ContentLength: int64(len(unauthorizedMsg)),
// 				}
// 				return false, response, goproxy.RejectConnect
// 			}
// 			challenge := interfaceChallenge.(*ntlm.ChallengeMessage)
// 			msg, err := ntlm.ParseAuthenticateMessage(rawProxyAuthPayload, 2)
// 			if err == nil {
// 				NtlmServerSessionV2.SetServerChallenge(challenge.ServerChallenge)
// 				err = NtlmServerSessionV2.ProcessAuthenticateMessage(msg)
// 				if err != nil {
// 					response := &http.Response{
// 						StatusCode:    407,
// 						ProtoMajor:    1,
// 						ProtoMinor:    1,
// 						Request:       req,
// 						Header:        http.Header{"Proxy-Authenticate": []string{"NTLM"}},
// 						Body:          ioutil.NopCloser(bytes.NewBuffer(unauthorizedMsg)),
// 						ContentLength: int64(len(unauthorizedMsg)),
// 					}
// 					return false, response, goproxy.RejectConnect
// 				} else {
// 					return true, nil, nil
// 				}
// 			} else {
// 				msg, err = ntlm.ParseAuthenticateMessage(rawProxyAuthPayload, 1)
// 				if err == nil {
// 					NtlmServerSessionV1.SetServerChallenge(challenge.ServerChallenge)
// 					err = NtlmServerSessionV1.ProcessAuthenticateMessage(msg)
// 				}
// 				if err != nil {
// 					response := &http.Response{
// 						StatusCode:    407,
// 						ProtoMajor:    1,
// 						ProtoMinor:    1,
// 						Request:       req,
// 						Header:        http.Header{"Proxy-Authenticate": []string{"NTLM"}},
// 						Body:          ioutil.NopCloser(bytes.NewBuffer(unauthorizedMsg)),
// 						ContentLength: int64(len(unauthorizedMsg)),
// 					}
// 					return false, response, goproxy.RejectConnect
// 				} else {
// 					return true, nil, nil
// 				}
// 			}
// 		}
// 	} else {
//
// 	}
// 	return false, nil, goproxy.RejectConnect
// }

func Unauthorized(req *http.Request, realm string) *http.Response {
	return &http.Response{
		StatusCode: 407,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Request:    req,
		Header:     http.Header{"Proxy-Authenticate": []string{"Basic realm=" + realm, "NTLM"}},
		// Header:        http.Header{"Proxy-Authenticate": []string{"NTLM"}},
		Body:          ioutil.NopCloser(bytes.NewBuffer(unauthorizedMsg)),
		ContentLength: int64(len(unauthorizedMsg)),
	}
}
