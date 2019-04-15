package models

//用户表
type User struct {
	BaseModel
	Username  string `json:"username"`
	Password  string `json:"-"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	RoleId    int    `json:"role_id"  gorm:"default:0"` //角色 默认 0 普通用户
	PersonId  int    `json:"person_id"`                 //身份证
	IconUrl   string `json:"icon_url"`                  //头像
	LicenseId string `json:"license_id"`                //执照信息 id
}

func InsertUser(user *User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//查找用户信息
func FindUserById(id uint64) (*User, error) {
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
