package main

import (
	"GoWebProject/internal/config"
)

func main() {
	viper := config.NewViper()
	server := config.NewGin()
	log, err := config.NewZapLogger(viper)
	if err != nil {
		panic(err)
	}
	orm := config.NewGorm(viper, log)
	validator := config.NewValidator()

	app := config.AppConfig{
		DB:        orm,
		Server:    server,
		Config:    viper,
		Logger:    log,
		Validator: validator,
	}

	config.App(&app)

	if err := app.Server.Run(); err != nil {
		log.Error(err.Error())
		panic(err)
	}
}
