package main 

import ("log" 
		"net/http" 
		"encoding/json"
		"github.com/gorilla/mux"
		"./Conexion"
		"./structures"
		)

func main(){
	Conexion.InitializeDataBase()
	defer Conexion.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}",GetUser).Methods("GET")
	r.HandleFunc("/user/new",NewUser).Methods("POST")
	r.HandleFunc("/user/update/{id}",UpdateUser).Methods("PATCH")

    log.Println("Este servidor esta corriendo en el puerto 8000")
    log.Fatal(http.ListenAndServe(":8000",r))

}

func GetUser(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	user_id := vars["id"]

	status := "succes"
	var message string
	user := Conexion.GetUser(user_id)
	if(user.Id <= 0){
		status = "Error"
		message = "User not found"
	} 
	respose := structures.Response{status,user,message}
	json.NewEncoder(w).Encode(respose) 
}

func NewUser(w http.ResponseWriter,r *http.Request) {
	user := GetUserRequest(r)
	respose :=  structures.Response{"suces",Conexion.CreateUser(user),""}
	json.NewEncoder(w).Encode(respose) 
}

func GetUserRequest(r *http.Request) structures.User {
	var user structures.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil{
		log.Fatal(err)
	}
	return user
}

func UpdateUser(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]

	user := GetUserRequest(r)
	
	respose :=  structures.Response{"suces",Conexion.UpdateUser(user_id,user),""}
	json.NewEncoder(w).Encode(respose)
}

