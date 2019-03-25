package models

//用户表
type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"-"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	//角色
	RoleId int `json:"role_id"`
	//身份证
	PersonId int `json:"person_id"`
	//头像
	IconUrl string `json:"icon_url"`
	//执照信息 id
	LicenseId string `json:"license_id"`
}

//查找用户信息
func FindUserById(id int64) (*User, error) {
	var result User
	err := db.First(&result, id).Error
	return &result, err
}

func FindUserByUsername(username string) (*User, error) {
	var result User
	//// SELECT * FROM users WHERE username = "username" LIMIT 1;
	err := db.Where(&User{Username: username}).First(&result).Error
	return &result, err
}
