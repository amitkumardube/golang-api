package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
    "github.com/amitkumardube/CRUD/delete_resource"
    "github.com/amitkumardube/CRUD/update_resource"
    "strings"
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
    if final_output["terraform"] == "" {
        return "Hello Terraform"
    }else{
        return final_output["terraform"]
    }

}

func set_final_output(new_string string)  {
    final_output["terraform"] = new_string
}

func write_to_rw(w http.ResponseWriter){
    output := return_final_output()
    fmt.Fprintf(w,output)
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
    new_value := strings.TrimSpace(man.New_string) // this is the value that will replace the old value
    if new_value == "" {
        fmt.Fprintf(w, "please pass parameter new_string in the json body to create a new string")
        return
    }
    set_final_output(new_value)
    // writing the value back to response writer
    write_to_rw(w)
}

func read(w http.ResponseWriter, req *http.Request){
    log.Println("Calling the read function")
    write_to_rw(w)
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
    old_value := strings.TrimSpace(man.Old_string)   // this is the value to be replaced
    new_value := strings.TrimSpace(man.New_string)   // this is the value that will replace the old value
    old_string := return_final_output()   // this is the old string
    new_string , err := update_resource.Update_resource(old_string,old_value,new_value)
    if err != nil {
        fmt.Fprintf(w,err.Error())
        return
    }
    // let's set the value of the map to replace value
    set_final_output(new_string)
    // writing the value back to response writer
    write_to_rw(w)
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
    old_value := strings.TrimSpace(man.Old_string)   // this is the value to be deleted
    old_string := return_final_output()   // this is the old string
    new_string , err := delete_resource.Delete_resource(old_string,old_value)
    if err != nil {
        fmt.Fprintf(w,err.Error())
        return
    }
    // let's set the value of the map to replace value
    set_final_output(new_string)
    // writing the value back to response writer
    write_to_rw(w)
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