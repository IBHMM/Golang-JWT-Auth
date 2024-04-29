package config

import "github.com/joho/godotenv"

type Variables struct {
	Port        int
	Db_name     string
	Db_user     string
	Db_password string
	Db_port     int
	Secret      string
}

type Load interface {
	Load() error
}

func (variable *Variables) Load() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	} else {
		return nil
	}
}
