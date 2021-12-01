package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"project-management-demo-backend/pkg/util/environment"

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

// ReadConfig configures config file
func ReadConfig() {
	Config := &C

	if environment.IsDev() {
		viper.AddConfigPath(filepath.Join("./", "config"))
	} else if environment.IsTest() {
		viper.AddConfigPath(filepath.Join("./", "config.test"))
	} else {
		viper.AddConfigPath(filepath.Join("/srv", "config"))
	}

	viper.SetConfigType("yml")
	viper.SetConfigName("config")
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
