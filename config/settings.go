package config

import "os"

type SettingsStruct struct {
	Port string
	Host string
}

func InitSettings() *SettingsStruct {
	return &SettingsStruct{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}
}
