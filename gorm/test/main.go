package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID   string // 名为`ID`的字段会默认作为表的主键
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.Table("test_table").CreateTable(&User{})
	fmt.Println(db.HasTable("test_table"))
	first := User{ID:"1",Name:"liyan"}
	db.NewRecord(first)
	db.Create(&first)
	var userinfo User
	db.First(&userinfo,"ID = ?","1")
	fmt.Println(userinfo)
}