package session

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sketch873/go-departments-project/internal/api"
	"github.com/sketch873/go-departments-project/internal/environment"
)

func SetSession(c *gin.Context, value string) {
	provider, _ := c.Value("provider").(*api.Provider)
	sessionID := uuid.NewString()
	maxAge, _ := strconv.Atoi(environment.GetEnv("SESSION_MAX_AGE"))
	c.SetCookie("session", sessionID, maxAge, "/", "", true, true)
	provider.Redis.SetNX(sessionID, value, time.Duration(maxAge)*time.Second)
}

func GetSession(c *gin.Context) (key string, err error) {

	provider, _ := c.Value("provider").(*api.Provider)
	val, err := c.Cookie("session")
	if err != nil {
		return "", errors.New("no cookie")
	}

	val, err = provider.Redis.Get(val)
	if err != nil {
		return "", errors.New("no session")
	}

	return val, nil
}

func DeleteSession(c *gin.Context) (err error) {
	provider, _ := c.Value("provider").(*api.Provider)
	val, err := c.Cookie("session")
	if err != nil {
		return errors.New("no cookie")
	}

	provider.Redis.Delete(val)
	c.SetCookie("session", "", -1, "/", "", false, true)

	return nil
}