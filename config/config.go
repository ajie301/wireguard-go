package config

// 隧道协议
const (
	PROTO_WIREGUARD  = iota // 原生wireguard协议
	PROTO_DEEPTUN_V1        // 混淆协议
)

var c = &conf{
	protocol:     PROTO_WIREGUARD,
	netinfcreate: true,
}

func SetProtocol(proto int) {
	c.protocol = proto
}

func SetNetInf(create bool) {
	c.netinfcreate = create
}

func Protocol() int {
	return c.protocol
}

func NetInf() bool {
	return c.netinfcreate
}

// 全局配置数据
type conf struct {
	protocol     int  // 隧道数据协议类型
	netinfcreate bool // 是否创建网卡
}
