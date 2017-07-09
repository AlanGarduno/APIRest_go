package main 

import ("log" 
		"net/http" 
		"github.com/gorilla/mux")

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/user/",GetUser).Methods("GET")
    log.Println("Este servidor esta corriendo")
    log.Fatal(http.ListenAndServe(":8000",r))

}

func GetUser(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Mustangs\n"))
}
