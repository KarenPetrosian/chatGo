package main

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB

func main() {

	service, error := service.NewUserService()
	if error != nil {
		fmt.Println(error)
		return
	}
	println(service)

	//var err error
	//dsn := "host=localhost user=postgres password=karen1984 dbname=testdb port=5432 TimeZone=Asia/Shanghai"
	//db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	logrus.WithError(err)
	//	panic(err)
	//}
	//
	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	return
	//}
	//
	//u := &model.User{}
	//db.Delete(u, 1)
	//router := gin.Default()
	//router.Use(gin.Recovery())
	//groupUser := router.Group("/users")
	//_ = router.Group("/customer")
	//
	//groupUser.GET("/:id/add", createUser)
	//groupUser.GET("", createUser)
	//router.POST("/users", createUser)

	//err = router.Run("localhost:9000")
	//if err != nil {
	//	return
	//}
}

func createUser(context *gin.Context) {
	id := context.Param("id")
	val, _ := strconv.Atoi(id)
	u := &model.User{
		ID:       val,
		Username: "username",
	}
	go db.Create(u)
	fmt.Println("Done")
	context.JSON(200, u)
}
