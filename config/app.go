package config

//importing G(ORM) which is Go Object Relational Mapping, built by jinzhu
//Using MySQL dialect specifically
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//creating global variable db which retrieves from gorm using pointer reference
var (
	db *gorm.DB
)

//creating Connect function to config the Database and handle errors if occured and returning db variable
func Connect() {
	d, err := gorm.Open("mysql", "username:password/table?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

//creating a function "getDB" to return db var which points to gorm.db
func GetDB() *gorm.DB {
	return db
}
