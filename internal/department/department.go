package department

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sketch873/go-departments-project/internal/user"
	"github.com/sketch873/go-departments-project/pkg/mysql"
)

type Department struct {
	Name     string    `json:"name"`
	Uuid  string    `json:"uuid"`
	Description  string    `json:"description"`
	Flags int    `json:"flags"`
}

func createDepartment(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}


	if _, exists := jsonData["name"]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Invalid name",
		})
		return
	}
	name := jsonData["name"].(string)

	if _, exists := jsonData["flags"]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Invalid flags",
		})
		return
	}
	flags := jsonData["flags"].(float64)

	if _, exists := jsonData["description"]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Invalid description",
		})
		return
	}
	description := jsonData["description"].(string)

	if len(name) > 255 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Invalid name",
		})
		return
	}

	if flags < 0 || flags > 7 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Invalid flags",
		})
		return
	}

	db, _ := mysql.GetDatabase(nil)

	uuid := "d-" + uuid.New().String()

	db.Exec("CALL CreateDepartment(?, ?, ?, ?)", name, flags, description, uuid)

	c.JSON(http.StatusOK, uuid)
}

func getDepartment(c *gin.Context) {
	uuid := c.Param("uuid")

	db, _ := mysql.GetDatabase(nil)

	var department Department;

	db.Raw("CALL GetDepartmentByUuid(?)", uuid).Scan(&department)

	c.JSON(http.StatusOK, department)
}

func updateDepartment(c *gin.Context) {
	uuid := c.Param("uuid")

	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	db, _ := mysql.GetDatabase(nil)

	if _, exists := jsonData["name"]; exists {
		db.Exec("CALL UpdateDepartmentName(?, ?)", jsonData["name"], uuid)
	}

	if _, exists := jsonData["flags"]; exists {
		db.Exec("CALL UpdateDepartmentflags(?, ?)", jsonData["flags"], uuid)
	}

	if _, exists := jsonData["description"]; exists {
		db.Exec("CALL UpdateDepartmentDescription(?, ?)", jsonData["description"], uuid)
	}

}

func getDepartmentUsers(c *gin.Context) {
	uuid := c.Param("uuid")

	db, _ := mysql.GetDatabase(nil)

	var users []user.User;

	db.Raw("CALL GetDepartmentUsers(?)", uuid).Scan(&users)

	c.JSON(http.StatusOK, users)
}

func getActiveDepartments(c *gin.Context) {
	db, _ := mysql.GetDatabase(nil)

	var departments []Department;

	db.Raw("CALL GetActiveDepartments()").Scan(&departments)

	c.JSON(http.StatusOK, departments)
}

func getDeletedDepartments(c *gin.Context) {
	// TODO
}

func getApprovedDepartments(c *gin.Context) {
	db, _ := mysql.GetDatabase(nil)

	var departments []Department;

	db.Raw("CALL GetExistingDepartments()").Scan(&departments)

	c.JSON(http.StatusOK, departments)
}

func getAllDepartments(c *gin.Context) {
	db, _ := mysql.GetDatabase(nil)

	var departments []Department;

	db.Raw("CALL GetAllDepartments()").Scan(&departments)

	c.JSON(http.StatusOK, departments)
}

func updateDepartmentUserList(c *gin.Context) {

	userUuid := c.Param("uuid")
	departmentUuid := c.Param("departmentUuid")

	db, _ := mysql.GetDatabase(nil)

	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	if _, exists := jsonData["op"]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Invalid operation",
		})
		return
	}

	switch jsonData["op"] {
	case "add":
		tx := db.Exec("CALL AddUserInDepartment(?, ?)", userUuid, departmentUuid)
		if tx.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "fail",
				"error":  tx.Error.Error(),
			})
			return
		}
	case "delete":
		tx := db.Exec("CALL DeleteUserFromDepartment(?, ?)", userUuid, departmentUuid)
		if tx.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "fail",
				"error":  tx.Error.Error(),
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Invalid operation",
		})
		return
	}

	c.JSON(http.StatusOK, "Ok")
}

func SetDepartmentControllers(rg *gin.RouterGroup) {
	/*  Ops for one department  */
	rg.PUT("/create", createDepartment)
	rg.GET("/one/:uuid", getDepartment)
	rg.PATCH("/one/:uuid", updateDepartment)
	rg.GET("/one/:uuid/users", getDepartmentUsers)
	rg.PATCH("/add/:uuid/user/:departmentUuid", updateDepartmentUserList)

	/*  Ops for multiple departments  */
	rg.GET("/active", getActiveDepartments)
	rg.GET("/deleted", getDeletedDepartments)
	rg.GET("/approved", getApprovedDepartments)
	rg.GET("/all", getAllDepartments)
}
