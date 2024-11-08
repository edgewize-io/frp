package auth

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/util/util"
	"strconv"
)

func (auth *OidcAuthProvider) SetCrypto(c *msg.CryptoLogin, login *msg.Login) {
	bs, _ := json.Marshal(login)
	c.Auth = util.RandomString(16)
	key := util.GenerateAesKey(c.Auth)[:8]
	iv := util.GenerateAesKey(strconv.FormatInt(c.TimeStamp, 16))[:8]
	c.Sign = hex.EncodeToString(util.Encrypt(string(bs), key, []byte(iv)))
}

func (auth *OidcAuthConsumer) VerifyCrypto(c *msg.CryptoLogin) *msg.Login {
	var login *msg.Login
	bs, err := hex.DecodeString(c.Sign)
	if err != nil {
		fmt.Println("verify crypto error,", err)
		return nil
	}
	key := util.GenerateAesKey(c.Auth)[:8]
	iv := util.GenerateAesKey(strconv.FormatInt(c.TimeStamp, 16))[:8]
	before := util.Decrypt(bs, key, []byte(iv))
	_ = json.Unmarshal(before, login)
	return login
}
