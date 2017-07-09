package main 

import ("log" 
		"net/http" 
		"encoding/json"
		"github.com/gorilla/mux")

type User struct {
	UserName string  `json:UserName`
	FirstName string `json:FirstName`
	LastName string `json:LastName`
}

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}",GetUser).Methods("GET")
    log.Println("Este servidor esta corriendo")
    log.Fatal(http.ListenAndServe(":8000",r))

}

func GetUser(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	user_id := vars['id']
	user := User{"Alan","Test_1", "Test_2"}
	json.NewEncoder(w).Encode(user) 
}
