package auth

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	secretKey = "ascomelinyicun20191219"
)
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

type CustomClaims struct {
	Username string
	ID       int
	jwt.StandardClaims
}

func CreatCookie(username string, id int) (string, time.Time) {
	// Expires the token and cookie in 24 hours
	expireToken := time.Now().Add(time.Hour * time.Duration(24)).Unix()
	expireCookie := time.Now().Add(time.Hour * time.Duration(24))
	// We'll manually assign the claims but in production you'd insert values from a database
	claims := CustomClaims{
		username,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    username,
		},
	}
	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signs the token with a secret.
	signedToken, _ := token.SignedString([]byte(secretKey))
	return signedToken, expireCookie
}

func CheckAjax(r *http.Request) bool {
	header := r.Header.Get("X-Requested-With")
	if header == "XMLHttpRequest" {
		return true
	} else {
		return false
	}
}

func resPonseError(c *gin.Context) {
	fmt.Printf("\033[1;31m Response Error ==========\033[0m\n")
	isajax := CheckAjax(c.Request)
	if isajax {
		c.JSON(http.StatusUnauthorized, "アクセス権限ありません")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/user/login")
	}
}

func ValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If no Auth cookie is set then return a 401
		cookie, err := c.Cookie("_tk")
		fmt.Printf("\033[1;34m ====== Get cookie = %s\033[0m\n", cookie)
		if err != nil {
			fmt.Printf("\033[1;31m Get cookie error : %s \033[0m\n", err.Error())
			resPonseError(c)
		} else {
			token, err := jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Prevents a known exploit
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("\033[1;31m Unexpected signing method %v \033[0m", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				fmt.Printf("\033[1;31m Token parse error : %s \033[0m\n", err.Error())
				resPonseError(c)
				return
			} else {
				// Validate the token and save the token's claims to a context
				if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
					fmt.Printf("claims.Username : %s, Now Route to URL : %s \n", claims.Username, c.Request.RequestURI)
					signedToken, expireCookie := CreatCookie(claims.Username, claims.ID)
					cook := http.Cookie{
						Name:    "_tk",
						Value:   signedToken,
						Path:    "/",
						Expires: expireCookie,
					}
					http.SetCookie(c.Writer, &cook)
					c.Next()
					return
				} else {
					// fmt.Printf("\033[1;31m  Valid Error : user name is %s %v \033[0m\n", claims.Username)
					resPonseError(c)
					return
				}
			}
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Set("content-type", "application/json")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
