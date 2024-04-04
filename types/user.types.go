package types

type UserRegister struct {
	UserEmail    string `json:"user_email" binding:"email,required"`
	UserPassword string `json:"user_password" binding:"required"`
	UserRoleID   uint   `json:"user_role_id" binding:"required"`
}
type UserLogin struct {
	UserEmail    string `json:"user_email" binding:"email,required"`
	UserPassword string `json:"user_password" binding:"required"`
}
type RestaurantRequest struct {
	RestaurantName    string         `json:"restaurant_name" binding:"required"`
	RestaurantContact string         `json:"restaurant_contact" binding:"required"`
	RestaurantAddress AddressRequest `json:"restaurant_address" binding:"required"`
}
type AddressRequest struct {
	AddressStreet  string `json:"address_street" binding:"required"`
	AddressCity    string `json:"address_city" binding:"required"`
	AddressState   string `json:"address_state" binding:"required"`
	AddressCountry string `json:"address_country" binding:"required"`
	AddressZipcode string `json:"address_zipcode" binding:"required"`
}

type BaseHttpResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}
type BaseErrorResponse struct {
	Status     string      `json:"status"`
	Error      interface{} `json:"error"`
	StatusCode int         `json:"statusCode"`
}

func GenerateResponse(data interface{}, statusCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Status:     "success",
		StatusCode: statusCode,
		Data:       data,
	}
}

func GenerateErrorResponse(err interface{}, statusCode int) *BaseErrorResponse {
	return &BaseErrorResponse{
		Status:     "error",
		StatusCode: statusCode,
		Error:      err,
	}
}
