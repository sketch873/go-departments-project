package api

import (
	"github.com/sketch873/go-departments-project/internal/environment"
	"github.com/sketch873/go-departments-project/pkg/mysql"
	"github.com/sketch873/go-departments-project/pkg/redis"
	"gorm.io/gorm"
)

type Provider struct {
	Db     *gorm.DB
	Redis  redis.Redis
}

func Init() *Provider {
	environment.LoadEnv()

	db, err := mysql.GetDatabase(nil)
	if err != nil {
		panic(err)
	}

	redis := redis.CreateClient()

	return &Provider{Db: db, Redis: redis}
}