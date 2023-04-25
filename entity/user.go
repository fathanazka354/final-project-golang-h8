package entity

import (
	"final-project/pkg/errs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secret_key = "RAHASIA"

var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")

type User struct {
	Id        int       `json:"id" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email,unique"`
	Password  string    `json:"password" validate:"required,min=6,max=255"`
	Age       int       `json:"age" validate:"required,numeric"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) HashPassword() errs.MessageErr {
	salt := 8

	userPassword := []byte(u.Password)

	result, err := bcrypt.GenerateFromPassword(userPassword, salt)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	u.Password = string(result)
	return nil
}

func (u *User) parseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}
		return []byte(secret_key), nil
	})

	if err != nil {
		return nil, invalidTokenErr
	}

	return token, nil

}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	if id, ok := claim["id"].(float64); !ok {
		return invalidTokenErr
	} else {
		u.Id = int(id)
	}

	if email, ok := claim["email"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Email = email
	}

	return nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return invalidTokenErr
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return invalidTokenErr
	}

	tokenString := splitToken[1]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return invalidTokenErr
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    u.Id,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 10).Unix(),
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret_key))
	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
