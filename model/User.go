package model

import (
	"fmt"

	"github.com/Farheen-cell/UserCRUD/config"
	_ "github.com/go-sql-driver/mysql"
)

// GetAllUser fetch all user
func GetAllUsers(user *User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreatUser ...	Insert new User

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID...Fetch only one user by id

func GetUserById(user *User, id string) (err error) {
	if err = config.DB.Where("id =?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ...update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user

func DeleteUser(user *User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
