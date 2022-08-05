# go_restapi
learn sturcture of rest api.

## สิ่งที่มีใน Example นี้

- Model / Seed
- Controller
- Middleware
- Dockerfile / Build Command
- Json web token (Authen JWT)
- Test
- Swagger
- Graceful shutdown

```sh
go mod init go_restapi
```

```sh
touch main.go
touch .env
```

```sh
mkdir Auth Config Controllers Middleware Models Routers Seeder Utility
```

```sh
go get github.com/dgrijalva/jwt-go v3.2.0+incompatible
go get github.com/gin-gonic/gin v1.8.1
go get github.com/golang-jwt/jwt v3.2.2+incompatible
go get github.com/joho/godotenv v1.4.0
go get github.com/pseidemann/finish v1.0.0
go get gorm.io/driver/sqlite v1.3.6
go get gorm.io/gorm v1.23.4
```

```sh
go mod tidy
```