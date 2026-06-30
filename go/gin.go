package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my-secret-key")

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	users  = make(map[int]User)
	nextID = 1
	mu     sync.RWMutex
)

//////////////////////////////////////////////////
// Logging Middleware
//////////////////////////////////////////////////

func LoggerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		log.Printf(
			"%s %s %d %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}

//////////////////////////////////////////////////
// JWT Middleware
//////////////////////////////////////////////////

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			},
		)

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			c.Abort()
			return
		}

		c.Next()
	}
}

//////////////////////////////////////////////////
// Login API
//////////////////////////////////////////////////

func Login(c *gin.Context) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": "admin",
			"exp":  time.Now().Add(
				time.Hour,
			).Unix(),
		},
	)

	tokenString, _ := token.SignedString(
		secretKey,
	)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

//////////////////////////////////////////////////
// GET
//////////////////////////////////////////////////

func GetUsers(c *gin.Context) {

	mu.RLock()
	defer mu.RUnlock()

	var result []User

	for _, user := range users {
		result = append(result, user)
	}

	c.JSON(http.StatusOK, result)
}

//////////////////////////////////////////////////
// POST
//////////////////////////////////////////////////

func CreateUser(c *gin.Context) {

	var user User

	if err := c.ShouldBindJSON(
		&user,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	mu.Lock()

	user.ID = nextID
	nextID++

	users[user.ID] = user

	mu.Unlock()

	c.JSON(
		http.StatusCreated,
		user,
	)
}

//////////////////////////////////////////////////
// PUT
//////////////////////////////////////////////////

func UpdateUser(c *gin.Context) {

	id, _ := strconv.Atoi(
		c.Param("id"),
	)

	var user User

	if err := c.ShouldBindJSON(
		&user,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := users[id]; !ok {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "user not found",
			},
		)

		return
	}

	user.ID = id

	users[id] = user

	c.JSON(
		http.StatusOK,
		user,
	)
}

//////////////////////////////////////////////////
// DELETE
//////////////////////////////////////////////////

func DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(
		c.Param("id"),
	)

	mu.Lock()
	defer mu.Unlock()

	delete(users, id)

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "deleted",
		},
	)
}

//////////////////////////////////////////////////
// Main
//////////////////////////////////////////////////

func main() {

	r := gin.New()

	r.Use(LoggerMiddleware())
	r.Use(gin.Recovery())

	r.POST("/login", Login)

	api := r.Group("/api")

	api.Use(AuthMiddleware())

	{
		api.GET("/users", GetUsers)

		api.POST(
			"/users",
			CreateUser,
		)

		api.PUT(
			"/users/:id",
			UpdateUser,
		)

		api.DELETE(
			"/users/:id",
			DeleteUser,
		)
	}

	r.Run(":8080")
}