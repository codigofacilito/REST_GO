package connect

import
(
	"log"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "../structures"
)

var connection *gorm.DB
const engine_slq string = "mysql"

const username string = "root"
const password string = ""
const database string = "taller"

func InitializeDataBase(){
	connection = ConnectORM( GetConnectionString() )
	log.Println("Conexion a la base de datos establecida.")
}

func ConnectORM(stringConnection string) *gorm.DB{
	connection, err := gorm.Open(engine_slq, stringConnection)
	if err != nil{
		log.Println(err)
		return nil;
	}
	return connection
}

func GetConnectionString() string{
	return username + ":"+ password + "@/" + database 
}

func CloseConnection(){
	connection.Close()
	log.Println("La conexion se ha cerrado.")
}

func GetUser(id string) structures.User{
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	return user
}

func CreateUser(user structures.User) structures.User {
	connection.Create(&user)//Se asigna el id
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


func DeleteUser(id string){
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	connection.Delete(&user)
}




