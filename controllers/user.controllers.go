package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarun-kavipurapu/gin-gorm/db"
	"github.com/tarun-kavipurapu/gin-gorm/db/models"
	"github.com/tarun-kavipurapu/gin-gorm/types"
	"github.com/tarun-kavipurapu/gin-gorm/utils"
)

func Test(c *gin.Context) {
	log.Print("i am in controllers")
	c.JSON(http.StatusAccepted, gin.H{"message": "Everything ok"})
}

// 1->customer
// 2->restaurant

func FindUsers(c *gin.Context) {
	var users []models.User

	err := db.DB.Find(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err, http.StatusBadRequest))
	}

	c.JSON(http.StatusOK, types.GenerateResponse(gin.H{"data": users}, http.StatusOK))

}

// POST /users
// Create a new user
func SignUp(c *gin.Context) {
	var input types.UserRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		// Handle parsing/binding error
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	user := models.User{
		UserEmail:    input.UserEmail,
		UserPassword: input.UserPassword,
		UserRoleID:   input.UserRoleID,
	}

	//check if the user with same Email exists with the same Role

	userDetails, err := user.Save()
	if err != nil {
		// Handle database error
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	c.JSON(http.StatusOK, types.GenerateResponse(gin.H{"message": "User created successfully", "user": userDetails}, http.StatusOK))

}

func LoginUser(c *gin.Context) {
	var input types.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	user, err := models.GetUserByEmail(input.UserEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	err = user.ValidatePassword(input.UserPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	acesstoken, err := utils.GenerateJWT(user)
	// log.Print(acesstoken)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
	}
	refreshToken, err := models.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
	}
	c.JSON(http.StatusOK, types.GenerateResponse(gin.H{"email": input.UserEmail, "accessToken": acesstoken, "refreshToken": refreshToken}, http.StatusOK))

}

// func SeedRole(c *gin.Context) {
// 	// Define roles to be inserted into the database
// 	var roles = []models.Role{
// 		{RoleName: "customer"},
// 		{RoleName: "restaurant"},
// 	}

// 	// Attempt to save roles to the database
// 	err := db.DB.Create(&roles).Error
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// If roles were successfully saved, respond with success message and the roles
// 	c.JSON(http.StatusOK, gin.H{"message": "Roles created successfully", "roles": roles})
// }
