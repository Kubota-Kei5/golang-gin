package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func ConnectToDB() (db *gorm.DB, err error) {
	user := GetEnvDefault("MYSQL_USER", "app")
	password := GetEnvDefault("MYSQL_PASSWORD", "password")
	database := GetEnvDefault("MYSQL_DATABASE", "album_database")
	host := GetEnvDefault("MYSQL_HOST", "mysql")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		user,
		password,
		host,
		3306,
		database)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

type Album struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
}

func (a *Album) Create() (*Album, error) {
	if err := DB.Create(&a).Error; err != nil {
		return a, err
	}
	return a, nil
}

func (a *Album) Save() error {
	if err := DB.Save(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Album) Delete() error {
	if err := DB.Where("id = ?", a.ID).Delete(&Album{}).Error; err != nil {
		return err
	}
	return nil
}

func AlbumFindOne(ID int) (*Album, error) {
	var album Album
	if err := DB.First(&album, ID).Error; err != nil {
		return nil, err
	}
	return &album, nil
}