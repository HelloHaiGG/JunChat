package models

import (
	"JunChat/common/igorm"
	"errors"
)

type UserInfo struct {
	CreateAt int64  `json:"create_at"`
	Uid      string `json:"uid" gorm:"PRIMARY_KEY;column:uid"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Password string `json:"password" gorm:"password"`
	IsLogin  bool   `json:"is_login" gorm:"is_login"`
}

func (p *UserInfo) FindByUid() error {
	if p.Uid == "" {
		return errors.New("Uid 不能为空. ")
	}
	if err := igorm.DbClient.Model(UserInfo{}).First(p).Error; err != nil {
		return err
	}
	return nil
}

func (p *UserInfo) FindByName() error {
	if p.UserName == "" {
		return errors.New("User Name 不能为空. ")
	}
	if err := igorm.DbClient.Model(UserInfo{}).First(p).Error; err != nil {
		return err
	}
	return nil
}

func (p *UserInfo) Register() error {
	if p.Uid == "" || p.UserName == "" || p.Password == "" {
		return errors.New("缺少用户信息. ")
	}
	if err := igorm.DbClient.Model(UserInfo{}).Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (p *UserInfo) LoginByUid() error {
	if p.Uid == "" || p.Password == "" {
		return errors.New("缺少用户信息. ")
	}
	if err := igorm.DbClient.Model(UserInfo{}).
		Where("uid = ? and password = ?", p.Uid, p.Password).
		Update("is_login", true).Error; err != nil {
		return err
	}
	return nil
}

func (p *UserInfo) LoginByUserName() error {
	if p.UserName == "" || p.Password == "" {
		return errors.New("缺少用户信息. ")
	}
	if err := igorm.DbClient.Model(UserInfo{}).
		Where("user_name = ? and password = ?", p.UserName, p.Password).
		Update("is_login", true).Error; err != nil {
		return err
	}
	return nil
}
