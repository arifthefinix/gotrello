package main

import (
	"gotrello/Database"
	"gotrello/Routes"
)

func main() {
	if err := Database.Migrate(); err != nil {
		panic(err)
	}
	router := Routes.SetupRouter()
	router.Run(":8031")
}