package auth

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/util/util"
	"strconv"
)

func (auth *OidcAuthProvider) SetCrypto(login *msg.Login) msg.CryptoLogin {
	var c msg.CryptoLogin
	bs, _ := json.Marshal(login)
	c.Auth = util.RandomString(16)
	key := util.GenerateAesKey(c.Auth)[:8]
	iv := util.GenerateAesKey(strconv.FormatInt(c.TimeStamp, 16))[:8]
	c.Sign = hex.EncodeToString(util.Encrypt(string(bs), key, []byte(iv)))
	return c
}

func (auth *OidcAuthConsumer) VerifyCrypto(c *msg.CryptoLogin) msg.Login {
	var login msg.Login
	bs, err := hex.DecodeString(c.Sign)
	if err != nil {
		fmt.Println("verify crypto error,", err)
		return msg.Login{}
	}
	key := util.GenerateAesKey(c.Auth)[:8]
	iv := util.GenerateAesKey(strconv.FormatInt(c.TimeStamp, 16))[:8]
	before := util.Decrypt(bs, key, []byte(iv))
	_ = json.Unmarshal(before, &login)
	return login
}
