package configs

import (
	"fmt"
	"log"

	"github.com/Ctere1/golang-api/pkgs/api"
	"github.com/Ctere1/golang-api/pkgs/storage"
	"github.com/spf13/viper"
)

var Cfg Configurations

type Configurations struct {
	Api      apiconf
	Database databaseconf
}

type apiconf struct {
	ListenAddress string
	ListenPort    string
	BearerToken   string
}

type databaseconf struct {
	Server,
	Port,
	Username,
	Password,
	Dbname,
	Sslmode,
	Timezone string
	MaxOpenConns,
	MaxIdleConns,
	ConnMaxLifetimeInMinutes int
}

func (c *Configurations) Load() {

	viper.SetConfigType("toml")

	viper.SetConfigName("app.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Println(err)
	}
}

func LoadConfigs() {
	//load configs from file
	Cfg.Load()

	//database configuraion
	storage.Dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", Cfg.Database.Server, Cfg.Database.Username, Cfg.Database.Password, Cfg.Database.Dbname, Cfg.Database.Port, Cfg.Database.Sslmode, Cfg.Database.Timezone)
	storage.MaxOpenConns = Cfg.Database.MaxOpenConns
	storage.MaxIdleConns = Cfg.Database.MaxIdleConns
	storage.ConnMaxLifetimeInMinutes = Cfg.Database.ConnMaxLifetimeInMinutes

	// API configuration
	api.ListenAddress = Cfg.Api.ListenAddress
	api.ListenPort = ":" + Cfg.Api.ListenPort
	api.BearerToken = Cfg.Api.BearerToken

}
