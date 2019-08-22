package setting

import (
	"github.com/BurntSushi/toml"
	"log"
	"rollcat/pkg/constants"
	"sync"
	"time"
)

var (
	once sync.Once

	conf         *Conf
	GinLogPath   string
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
	App      `toml:"app"`
	Server   `toml:"server"`
	Database `toml:"database"`
}

type App struct {
	GinLogPath string `toml:"ginLogPath"`
	SecretKey  string `toml:"secretKey"`
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
		if _, err := toml.DecodeFile(constants.FPath, &conf); err != nil {
			log.Fatalf("Failed to parse 'conf/conf.toml': %v ", err)
		}

		// App
		GinLogPath = conf.GinLogPath

		// Server
		if conf.RunMode == "" {
			RunMode = constants.DebugMode
		} else {
			RunMode = conf.RunMode
		}
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
