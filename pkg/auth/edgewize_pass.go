package auth

import "github.com/fatedier/frp/pkg/msg"

func (p *alwaysPass) VerifyCrypto(c *msg.CryptoLogin) msg.Login {
	return msg.Login{}
}
