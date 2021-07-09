package main

import (
	"fmt"

	"github.com/danx-dev/kidspointsystem/pkg/db"
	"github.com/danx-dev/kidspointsystem/pkg/server"
)

func main() {
	fmt.Println("Program running!")
	db.AddKid(db.Kid{Id: 1, Name: "Ella", Points: 0})
	db.AddKid(db.Kid{Id: 2, Name: "Frida", Points: 0})
	db.AddKid(db.Kid{Id: 3, Name: "Karla", Points: 0})
	db.GetAllKids()
	server.NewServer()

}
