package Conexion

import( "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "../structures")

var connection *gorm.DB

const engine_sql string = "mysql"
const username string = "root"
const password string = "0709"
const database string = "apigo"

func InitializeDataBase() {
	connection = ConnectORM(CreateString() )
	log.Println("La conexion con la base de datos fue exitosa")
}

func CloseConnection() {
	connection.Close()
	log.Println("La conexion con la base de datos fue cerrada")
}

func ConnectORM(stringConnection string) *gorm.DB {
	connection,er := gorm.Open(engine_sql,stringConnection)
	if er != nil{
		log.Fatal(er)
		return nil
	}
	return connection
}

func GetUser(id string) structures.User {
	user := structures.User{}
	connection.Where("id = ?",id).First(&user)
	return user
}

func CreateString() string {
	return username+":"+password+"@/"+database
}

func CreateUser(user structures.User)structures.User {
	connection.Create(&user)
	return user
	
}