package main

import (
	"api/internal/api"
	"api/internal/config"
    "github.com/spf13/viper"
	log "github.com/sirupsen/logrus" 
)


func main(){
	log.SetFormatter(&log.JSONFormatter{})

	var appCfg config.Config
	var pgConf config.PGStoreConfg

    viper.AddConfigPath(".")
    viper.SetConfigName(".env")
    viper.SetConfigType("env")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Error("Error reading config file, %s", err)
		return
    }
	if err := viper.Unmarshal(&appCfg); err != nil {
        log.Error("Error Unmarshal AppConf, %s", err)
		return
	}

	if err := viper.Unmarshal(&pgConf); err != nil {
        log.Error("Error Unmarshal DBConf, %s", err)
		return
	}
	appCfg.PgConf = &pgConf
	api := api.New(&appCfg)

	if err := api.Run(); err != nil {
        log.Error("Error running app, %s", err)	
	}
}