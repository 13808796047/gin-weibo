package models

import "github.com/jinzhu/gorm"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	gorm.Model
}

func GetUserTotal() (int, error) {
	var count int
	if err := DB.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetUsers(page int, size int) (users []User) {
	// limit定位每页大小, Offset定位偏移的查询位置.并且先进行条件筛选,最后做分页操作.
	DB.Offset((page - 1) * size).Limit(size).Find(&users)
	return
}
func AddUser(user *User) {
	DB.Create(&user)
	return
}
func UserDetailByName(name string) (user User) {
	DB.Where("name = ?", name).First(&user)
	return
}

func UserDetailByEmail(email string) (user User) {
	DB.Where("email = ?", email).First(&user)
	return
}

func UserDetail(id uint) (user User) {
	DB.Where("id = ?", id).First(&user)
	return
}
