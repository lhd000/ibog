package model

import (
	"log"
	"time"
)

type Admin struct {
	ID       int64  `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	UserName string `gorm:"type:varchar(30);unique;not null" form:"username" json:"username"`
	PassWord string `gorm:"type:varchar(64);unique;not null" form:"password" json:"password"`
	AddTime  time.Time
}

func (a *Admin) Add() error {

	a.AddTime = time.Now()
	a.PassWord = md5V(a.PassWord)

	LOCK.Lock()
	defer LOCK.Unlock()
	return DB.Create(a).Error
}

func AdminList() []*Admin {

	list := make([]*Admin, 0)
	DB.Find(&list)
	return list
}

func (a *Admin) DeleteByID() error {

	LOCK.Lock()
	defer LOCK.Unlock()
	return DB.Where("id = ?", a.ID).Delete(a).Error
}

func (a *Admin) Login() int64 {

	re := Admin{}
	DB.Where("user_name = ? and pass_word = ?", a.UserName, md5V(a.PassWord)).First(&re)
	log.Println("the login result is:", re)
	return re.ID
}
