package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = "5433"
	user     = "postgres"
	password = "jokammaestro"
	dbname   = "h8-final-project"
	dialect  = "postgres"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}
}

func handleCreateRequiredTables() {
	userTable := `
		CREATE TABLE IF NOT EXISTS "users" (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL,
			password TEXT NOT NULL,
			age int NOT NULL,
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now()
		);
	`

	photoTable :=
		`
	CREATE TABLE IF NOT EXISTS "photos" (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		caption VARCHAR(255) NOT NULL,
		photoUrl TEXT NOT NULL,
		userId int NOT NULL,
		createdAt timestamptz DEFAULT now(),
		updatedAt timestamptz DEFAULT now(),
		CONSTRAINT photos_user_id_fk
			FOREIGN KEY(userId)
				REFERENCES users(id)
					ON DELETE CASCADE
	);
	`

	commentTable :=
		`
	CREATE TABLE IF NOT EXISTS "comments" (
		id SERIAL PRIMARY KEY,
		userId int NOT NULL,
		photoId int NOT NULL,
		message VARCHAR(255) NOT NULL,
		createdAt timestamptz DEFAULT now(),
		updatedAt timestamptz DEFAULT now(),
		CONSTRAINT comment_user_id_fk
			FOREIGN KEY(userId)
				REFERENCES users(id)
					ON DELETE CASCADE,
		CONSTRAINT comment_photo_id_fk
			FOREIGN KEY(photoId)
				REFERENCES photos(id)
					ON DELETE CASCADE
	);
	`

	socialMediaTable :=
		`
	CREATE TABLE IF NOT EXISTS "socialMedias" (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		social_media_url VARCHAR(255) NOT NULL,
		userId int NOT NULL,
		createdAt timestamptz DEFAULT now(),
		updatedAt timestamptz DEFAULT now(),
		CONSTRAINT socialMedia_user_id_fk
			FOREIGN KEY(userId)
				REFERENCES users(id)
					ON DELETE CASCADE
	);
	`

	createTableQueries := fmt.Sprintf("%s %s %s %s", userTable, photoTable, commentTable, socialMediaTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}

}

func InitializeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
