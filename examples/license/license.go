package main

import (
	"log"

	"github.com/nenavizhuleto/golic"
)

type License struct {
	Param1 any `json:"param1"`
	Param2 any `json:"param2"`
}

func main() {
	service, err := golic.Open("driverName", "DSN")
	if err != nil {
		panic(err)
	}

	verified, err := service.Verify()
	if err != nil {
		panic(err)
	}

	if !verified {
		panic("license not verified")
	}

	var license License
	if err := service.Scan(&license); err != nil {
		panic(err)
	}

	log.Println(license)
}
