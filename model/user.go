package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/mesment/server/pkg/constvar"
	"github.com/mesment/server/pkg/errno"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=4,max=128"`
}

func (c *UserModel) TableName() string {
	return "users"
}

//新增一个用户账号
func (u *UserModel) Create() error {
	return DB.DB.Create(&u).Error
}

//删除指定id的用户
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.DB.Delete(&user).Error
}

//更新一个用户的信息
func (u *UserModel) UpdateUser() error {
	return DB.DB.Save(&u).Error
}

func GetUser(name string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.DB.Where("username = ?", name).First(&u)
	return u, d.Error
}

func IsUsernameExist(name string) bool {
	var count uint64
	if err := DB.DB.Model(&UserModel{}).Where("username = ?", name).Count(&count).Error; err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func ListUsers(name string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", name)

	if err := DB.DB.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.DB.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil

}

func (u *UserModel) Compare(pwd string) bool {
	if err := u.Encrypt(); err != nil {
		return false
	}
	return u.Password == pwd
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	h := md5.New()
	h.Write([]byte(u.Password))
	hex.EncodeToString(h.Sum(nil))
	u.Password = u.MD5Password()
	return
}

func (u UserModel) MD5Password() string {
	h := md5.New()
	h.Write([]byte(u.Password))
	return hex.EncodeToString(h.Sum(nil))

}

// Validate the fields.
func (u *UserModel) Validate() error {
	var err error
	//对用户名密码进行检查
	if u.Username == "" {
		err = errno.New(errno.ErrUserNameIsEmpty, fmt.Errorf("username:%s", u.Username))
	}

	if u.Password == "" {
		err = errno.New(errno.ErrPasswordIsEmpty, fmt.Errorf("password:%s", u.Password))
	}
	return err
}
