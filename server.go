package main

import(
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type User struct {
  Username string `json:"username"`
  First_Name string `json:"first_name"`
  Last_Name string `json:"last_name"`
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/user/", GetUser).Methods("GET")
  log.Println("El servidor se encuentra en el puerto 8000")

  log.Fatal(http.ListenAndServe(":8000", r))
}


func GetUser(w http.ResponseWriter, r* http.Request) {
  user := User{"Gerardo Cetzal", "Gerardo", "Cetzal"}
  json.NewEncoder(w).Encode(user)
}
