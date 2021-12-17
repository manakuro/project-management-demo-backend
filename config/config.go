package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	AppEnv   string
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

// Application Environment name
const (
	Development = "development"
	Test        = "test"
	E2E         = "e2e"
)

// ReadConfigOption is a config option
type ReadConfigOption struct {
	AppEnv string
}

// ReadConfig configures config file
func ReadConfig(option ReadConfigOption) {
	Config := &C

	if option.AppEnv == Test {
		setTest()
	} else if option.AppEnv == E2E {
		setE2E()
	} else if option.AppEnv == Development || option.AppEnv == "" {
		setDev()
	} else {
		setProd()
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

	if Config.AppEnv == Development {
		spew.Dump(C)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func setDev() {
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.SetConfigName("config")
}

func setTest() {
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.SetConfigName("config.test")
}

func setE2E() {
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.SetConfigName("config.e2e")
}

func setProd() {
	viper.AddConfigPath(filepath.Join("/srv", "config"))
	viper.SetConfigName("config")
}
