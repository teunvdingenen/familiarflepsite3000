package config

import "time"

type GeneralConfig struct {
	DB   DBConfig
	Auth AUTHConfig
}

type DBConfig struct {
	DatabaseHost string
	DatabaseName string
}

type AUTHConfig struct {
	Issuer        string
	SigningKey    []byte
	ValidDuration time.Duration
}

var Config = GeneralConfig{
	DB: DBConfig{
		DatabaseHost: "mongodb://localhost:27017",
		DatabaseName: "familiarflepsite3000",
	},
	Auth: AUTHConfig{
		Issuer:        "familiarflepsite",
		SigningKey:    []byte("somekey"),
		ValidDuration: time.Hour * 1,
	},
}
