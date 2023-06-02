package repo

import (
	"awesomeProject/general"
	"awesomeProject/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() (*UserRepo, *general.Error) {
	user := &UserRepo{}
	var error error
	dsn := "host=localhost user=postgres password=karen1984 dbname=testdb port=5432 TimeZone=Asia/Shanghai"
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		panic(error)
		err := general.NewError("Error Connecting to DB", 1001)
		return nil, err
	}
	error = db.AutoMigrate(&model.User{})
	if error != nil {
		panic(error)
		err := general.NewError("Error Migrating DB", 1001)
		return nil, err
	}
	user.db = db
	return user, nil
}

func (repo *UserRepo) CreateUser(user model.User) {
	u := &model.User{
		ID:       user.ID,
		Username: user.Username,
	}
	aa := repo.db.Create(u)
	println(aa.Error)
	fmt.Println("User Created")
}

func (repo *UserRepo) GetUserById(id int) *gorm.DB {
	var user = model.User{ID: id}
	return repo.db.Find(&user)
}

func (repo *UserRepo) UpdateUsername(id int, username string) {
	repo.db.Save(model.User{ID: id, Username: username})
}

func (repo *UserRepo) DeleteUser(id int) {
	repo.db.Delete(model.User{ID: id})
}
