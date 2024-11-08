package msg

const (
	TypeConnHead    = '~'
	TypeCryptoLogin = '!'
)

func init() {
	msgTypeMap[TypeConnHead] = ConnectHead{}
	msgTypeMap[TypeCryptoLogin] = CryptoLogin{}
	msgCtl.RegisterMsg(TypeConnHead, ConnectHead{})
	msgCtl.RegisterMsg(TypeCryptoLogin, CryptoLogin{})
}

type CryptoLogin struct {
	TimeStamp int64  `json:"time_stamp,omitempty"`
	Auth      string `json:"auth"`
	Sign      string `json:"sign,omitempty"`
}

type ConnectHead struct {
	ServiceName string `json:"service_name"`
}
