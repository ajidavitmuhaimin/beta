package main

import "fmt"
import "net/http"



func main(){

   http.HandleFunc("/send",sendHandler)
   http.HandleFunc("/conn",connHandler)
   http.ListenAndServe(":1395",nil)

}