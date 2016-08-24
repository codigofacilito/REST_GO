package structures

type User struct{
	Id 					int 		`json:"id"`
	Username 		string 	`json:"username"`
	First_Name 	string	`json:"first_name"`
	Last_Name 	string	`json:"last_name"`
}

type Reponse struct{
	Status 			string 	`json:"status"`
	Data 				User 		`json:"data"`
	Message 		string 	`json:"message"`
}
