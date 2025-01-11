package user

import (
	"github.com/google/uuid"
	"github.com/sketch873/go-departments-project/pkg/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Uuid      string    `json:"uuid"`
}

type REGISTER struct {
	EMAIL         string `json:"email" binding:"required"`
	PASSWORD      string `json:"password" binding:"required"`
	CONF_PASSWORD string `json:"confirm_password" binding:"required"`
	USERNAME      string `json:"username" binding:"required"`
	FIRST_NAME    string `json:"first_name" binding:"required"`
	LAST_NAME     string `json:"last_name" binding:"required"`
}

type LOGIN struct {
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

func GetUserByUuid(uuid string) (*User, error) {
	db, _ := mysql.GetDatabase(nil)

	var user User;

	db.Raw("CALL GetUserByUuid(?)", uuid).Scan(&user)

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {

	var user User

	db, _ := mysql.GetDatabase(nil)

	db.Raw("CALL GetUserByEmail(?)", email).Scan(&user)

	return &user, nil
}

func CreateUser(creds REGISTER) (*string, error) {
	uuid := "u-" + uuid.New().String()

	db, _ := mysql.GetDatabase(nil)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.PASSWORD), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	tx := db.Exec("CALL CreateUser(?, ?, ?, ?, ?, ?)", creds.EMAIL, string(hashedPassword), creds.USERNAME, creds.FIRST_NAME, creds.LAST_NAME, uuid)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &uuid, nil
}