package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sketch873/go-departments-project/internal/session"
	userModel "github.com/sketch873/go-departments-project/pkg/mysql/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User = userModel.User
type LOGIN = userModel.LOGIN
type REGISTER = userModel.REGISTER

func getUser(c *gin.Context) {

	val, err := session.GetSession(c)
	if err != nil {
		c.JSON(http.StatusNotFound, "")
	}

	user, err := userModel.GetUserByUuid(val)
	if err != nil {
		c.JSON(http.StatusNotFound, "")
	}

	fmt.Println(time.Now().UTC())

	c.JSON(http.StatusFound, user)
}

func register(c *gin.Context) {

	var creds REGISTER

	if err := c.BindJSON(&creds); err != nil {
		panic(err)
	}

	if len(creds.PASSWORD) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": nil,
			"error":   "Password must contain at least 8 characters",
		})
		return
	}

	if strings.ToLower(creds.PASSWORD) == creds.PASSWORD {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": nil,
			"error":   "Password must contain at least one upper-case character",
		})
		return
	}

	if strings.ToUpper(creds.PASSWORD) == creds.PASSWORD {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": nil,
			"error":   "Password must contain at least one lower-case character",
		})
		return
	}

	f := func(r rune) bool {
		return (r < 'A' || r > 'Z') && (r < 'a' || r > 'z')
	}

	if strings.IndexFunc(creds.PASSWORD, f) == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": nil,
			"error":   "Password must contain at least one special character",
		})
		return
	}

	if creds.PASSWORD != creds.CONF_PASSWORD {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": nil,
			"error":   "Passwords must match",
		})
		return
	}

	uuid, err := userModel.CreateUser(creds)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, uuid)
}

func login(c *gin.Context) {
	var login LOGIN
	c.BindJSON(&login)

	user, err := userModel.GetUserByEmail(login.EMAIL)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.PASSWORD))
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	session.SetSession(c, user.Uuid)
}

func logout(c *gin.Context) {
	session.DeleteSession(c)
}

func SetUserControllers(rg *gin.RouterGroup, engine *gin.Engine) {
	engine.POST("/register", register)
	engine.POST("/login", login)
	rg.GET("/me", getUser)
	engine.POST("/logout", logout)
}
