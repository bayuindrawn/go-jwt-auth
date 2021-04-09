package main

import (
	"fmt"
	"net/http"

	"github.com/bayuindrawn/go-auth-jwt/config"
	"github.com/bayuindrawn/go-auth-jwt/controllers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	db, err := config.BDInit()
	if err != nil {
		panic(err.Error())
	}
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/login", loginHandler)
	router.GET("/person/:id", auth, inDB.GetPerson)
	router.GET("/persons", auth, inDB.GetPersons)
	router.POST("/person", auth, inDB.CreatePerson)
	router.PUT("/person", auth, inDB.UpdatePerson)
	router.DELETE("/person/:id", auth, inDB.DeletePerson)
	router.Run(":3000")
}

func loginHandler(c *gin.Context) {
	var user Credential
	err := c.Bind(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "wrong username or password",
		})
		return
	}
	if c.PostForm("username") != "myname" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
		return

	} else {
		if c.PostForm("password") != "myname123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
			return
		}
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token varified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
