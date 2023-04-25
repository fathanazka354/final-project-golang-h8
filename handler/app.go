package handler

import "final-project/database"

func StartApp() {
	database.InitializeDatabase()
	db := database.GetDatabaseInstance()
	_ = db
}
