package peach

//Client 配置客户端通用接口
type Client interface {
	Get(k string) *Value
}
