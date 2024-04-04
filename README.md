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