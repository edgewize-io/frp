package auth

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/util/util"
	"strconv"
)

func (auth *TokenAuthSetterVerifier) SetCrypto(c *msg.CryptoLogin, login *msg.Login) {
	bs, _ := json.Marshal(login)
	key := util.GenerateAesKey(auth.token)[:8]
	iv := util.GenerateAesKey(strconv.FormatInt(c.TimeStamp, 16))[:8]
	c.Auth = util.RandomString(16)
	c.Sign = hex.EncodeToString(util.Encrypt(string(bs), key, []byte(iv)))
}

func (auth *TokenAuthSetterVerifier) VerifyCrypto(c *msg.CryptoLogin) *msg.Login {
	var login msg.Login
	bs, err := hex.DecodeString(c.Sign)
	if err != nil {
		fmt.Println("verify crypto error,", err)
		return nil
	}
	key := util.GenerateAesKey(auth.token)[:8]
	iv := util.GenerateAesKey(strconv.FormatInt(c.TimeStamp, 16))[:8]
	before := util.Decrypt(bs, key, []byte(iv))
	_ = json.Unmarshal(before, &login)
	return &login
}
