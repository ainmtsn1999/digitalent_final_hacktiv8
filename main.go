package main

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/app"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/config"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
