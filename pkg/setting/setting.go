package setting

import (
	"github.com/BurntSushi/toml"
	"log"
	"sync"
	"time"
)

var (
	once         sync.Once
	conf         *Conf
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Type         string
	Host         string
	Name         string
	User         string
	Password     string
)

type Conf struct {
	Server   `toml:"server"`
	Database `toml:"database"`
}

type Server struct {
	RunMode      string        `toml:"runMode"`
	HttpPort     int           `toml:"httpPort"`
	ReadTimeout  time.Duration `toml:"readTimeout"`
	WriteTimeout time.Duration `toml:"writeTimeout"`
}

type Database struct {
	Type     string `toml:"type"`
	Host     string `toml:"host"`
	Name     string `toml:"name"`
	Password string `toml:"password"`
	User     string `toml:"user"`
}

func init() {
	once.Do(func() {
		if _, err := toml.DecodeFile("conf/conf.toml", &conf); err != nil {
			log.Fatalf("Failed to parse 'conf/conf.toml': %v ", err)
		}

		// Server
		RunMode = conf.RunMode
		HttpPort = conf.HttpPort
		ReadTimeout = conf.ReadTimeout * time.Second
		WriteTimeout = conf.WriteTimeout * time.Second

		// Database
		Type = conf.Type
		Host = conf.Host
		Name = conf.Name
		User = conf.User
		Password = conf.Password
	})
}
