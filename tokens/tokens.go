package tokens

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SECRET = []byte("secret")
var api_key = "1234"

func CreateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println(token)
	return tokenStr, nil
}

// func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] != nil {
// 			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
// 				_, ok := t.Method.(*jwt.SigningMethodHMAC)
// 				if !ok {
// 					w.WriteHeader(http.StatusUnauthorized)
// 					w.Write([]byte("not authorized"))
// 				}
// 				return SECRET, nil
// 			})

// 			if err != nil {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				w.Write([]byte("not authorized: " + err.Error()))
// 			}

// 			if token.Valid {
// 				next(w, r)
// 			}
// 		} else {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("not authorized"))
// 		}
// 	})
// }

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Token")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				context.JSON(http.StatusUnauthorized, "not authorized")
			}
			return SECRET, nil
		})

		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		if token.Valid {
			context.Next()
		} else {
			context.JSON(http.StatusUnauthorized, "not authorized")
		}
	}
}

func GetJwt(c *gin.Context) {
	if c.Request.Header["Access"] != nil {
		if c.Request.Header["Access"][0] != api_key {
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, token)
		}
	}
}
