package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarun-kavipurapu/gin-gorm/db/models"
	"github.com/tarun-kavipurapu/gin-gorm/types"
	"github.com/tarun-kavipurapu/gin-gorm/utils"
)

func CreateRestaurant(c *gin.Context) {
	var requestBody types.RestaurantRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	// Validate request body
	//if err := validateRestaurantRequest(&requestBody); err != nil {
	//  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//return
	//}
	address := models.Address{
		Street:  requestBody.RestaurantAddress.AddressStreet,
		City:    requestBody.RestaurantAddress.AddressCity,
		State:   requestBody.RestaurantAddress.AddressState,
		Country: requestBody.RestaurantAddress.AddressCountry,
		Zipcode: requestBody.RestaurantAddress.AddressZipcode,
	}
	addressDetails, err := address.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
	}
	user := utils.CurrentUser(c)
	restaurant := models.Restaurant{
		RestaurantName:    requestBody.RestaurantName,
		RestaurantContact: requestBody.RestaurantContact,

		RestaurantUserID:    user.UserID,
		RestaurantAddressID: address.AddressID,
	}
	restaurantDetails, err := restaurant.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, types.GenerateErrorResponse(err.Error(), http.StatusBadRequest))
	}
	c.JSON(http.StatusOK, types.GenerateResponse(gin.H{"message": "Restaurant created successfully", "restaurant": restaurantDetails, "address": addressDetails}, http.StatusOK))

}

// func
