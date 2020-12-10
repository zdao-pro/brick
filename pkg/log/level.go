package log

//Level log level
type Level int

const (
	_debugLevel Level = iota
	_infoLevel
	_warnLevel
	_accessLevel
	_errorLevel
	_fetalLevel
)

var levelName = [...]string{
	_debugLevel:  "DEBUG",
	_infoLevel:   "INFO",
	_warnLevel:   "WARN",
	_accessLevel: "ACCESS",
	_errorLevel:  "ERROR",
	_fetalLevel:  "FETAL",
}

func (l Level) String() string {
	return levelName[l]
}
