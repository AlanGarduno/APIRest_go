package Conexion

import(
		"github.com/jinzhu/gorm"
		"github.com/jinzhu/gorm/dialects/mysql")

var connection *gorm.DB

const engine_sql string = "mysql"
const username string = "root"
const password string = "0709"
const database string = "apigo"

func InitializeDataBase() {
	
}
