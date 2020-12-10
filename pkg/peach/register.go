package peach

import (
	"fmt"
)

var (
	drivers = make(map[string]Driver)
)

//RegistDriver regist driver
func RegistDriver(name string, driver Driver) {
	if nil == driver {
		panic("driver is nil")
	}
	drivers[name] = driver
}

//GetDriver return Driver by DriverName
func GetDriver(name string) (Driver, error) {
	driver, ok := drivers[name]

	if !ok {
		return nil, fmt.Errorf("paladin: unknown driver %q (forgotten register?)", name)
	}

	return driver, nil
}
