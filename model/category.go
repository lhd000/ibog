package model

import (
	"time"
)

type Category struct {
	ID       int64     `json:"id" gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name     string    `json:"name" gorm:"type:varchar(30);unique;not null" form:"name" json:"name"`
	AliaName string    `json:"alianame" gorm:"type:varchar(30);unique;not null" form:"alianame" json:"alianame,Omitempty"`
	Sort     int64     `json:"sort" gorm:"type:int(10);not null" form:"sort" json:"sort"`
	ParentID int64     `json:"parentid" gorm:"type:int(10);not null" form:"parentid" json:"parentid"`
	KeyWord  string    `json:"keyword" gorm:"type:varchar(30);not null" form:"keyword" json:"keyword"`
	Des      string    `json:"des" gorm:"type:text" form:"des" json:"des"`
	AddTime  time.Time `json:"addtime" gorm:"type:varchar(30)" form:"addtime" json:"addtime,Omitempty"`
}

func AddCate(c *Category) error {

	LOCK.Lock()
	defer LOCK.Unlock()

	return DB.Create(c).Error
	//return DB.NewRecord(c)
}

func GetCateList() []Category {

	list := make([]Category, 0)
	DB.Find(&list)

	return list
}

func DeleteCategory(id int64) {

	LOCK.Lock()
	defer LOCK.Unlock()
	DB.Where("id = ?", id).Delete(&Category{ID: id})
}

func (c *Category) GetArticleListByCid() []*Article {

	list := make([]*Article, 0)
	DB.Where("cid = ?", c.ID).Order("id desc").Find(&list)
	return list
}

func (c *Category) GetCategoryByCid() *Category {

	re := &Category{}
	DB.Where("id = ?", c.ID).First(&re)
	return re

}
