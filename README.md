# Go Auth with JWT
<p align="center">
  <img src="https://www.nicepng.com/png/full/370-3707528_65159967-golang-logo.png" />
</p>

### Requirement

- Docker

### Run docker compose

```cmd
$ docker-compose up
```

### Dependencies

```go
require (
	github.com/gofiber/fiber/v2 v2.25.0
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.5
    	github.com/dgrijalva/jwt-go
    	golang.org/x/crypto/bcrypt
)
```

### Database

- PostgreSQL
