package db

type Parrent struct {
	Id        int64
	Name      string
	CreatedAt string
}

func AddParrent(parrent Parrent) {
	// db, _ := db.CreateDatabase()
	// statement, _ :=
	// 	db.Prepare("INSERT INTO Parrent (id,name,createdAt) VALUES(?, ?, ?)")
	// t := time.Now()
	// statement.Exec(parrent.Id, parrent.Name, t.String())
}
