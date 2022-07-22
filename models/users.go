package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-redis/redis"
	"github.com/gofrs/uuid"
	_ "github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// var client *redis.Client

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 	}

// 	dsn := os.Getenv("REDIS_DSN")
// 	if len(dsn) == 0 {
// 		dsn = "localhost:6379"
// 	}
// 	client = redis.NewClient(&redis.Options{
// 		Addr: dsn, //redis port
// 	})
// 	_, err := client.Ping().Result()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func CheckEmailExist(email string) (User, bool) {
	ut := User{}
	err := db.Get(&ut, "SELECT * FROM user WHERE email=?", email)
	return ut, err != sql.ErrNoRows
}

func CreateNewUser(u User) {
	db.MustExec("INSERT INTO user(email, password) VALUES (?, ?)", u.Email, u.Password)
}

func CreateToken(userId int) (*TokenDetails, error) {
	var err error
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	uAT, err := uuid.NewV4() //create an uuid
	if err != nil {
		return nil, err
	}
	td.AccessUuid = uAT.String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	uRT, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	td.RefreshUuid = uRT.String()

	secretAT, exists := os.LookupEnv("ACCESS_SECRET")
	if !exists {
		return nil, fmt.Errorf("no secret access found")
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["id_user"] = userId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(secretAT))
	if err != nil {
		return nil, err
	}

	// secretRT, exists := os.LookupEnv("REFRESH_SECRET")
	// if !exists {
	// 	return nil, fmt.Errorf("no secret access found")
	// }
	// rtClaims := jwt.MapClaims{}
	// rtClaims["refresh_uuid"] = td.RefreshUuid
	// rtClaims["id_user"] = userId
	// rtClaims["exp"] = td.RtExpires
	// rt := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// td.RefreshToken, err = rt.SignedString([]byte(secretRT))
	// if err != nil {
	// 	return nil, err
	// }
	return td, nil
}

func GetInfoUser(r *http.Request, token string) (User, error) {
	u := User{}
	secret, exists := os.LookupEnv("ACCESS_SECRET")
	if !exists {
		return u, fmt.Errorf("no access key found")
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return u, fmt.Errorf(err.Error())
	}

	err = db.Get(&u, "SELECT * FROM user WHERE id=?", claims["id_user"])
	if err != nil {
		return u, err
	}
	return u, nil
}
