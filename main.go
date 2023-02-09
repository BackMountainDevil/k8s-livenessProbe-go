package main

import (
   "fmt"
   "net/http"
)

func main() {
   fmt.Println("service start")

   http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
      fmt.Fprint(writer,"Hello Docker")
   })
   http.ListenAndServe(":8888",nil)

}

