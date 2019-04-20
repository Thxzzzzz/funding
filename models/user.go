package models

//用户表
type User struct {
	BaseModel
	Username         string `json:"username"`                  //用户名
	Password         string `json:"-"`                         //密码 这里 tag 设置为 - ,保证其永远不会被解析为 Json
	Nickname         string `json:"nickname"`                  //昵称
	Email            string `json:"email"`                     //邮箱
	Phone            string `json:"phone"`                     //手机号
	RoleId           int    `json:"role_id"  gorm:"default:0"` //角色  ( 0:普通用户(默认) )
	PersonId         int    `json:"person_id"`                 //身份证
	IconUrl          string `json:"icon_url"`                  //头像
	DefaultAddressId uint64 `json:"default_address_id"`        //默认地址
	LicenseId        string `json:"license_id"`                //执照信息 id
}

//插入一条新的用户信息
func InsertUser(user *User) error {
	err := db.Create(&user).Error
	return err
}

//查找用户信息
func FindUserById(id uint64) (*User, error) {
	var result User
	err := db.First(&result, id).Error
	return &result, err
}

//根据用户名查找用户信息
func FindUserByUsername(username string) (*User, error) {
	var result User
	//// SELECT * FROM users WHERE username = "username" LIMIT 1;
	err := db.Where(&User{Username: username}).First(&result).Error
	return &result, err
}

//根据 User.ID 来更新其他相应的字段
func UpdateUser(user *User) error {
	var ret User
	err := db.First(&ret, "id = ?", user.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&ret).Update(user).Error
	return err
}
