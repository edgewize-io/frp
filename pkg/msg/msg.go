// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import (
	"net"
	"reflect"
)

const (
	TypeLogin              = 0x01
	TypeLoginResp          = 0x02
	TypeNewProxy           = 0x03
	TypeNewProxyResp       = 0x04
	TypeCloseProxy         = 0x05
	TypeNewWorkConn        = 0x06
	TypeReqWorkConn        = 0x07
	TypeStartWorkConn      = 0x08
	TypeNewVisitorConn     = 0x09
	TypeNewVisitorConnResp = 0x0a
	TypePing               = 0x0b
	TypePong               = 0x0c
	TypeUDPPacket          = 0x0d
	TypeNatHoleVisitor     = 0x0e
	TypeNatHoleClient      = 0x10
	TypeNatHoleResp        = 0x11
	TypeNatHoleSid         = 0x12
	TypeNatHoleReport      = 0x0f
	TypeConnHead           = 0x13
	TypeCryptoLogin        = 0xf8
)

var msgTypeMap = map[byte]interface{}{
	TypeLogin:              Login{},
	TypeLoginResp:          LoginResp{},
	TypeNewProxy:           NewProxy{},
	TypeNewProxyResp:       NewProxyResp{},
	TypeCloseProxy:         CloseProxy{},
	TypeNewWorkConn:        NewWorkConn{},
	TypeReqWorkConn:        ReqWorkConn{},
	TypeStartWorkConn:      StartWorkConn{},
	TypeNewVisitorConn:     NewVisitorConn{},
	TypeNewVisitorConnResp: NewVisitorConnResp{},
	TypePing:               Ping{},
	TypePong:               Pong{},
	TypeUDPPacket:          UDPPacket{},
	TypeNatHoleVisitor:     NatHoleVisitor{},
	TypeNatHoleClient:      NatHoleClient{},
	TypeNatHoleResp:        NatHoleResp{},
	TypeNatHoleSid:         NatHoleSid{},
	TypeNatHoleReport:      NatHoleReport{},
	TypeConnHead:           ConnectHead{},
	TypeCryptoLogin:        CryptoLogin{},
}

var TypeNameNatHoleResp = reflect.TypeOf(&NatHoleResp{}).Elem().Name()

type ClientSpec struct {
	// Due to the support of VirtualClient, frps needs to know the client type in order to
	// differentiate the processing logic.
	// Optional values: ssh-tunnel
	Type string `json:"type,omitempty"`
	// If the value is true, the client will not require authentication.
	AlwaysAuthPass bool `json:"auth_type,omitempty"`
}

type CryptoLogin struct {
	TimeStamp int64  `json:"time_stamp,omitempty"`
	Auth      string `json:"auth"`
	Sign      string `json:"sign,omitempty"`
}

// When frpc start, client send this message to login to server.
type Login struct {
	Version      string            `json:"tag,omitempty"`
	Hostname     string            `json:"path,omitempty"`
	Os           string            `json:"system,omitempty"`
	Arch         string            `json:"types,omitempty"`
	User         string            `json:"people,omitempty"`
	PrivilegeKey string            `json:"identify,omitempty"`
	Timestamp    int64             `json:"time_second,omitempty"`
	RunID        string            `json:"self_id,omitempty"`
	Metas        map[string]string `json:"data_bucket,omitempty"`

	// Currently only effective for VirtualClient.
	ClientSpec ClientSpec `json:"other_data,omitempty"`

	// Some global configures.
	PoolCount int `json:"count,omitempty"`
}

type LoginResp struct {
	Version string `json:"tag,omitempty"`
	RunID   string `json:"self_id,omitempty"`
	Error   string `json:"error_msg,omitempty"`
}

// When frpc login success, send this message to frps for running a new proxy.
type NewProxy struct {
	ProxyName          string            `json:"proxy_name,omitempty"`
	ProxyType          string            `json:"proxy_type,omitempty"`
	UseEncryption      bool              `json:"use_encryption,omitempty"`
	UseCompression     bool              `json:"use_compression,omitempty"`
	BandwidthLimit     string            `json:"bandwidth_limit,omitempty"`
	BandwidthLimitMode string            `json:"bandwidth_limit_mode,omitempty"`
	Group              string            `json:"group,omitempty"`
	GroupKey           string            `json:"group_key,omitempty"`
	Metas              map[string]string `json:"metas,omitempty"`
	Annotations        map[string]string `json:"annotations,omitempty"`

	// tcp and udp only
	RemotePort int `json:"remote_port,omitempty"`

	// http and https only
	CustomDomains     []string          `json:"custom_domains,omitempty"`
	SubDomain         string            `json:"subdomain,omitempty"`
	Locations         []string          `json:"locations,omitempty"`
	HTTPUser          string            `json:"http_user,omitempty"`
	HTTPPwd           string            `json:"http_pwd,omitempty"`
	HostHeaderRewrite string            `json:"host_header_rewrite,omitempty"`
	Headers           map[string]string `json:"headers,omitempty"`
	ResponseHeaders   map[string]string `json:"response_headers,omitempty"`
	RouteByHTTPUser   string            `json:"route_by_http_user,omitempty"`

	// stcp, sudp, xtcp
	Sk         string   `json:"sk,omitempty"`
	AllowUsers []string `json:"allow_users,omitempty"`

	// tcpmux
	Multiplexer string `json:"multiplexer,omitempty"`
}

