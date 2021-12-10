package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"project-management-demo-backend/pkg/util/environment"
	"runtime"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Addr                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
			Charset   string
			Loc       string
		}
	}
	Server struct {
		Address string
	}
	Firebase struct {
		ServiceKey string
	}
}

// C is config variable
var C config

// ReadConfigOption is a config option
type ReadConfigOption struct {
	AppEnv string
}

// ReadConfig configures config file
func ReadConfig(option ReadConfigOption) {
	Config := &C

	if environment.IsDev() || os.Getenv("APP_ENV") == "" {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config")
	} else if environment.IsTest() || (option.AppEnv == environment.Test) {
		fmt.Println(rootDir())
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.test")
	} else if environment.IsE2E() || (option.AppEnv == environment.E2E) {
		fmt.Println(rootDir())
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.e2e")
	} else {
		viper.AddConfigPath(filepath.Join("/srv", "config"))
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//spew.Dump(C)
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
