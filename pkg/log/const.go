package log

const (
	_timeFormat = "2006-01-02T15:04:05.999999"

	// log level defined in level.go.
	_levelValue = "level_value"
	//  log level name: INFO, WARN...
	_level = "level"
	// log time.
	_time = "time"
	// request path.
	// _title = "title"
	// log file.
	_source = "source"
	// common log filed.
	_log = "log"
	// app name.
	_appID = "app_id"
	// container ID.
	_instanceID = "instance_id"
	// uniq ID from trace.
	_tid = "traceid"
	// request time.
	// _ts = "ts"
	// requester.
	_caller = "caller"
	// container environment: prod, pre, uat, fat.
	_deplyEnv = "env"
	// container area.
	_zone = "zone"
	// mirror flag
	_mirror = "mirror"
	// color.
	_color = "color"
	// env_color
	_envColor = "env_color"
	// cluster.
	_cluster = "cluster"
)

//D ...
type D Field
