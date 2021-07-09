package db

import (
	"fmt"
	"time"
)

type Kid struct {
	Id        int64
	Name      string
	Points    int64
	CreatedAt string `json:"createdAt" db:"createdAt"`
}

func AddKid(kid Kid) {
	db := GetDatabase()
	sql := `INSERT INTO kid (id,name,points,createdAt) VALUES(?,?,?,?)`
	db.Exec(sql, kid.Id, kid.Name, kid.Points, time.Now().String())
	fmt.Println("Kid Added")
}

func GetAllKids() []Kid {
	db := GetDatabase()
	kids := []Kid{}
	err := db.Select(&kids, "SELECT * FROM kid")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Println(kids)
	return kids
}

func FindKidInArray(kidId int64, kids []Kid) Kid {
	var kid Kid
	for _, v := range kids {
		if v.Id == kidId {
			kid = v
		}
	}
	return kid
}
