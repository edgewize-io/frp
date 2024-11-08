package msg

const (
	TypeConnHead    = 's'
	TypeCryptoLogin = 't'
)

func init() {
	msgTypeMap[TypeConnHead] = ConnectHead{}
	msgTypeMap[TypeCryptoLogin] = CryptoLogin{}
}

type CryptoLogin struct {
	TimeStamp int64  `json:"time_stamp,omitempty"`
	Auth      string `json:"auth"`
	Sign      string `json:"sign,omitempty"`
}

type ConnectHead struct {
	ServiceName string `json:"serviceName"`
}
