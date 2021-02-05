package model

import "time"

type Article struct {
	ID      int64     `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Title   string    `gorm:"type:varchar(30);unique;not null" form:"title" json:"title"`
	Content string    `gorm:"type:text;not null" form:"content" json:"content"`
	Keyword string    `gorm:"type:varchar(50)" form:"keyword" json:"keyword"`
	Des     string    `gorm:"type:text;" form:"des" json:"des"`
	Tag     string    `gorm:"type:varchar(50);" form:"tag" json:"tag"`
	Timg    string    `gorm:"type:text;" form:"timg" json:"timg"`
	Status  int64     `gorm:"type:int" form:"status" json:"status"`
	Cid     int64     `gorm:"type:int;" form:"cid" json:"cid"`
	AddTime time.Time `gorm:"type:DATETIME;" form:"addtime" json:"addtime"`
}

func (a *Article) Add() error {
	LOCK.Lock()
	defer LOCK.Unlock()
	return DB.Create(a).Error
}

func (a *Article) List() []*Article {

	list := make([]*Article, 0)
	DB.Order("id desc").Find(&list)
	return list
}

func (a *Article) DeleteArticleByID() error {
	LOCK.Lock()
	defer LOCK.Unlock()
	return DB.Where("id = ?", a.ID).Delete(a).Error
}

func (a *Article) FindArticleByID() *Article {

	re := &Article{}
	DB.Where("id = ? ", a.ID).First(&re)
	return re
}

func (a *Article) UpdateArticleByID() error {
	LOCK.Lock()
	defer LOCK.Unlock()
	return DB.Model(&a).Where("id= ?", a.ID).Updates(a).Error
}
