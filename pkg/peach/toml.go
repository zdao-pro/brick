package peach

import (
	"github.com/BurntSushi/toml"
)

//TomlClient 解析toml文件的客户端
type TomlClient struct {
	mapping map[string]interface{}
}

//NewToml 新建toml
func (t *TomlClient) NewToml(text []byte) (*TomlClient, error) {
	raws := map[string]interface{}{}
	if err := toml.Unmarshal(text, &raws); err != nil {
		return nil, err
	}

	return &TomlClient{}, nil
}

//Get 获取key
func (t *TomlClient) Get(k string) *Value {
	return t.mapping[k].(*Value)
}
