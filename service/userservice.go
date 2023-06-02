package service

import (
	"awesomeProject/general"
	"awesomeProject/model"
	"awesomeProject/repo"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService() (*UserService, *general.Error) {
	service := &UserService{}
	error := &general.Error{}
	service.repo, error = repo.NewUserRepo()

	if error != nil {
		return nil, error
	}
	router := gin.Default()
	router.Use(gin.Recovery())
	//groupUser := router.Group("/users")
	_ = router.Group("/customer")

	//groupUser.GET("/:id/add", CreateUser)
	//groupUser.GET("", CreateUser)
	router.GET("/users", func(c *gin.Context) {
		CreateUser(c, service)
	})
	router.Run("localhost:9000")

	return service, nil
}

func GetUsers(context *gin.Context) {

}

func CreateUser(context *gin.Context, service *UserService) {
	//service := &UserService{}
	error := &general.Error{}
	println(error)
	service.repo, error = repo.NewUserRepo()
	service.createUser(context)
}

func (service UserService) createUser(context *gin.Context) {
	//id := context.Param("id")
	//username := context.Param("username")
	//val, _ := strconv.Atoi(id)
	u := model.User{
		ID:       2,
		Username: "Karen",
	}
	//serviceA := service
	service.repo.CreateUser(u)

	fmt.Println("Done")
	context.JSON(200, u)
}
