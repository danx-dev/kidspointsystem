package db

type Assignment struct {
	Id          int64
	Description string
	Points      int64
	Status      string
	CreatedAt   string
	KidId       int64
}

func GetAllOpenAssignments() []Assignment {
	var assignments []Assignment
	db := GetDatabase()
	err := db.Select(&assignments, "SELECT * FROM assignment WHERE status='Open'")
	ErrorHandling(err)
	return assignments
}

func GetKidAssignments(kidId int64) []Assignment {
	var result []Assignment
	db := GetDatabase()
	db.Select(&result, "SELECT * FROM assignment WHERE kidid = ?", kidId)
	return result
}
