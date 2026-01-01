package main


import (
	"net/http"
	"fmt"
)

func main(){

	http.HandleFunc("/",
	func(w http.ResponseWriter, r *http.Request) {
		user :=r.UserAgent()
		fmt.Printf(user)
	})

	http.ListenAndServe(":3001",nil)
}