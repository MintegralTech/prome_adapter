package model

import (
    "github.com/spf13/viper"
)

type Config struct {
    remoteUrl   string

    buffer int
    jobNames []string
    whitelist []string
}

var defaultConfigPath "./conf"
var defaultConfigName "adapter"
var defaultConfigType "yaml"
var Conf *Config

func init() {
    Conf = NewConfig()
}

func NewConfig() *Config {
    config := &Config{}
    v := setViper(defaultConfigPath, defaultConfigName, defaultConfigType)
    config.buffer = v.GetInt("runtime.buffer")
    config.remoteUrl = v.GetString("data.remote_url")
    config.jobNames = v.GeStringSlice("data.jobNames")
    config.whitelist = v.GeStringSlice("filter.whitelist")
}

func setViper(cfgPath, cfgName, cfgType string) *viper.Viper {
    v := viper.New()
    v.AddConfigPath(cfgPath)
    v.SetConfigName(cfgName)
    v.SetConfigType(cfgType)
    if err := v.ReadInConfig(); err != nil {
        panic(err)
    }
    return v
}

