package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID int
	Phone string
	Gender int8
}

func (User)TableName()string  {
	return "hr_user"
}

//单插入
func Insert(db *gorm.DB){
	user := User{Phone: "嗷嗷嗷哦",Gender: 2}
	db.Create(&user)
	fmt.Println(user.ID)
}

//批量插入
func Inserts(db *gorm.DB)  {
	users := []User{{Phone: "嗷嗷嗷哦1",Gender: 2},{Phone: "嗷嗷嗷哦2",Gender: 2}}
	db.Create(&users)
	for _,v := range users {
		fmt.Println(v.ID)
	}
}

func main()  {
	dsn := "root:1234@tcp(127.0.0.1:3306)/hongren2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil {
		panic(err)
	}

	//单插入
	//Insert(db)
	//批量插入
	//Inserts(db)
	user := User{}
	db.First(&user)
	fmt.Printf("%#v",user)

}