type NewProxyResp struct {
	ProxyName  string `json:"proxy_name,omitempty"`
	RemoteAddr string `json:"remote_addr,omitempty"`
	Error      string `json:"error,omitempty"`
}

type CloseProxy struct {
	ProxyName string `json:"proxy_name,omitempty"`
}

type NewWorkConn struct {
	RunID        string `json:"run_id,omitempty"`
	PrivilegeKey string `json:"privilege_key,omitempty"`
	Timestamp    int64  `json:"timestamp,omitempty"`
}

type ReqWorkConn struct{}

type StartWorkConn struct {
	ProxyName string `json:"proxy_name,omitempty"`
	SrcAddr   string `json:"src_addr,omitempty"`
	DstAddr   string `json:"dst_addr,omitempty"`
	SrcPort   uint16 `json:"src_port,omitempty"`
	DstPort   uint16 `json:"dst_port,omitempty"`
	Error     string `json:"error,omitempty"`
}

type NewVisitorConn struct {
	RunID          string `json:"dog_name,omitempty"`
	ProxyName      string `json:"cat_name,omitempty"`
	SignKey        string `json:"elephant_name,omitempty"`
	Timestamp      int64  `json:"duck_name,omitempty"`
	UseEncryption  bool   `json:"monkey_name,omitempty"`
	UseCompression bool   `json:"snake_name,omitempty"`
}

type NewVisitorConnResp struct {
	ProxyName string `json:"panda_name,omitempty"`
	Error     string `json:"something,omitempty"`
}

type Ping struct {
	PrivilegeKey string `json:"privilege_key,omitempty"`
	Timestamp    int64  `json:"timestamp,omitempty"`
}

type Pong struct {
	Error string `json:"error,omitempty"`
}

type UDPPacket struct {
	Content    string       `json:"c,omitempty"`
	LocalAddr  *net.UDPAddr `json:"l,omitempty"`
	RemoteAddr *net.UDPAddr `json:"r,omitempty"`
}

type NatHoleVisitor struct {
	TransactionID string   `json:"transaction_id,omitempty"`
	ProxyName     string   `json:"proxy_name,omitempty"`
	PreCheck      bool     `json:"pre_check,omitempty"`
	Protocol      string   `json:"protocol,omitempty"`
	SignKey       string   `json:"sign_key,omitempty"`
	Timestamp     int64    `json:"timestamp,omitempty"`
	MappedAddrs   []string `json:"mapped_addrs,omitempty"`
	AssistedAddrs []string `json:"assisted_addrs,omitempty"`
}

type NatHoleClient struct {
	TransactionID string   `json:"transaction_id,omitempty"`
	ProxyName     string   `json:"proxy_name,omitempty"`
	Sid           string   `json:"sid,omitempty"`
	MappedAddrs   []string `json:"mapped_addrs,omitempty"`
	AssistedAddrs []string `json:"assisted_addrs,omitempty"`
}

type PortsRange struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type NatHoleDetectBehavior struct {
	Role              string       `json:"role,omitempty"` // sender or receiver
	Mode              int          `json:"mode,omitempty"` // 0, 1, 2...
	TTL               int          `json:"ttl,omitempty"`
	SendDelayMs       int          `json:"send_delay_ms,omitempty"`
	ReadTimeoutMs     int          `json:"read_timeout,omitempty"`
	CandidatePorts    []PortsRange `json:"candidate_ports,omitempty"`
	SendRandomPorts   int          `json:"send_random_ports,omitempty"`
	ListenRandomPorts int          `json:"listen_random_ports,omitempty"`
}

type NatHoleResp struct {
	TransactionID  string                `json:"transaction_id,omitempty"`
	Sid            string                `json:"sid,omitempty"`
	Protocol       string                `json:"protocol,omitempty"`
	CandidateAddrs []string              `json:"candidate_addrs,omitempty"`
	AssistedAddrs  []string              `json:"assisted_addrs,omitempty"`
	DetectBehavior NatHoleDetectBehavior `json:"detect_behavior,omitempty"`
	Error          string                `json:"error,omitempty"`
}

type NatHoleSid struct {
	TransactionID string `json:"transaction_id,omitempty"`
	Sid           string `json:"sid,omitempty"`
	Response      bool   `json:"response,omitempty"`
	Nonce         string `json:"nonce,omitempty"`
}

type NatHoleReport struct {
	Sid     string `json:"sid,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type ConnectHead struct {
	ServiceName string `json:"serviceName"`
}
