package connect

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  "../structures"
)

var connection *gorm.DB

const engine_sql string = "mysql"


const username string = "root"
const password string = "zxczxc.123"
const database string = "taller1"

func InitializeDatabase() {
  connection = ConnectORM(CreateString())
  log.Println("la conexion con la base de datos fue exitosa")
}

func CloseConnection() {
  connection.Close()
  log.Println("la conexion con la base de datos fue cerrada")
}

func ConnectORM(stringConnection string) *gorm.DB {
  connection, err := gorm.Open(engine_sql, stringConnection)

  if err != nil {
    log.Fatal(err)
    return nil
  }
  return connection
}

func CreateString() string {
  return username + ":" + password + "@/" + database
}

func GetUser(id string) structures.User {
  user := structures.User{}
  connection.Where("id = ?", id).First(&user)

  return user
}

func CreateUser(user structures.User) structures.User{
  connection.Create(&user) // Se le asigna un id
  return user
}

func UpdateUser(id string, user structures.User) structures.User{
  currentUser := structures.User{}
  connection.Where("id = ?", id).First(&currentUser)

  currentUser.Username = user.Username
  currentUser.First_Name = user.First_Name
  currentUser.Last_Name = user.Last_Name

  connection.Save(&currentUser)

  return currentUser
}

func DeleteUser(id string) {
  user := structures.User{}
  connection.Where("id = ?", id).First(&user)

  connection.Delete(&user)
}
