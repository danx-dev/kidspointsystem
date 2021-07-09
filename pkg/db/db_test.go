package db

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	DeleteDatabase()
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	DeleteDatabase()
	os.Exit(exitVal)
}

func TestAddKid(t *testing.T) {
	AddKid(Kid{Id: 99, Name: "Testing", Points: 0})
	kids := GetAllKids()
	firstKid := kids[0]
	fmt.Println(firstKid)
}

func TestGetAllOpenAssignments(t *testing.T) {
	result := GetAllOpenAssignments()
	fmt.Println(result)

}

func DeleteDatabase() {
	if _, err := os.Stat("./pointsystem.db"); err == nil {
		err := os.Remove("./pointsystem.db")
		if err != nil {
			fmt.Println(err)
		}
	}
}
