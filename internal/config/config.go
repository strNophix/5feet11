package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	ScyllaDB struct {
		Hosts []string
	}
	MachineID int
}
