package models

// User - struct of user
type User struct{
	ID uint `gorm:"AUTO_INCREMENT" json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	Age uint32 `json:"age"`
}
// GetUsers - get user list
func GetUsers()(users []User,err error){
	err = db.Debug().Find(&users).Error
	return
}

// CreateUser - create user 
func (u *User)CreateUser()(err error){
	err = db.Create(&u).Error
	return
}

// GetUserDetail - get user detail
func (u *User)GetUserDetail()(err error){
	err = db.First(&u).Error
	return
}

//UpdateUser - update user
func (u *User)UpdateUser()(err error){
	err = db.Model(&u).Update(&u).Error
	return
}

//DeleteUser - delete user
func (u *User)DeleteUser()(err error){
	err =db.Delete(&u).Error
	return
}
