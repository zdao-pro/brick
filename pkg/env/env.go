package env

import (
	"os"
)

// env configuration.
const (
	DeployEnvDev    = "dev"
	DeployEnvPre    = "pre"
	DeployEnvOnline = "online"
)

var (
	// Hostname machine hostname.
	Hostname string
	// AppID is global unique application id, register by service tree.
	// such as main.arch.disocvery.
	AppID string
	// DeployEnv deploy env where app at.
	DeployEnv string
)

func init() {
	Hostname = os.Getenv("HOSTNAME")
	AppID = os.Getenv("APPID")
	DeployEnv = os.Getenv("DEPLOYENV")
}

//GetAppID 获取应用ID
func GetAppID() string {
	return AppID
}

//GetEnv 获取应用环境变量
func GetEnv() string {
	return DeployEnv
}

//GetHostname 获取应用主机名
func GetHostname() string {
	return Hostname
}

//IsOnline 是否为生产环境
func IsOnline() bool {
	if DeployEnv == DeployEnvOnline {
		return true
	}
	return false
}

//IsDev 是否为开发环境
func IsDev() bool {
	if DeployEnv == DeployEnvPre {
		return true
	}
	return false
}
