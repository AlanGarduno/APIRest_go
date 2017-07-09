package main 

import ("log" 
		"net/http" 
		"encoding/json"
		"github.com/gorilla/mux"
		"./Conexion"
		//"./structures"
		)

func main(){
	Conexion.InitializeDataBase()
	defer Conexion.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}",GetUser).Methods("GET")
    log.Println("Este servidor esta corriendo en el puerto 8000")
    log.Fatal(http.ListenAndServe(":8000",r))

}

func GetUser(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	user_id := vars["id"]

	user := Conexion.GetUser(user_id)
	json.NewEncoder(w).Encode(user) 
}
