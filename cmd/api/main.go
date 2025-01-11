package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sketch873/go-departments-project/internal/api"
	"github.com/sketch873/go-departments-project/internal/department"
	"github.com/sketch873/go-departments-project/internal/middleware"
	"github.com/sketch873/go-departments-project/internal/user"
)

var provider *api.Provider

func main() {
	provider = api.Init()

	router := gin.Default()
	router.Use(middleware.EntryMiddleware(provider))

	userEndpoints := router.Group("/user")
	userEndpoints.Use(middleware.AuthMiddleware())
	user.SetUserControllers(userEndpoints, router)

	departmentEndpoints := router.Group("/department")
	departmentEndpoints.Use(middleware.AuthMiddleware())
	department.SetDepartmentControllers(departmentEndpoints)


	router.Run()
}