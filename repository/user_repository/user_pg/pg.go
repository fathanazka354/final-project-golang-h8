package user_pg

import (
	"database/sql"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/user_repository"
)

const (
	createUserQuery = `
		INSERT INTO "users"
		(
			"username",
			"email",
			"password",
			"age"
		) VALUES ($1,$2,$3,$4);
	`

	getUserByEmailQuery = `
		SELECT id, username, email, password FROM "users"
		WHERE email = $1;
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) CreateNewUser(user entity.User) errs.MessageErr {
	_, err := u.db.Exec(createUserQuery, user.Username, user.Email, user.Password, user.Age)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	return nil, nil
}

func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	row := u.db.QueryRow(getUserByEmailQuery, userEmail)

	var user entity.User

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &user, nil
}
