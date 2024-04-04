# test-golang-api

### Routes information
```
GET    /api/v1/users/test        --> github.com/tarun-kavipurapu/gin-gorm/controllers.Test (3 handlers)
2024-04-04 16:19:21 [GIN-debug] GET    /api/v1/users/user        --> github.com/tarun-kavipurapu/gin-gorm/controllers.FindUsers (3 handlers)
2024-04-04 16:19:21 [GIN-debug] POST   /api/v1/users/signup      --> github.com/tarun-kavipurapu/gin-gorm/controllers.SignUp (3 handlers)
2024-04-04 16:19:21 [GIN-debug] POST   /api/v1/users/login       --> github.com/tarun-kavipurapu/gin-gorm/controllers.LoginUser (3 handlers)
2024-04-04 16:19:21 [GIN-debug] POST   /api/v1/restaurant/createRestaurant --> github.com/tarun-kavipurapu/gin-gorm/controllers.CreateRestaurant (4 handlers)
```

### Setup information 
```
docker build -t test-golang .
```

```
docker compose up
```
### Information for testing API
- Only The CreateRestaurant endpoint is written
- In Database 1-->Customer 2-->Restaurant so while signing up send the role_id as 2 for acessing the create restaurant endpoint
#### Sign up
```
{
  "user_email":"restaurant1@gmail.com",
  "user_password":"12345678",
  "user_role_id":2
}
```
#### Login
```
{
  "user_email":"restaurant1@gmail.com",
  "user_password":"12345678"
}
```
#### Create Restaurant
```
provide the Auth token in here which is obtained from the login endpoint
```
```
{
  "restaurant_name": "Your Restaurant Name",
  "restaurant_contact": "+15551234567",
  "restaurant_address": {
    "address_street": "123 Main Street",
    "address_city": "Anytown",
    "address_state": "CA",
    "address_country": "USA",
    "address_zipcode": "12345"
  }
}
```