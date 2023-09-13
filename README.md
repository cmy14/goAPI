# goAPI
Création d'api rest en go
# initialiser un projet
go mod init api-go
# installer gin pour l'api  rest

go get github.com/gin-gonic/gin

# installer gorm

go get -u gorm.io/gorm
# AJOUT  FRAMEWORK
go get -u gorm.io/driver/postgres

@
// package that we will be used to authenticate and generate our JWT
go get -u github.com/dgrijalva/jwt-go
// to help manage our environment variables
   go get    -u github.com/joho/godotenv
// to encrypt our user's password
go get -u golang.org/x/crypto