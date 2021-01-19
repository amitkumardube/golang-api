package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
    "golang-api/update_resource"
    "golang-api/delete_resource"
    )

type string_manipulation struct {
    Old_string  string  `json:"old_string,omitempty"`
    New_string string `json:"new_string,omitempty"`
}

// defining a map which will keep the value for the final string

var final_output map[string]string

// creating the init function to initialize the map
func init(){
    final_output = make(map[string]string)
    final_output["terraform"] = ""
}

func return_final_output() string{
    return final_output["terraform"]
}

func set_final_output(new_string string)  {
    final_output["terraform"] = new_string
}

func create(w http.ResponseWriter, req *http.Request){
    log.Println("Calling the create function")
    body := json.NewDecoder(req.Body)
    body.UseNumber()
    var man string_manipulation
    err := body.Decode(&man)
    if err != nil {
        log.Println(err.Error())
        return
    }
    new_value := man.New_string // this is the value that will replace the old value
    set_final_output(new_value)

}

func read(w http.ResponseWriter, req *http.Request){
    log.Println("Calling the read function")
    output := return_final_output()
    fmt.Fprintf(w, output)
}

func update(w http.ResponseWriter, req *http.Request){
    log.Println("Calling the update function")
    body := json.NewDecoder(req.Body)
    body.UseNumber()
    var man string_manipulation
    err := body.Decode(&man)
    if err != nil {
        log.Println(err.Error())
        return
    }
    old_value := man.Old_string   // this is the value to be replaced
    new_value := man.New_string   // this is the value that will replace the old value
    old_string := return_final_output()   // this is the old string
    new_string := update_resource.Update_resource(old_string,old_value,new_value)
    // let's set the value of the map to replace value
    set_final_output(new_string)
}

func delete(w http.ResponseWriter, req *http.Request){
    log.Println("Calling the delete function")
    body := json.NewDecoder(req.Body)
    body.UseNumber()
    var man string_manipulation
    err := body.Decode(&man)
    if err != nil {
        log.Println(err.Error())
        return
    }
    old_value := man.Old_string   // this is the value to be replaced
    old_string := return_final_output()   // this is the old string
    new_string := delete_resource.Delete_resource(old_string,old_value)
    // let's set the value of the map to replace value
    set_final_output(new_string)
}

func main(){

    // using the golang http package to implement mapping function for each type of API call
    // they are corresponding to the CRUD operations
    http.HandleFunc("/resource/create",create)
    http.HandleFunc("/resource/read",read)
    http.HandleFunc("/resource/update",update)
    http.HandleFunc("/resource/delete",delete)

    fmt.Println("Launching application on port 8000")
    http.ListenAndServe(":8000", nil)
}