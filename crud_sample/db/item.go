package db

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
}

type ItemRepository interface {
	GetList() *[]Item
	FindById(id string) *Item
	Insert(item *Item) error
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

func (i *ItemRepositoryImpl) FindById(id string) *Item {
	db := getDbConnection()
	defer db.Close()

	item := Item{}
	if db.First(&item, id).RecordNotFound() {
		return nil
	}
	return &item
}

func (t *ItemRepositoryImpl) Insert(item *Item) error {
	db := getDbConnection()
	defer db.Close()

	db.Create(item)
	return nil
}
