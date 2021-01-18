package main

import (
    "fmt"
    "net/http"
    )

func create_resource(w http.ResponseWriter, req *http.Request){

}

func read_resource(w http.ResponseWriter, req *http.Request){
    fmt.Fprintf(w,"Hello World")
}

func update_resource(w http.ResponseWriter, req *http.Request){

}

func delete_resource(w http.ResponseWriter, req *http.Request){

}

func main(){

    // using the golang http package to implement mapping function for each type of API call
    // they are corresponding to the CRUD operations
    http.HandleFunc("/resource/create",create_resource)
    http.HandleFunc("/resource/read",read_resource)
    http.HandleFunc("/resource/update",update_resource)
    http.HandleFunc("/resource/delete",delete_resource)

    http.ListenAndServe(":8000", nil)
}