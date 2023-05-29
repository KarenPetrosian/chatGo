package main

import (
	"awesomeProject/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=karen1984 dbname=testdb port=5432 TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.WithError(err)
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}

	u := &model.User{}
	//db.Create(u)
	db.Delete(u, 1)

	//logrus.Infof("DB %s", db)
	//mux := http.NewServeMux()
	//mux.HandleFunc("/users", model.CreateUser)
	//http.ListenAndServe(":9000", mux)

	router := gin.Default()

	router.Use(gin.Recovery())

	groupUser := router.Group("/users")
	_ = router.Group("/customer")

	groupUser.GET("/:id/add", createUser)
	groupUser.GET("", createUser)
	//router.POST("/users", createUser)

	router.Run("localhost:9000")
}

func createUser(context *gin.Context) {
	//username := context.Query("username")
	//mail := context.Query("email")
	id := context.Param("id")

	//if username == "" || mail == "" {
	//	context.JSON(400, "BAD REQUEST")
	//	return
	//
	//}

	val, _ := strconv.Atoi(id)

	u := &model.User{
		ID:       val,
		Username: "username",
		Email:    "mail",
	}

	go db.Create(u)
	fmt.Println("Done")
	context.JSON(200, u)
}
