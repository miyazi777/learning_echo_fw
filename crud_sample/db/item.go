package db

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
}

type ItemRepository interface {
	GetList() *[]Item
}

type ItemRepositoryImpl struct{}

func (i *ItemRepositoryImpl) GetList() *[]Item {
	db := getDbConnection()
	defer db.Close()

	db = db.Order("id desc")

	items := []Item{}
	db.Find(&items)

	return &items
}
