package apollo

import (
	"brick/pkg/env"
	"brick/pkg/peach"
	"fmt"
	"os"

	"github.com/shima-park/agollo"
)

var (
	apolloConfig *Config
)

func init() {
	appid := env.GetAppID()
	configServerURL := os.Getenv("ConfigServerURL")
	nameSpace := os.Getenv("NameSpace")

	apolloConfig = &Config{
		configServerURL: configServerURL,
		appid:           appid,
		nameSpace:       nameSpace,
	}
	fmt.Println(apolloConfig)

	peach.RegistDriver("apollo", &appoloDriver{})
}

type appoloDriver struct {
}

func (ad *appoloDriver) New() (peach.Client, error) {
	apolloClientObj, err := newApolloClient(apolloConfig)
	if nil != err {
		return nil, fmt.Errorf("paladin: unknown appoloDriver (forgotten register?)")
	}

	return apolloClientObj, nil
}

//Config apollo config struct
type Config struct {
	configServerURL string
	appid           string
	nameSpace       string
}

//newApolloClient 返回一个apollo配置中心
func newApolloClient(c *Config) (peach.Client, error) {

	a, err := agollo.New(c.configServerURL, c.appid, agollo.AutoFetchOnCacheMiss())
	if err != nil {
		return nil, err
	}

	errorCh := a.Start()

	if errorCh != nil {
		fmt.Println("Start failed.......")
	}

	if nil != err {
		panic(err)
	}
	apoloClientInstance := &ApoClient{}
	apoloClientInstance.SetApollo(a)

	return apoloClientInstance, nil
}

//ApoClient 实例化的apollo配置中心
type ApoClient struct {
	a agollo.Agollo
}

//SetApollo 设置apollo客户端
func (ac *ApoClient) SetApollo(ap agollo.Agollo) {
	ac.a = ap
}

//Get 获取NameSpace下key的值
func (ac *ApoClient) Get(k string) *peach.Value {
	v := ac.a.Get(k, agollo.WithNamespace(apolloConfig.nameSpace))
	return peach.NewValue(v, v)
}

//GetKey 获取NameSpace下key的值
func (ac *ApoClient) GetKey(key string) string {
	k := ac.a.Get(key, agollo.WithNamespace(apolloConfig.nameSpace))
	fmt.Println(k)
	return k
}
