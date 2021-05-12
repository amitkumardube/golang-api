package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/amitkumardube/CRUD/delete_resource"
	"github.com/amitkumardube/CRUD/update_resource"

	"github.com/go-redis/redis/v8"
	"context"
)

type string_manipulation struct {
	Old_string string `json:"old_string,omitempty"`
	New_string string `json:"new_string,omitempty"`
}

// defining a map which will keep the value for the final string

var final_output map[string]string

// variables for redis
var redis_addr = "10.3.246.134"
var redis_port = "6379"

// creating an instance of redis

func invoke_redis(){
	var ctx = context.Background()
        client := redis.NewClient(&redis.Options{
                Addr: redis_addr+":"+redis_port,
                Password: "",
                DB: 0,
        })
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(pong, err)	
	}
	fmt.Println(pong, err)
}

// creating the init function to initialize the map
func init() {
	final_output = make(map[string]string)
	final_output["terraform"] = ""

	// initializing redis
	invoke_redis()
}

func return_final_output() string {
	if final_output["terraform"] == "" {
		return "Hello Terraform \n"
	} else {
		return final_output["terraform"]
	}

}

// We need to check if end point was called with correct method.
func check_req_method(w http.ResponseWriter, r http.Request, req_type string) {
	req_method := r.Method
	allowed_method := req_type
	if strings.ToLower(req_method) == strings.ToLower(allowed_method) {
		fmt.Sprintf("Call to end point is made with %s method which is correct. Processing Further.", req_type)
	} else {
		set_final_output(fmt.Sprintf("Call to end point is made with wrong Method - %s. Allowed Method is %s", req_method, allowed_method))
		write_to_rw(w)
		return
	}
}

func set_final_output(new_string string) {
	final_output["terraform"] = new_string
}

func write_to_rw(w http.ResponseWriter) {
	output := return_final_output()
	fmt.Fprintf(w, output)
}

// this function will reset the map to empty string

func reset(w http.ResponseWriter, req *http.Request) {
	set_final_output("")
	write_to_rw(w)
}

func create(w http.ResponseWriter, req *http.Request) {
	log.Println("Calling the create function")
	check_req_method(w, *req, "POST")
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

func read(w http.ResponseWriter, req *http.Request) {
	log.Println("Calling the read function")
	write_to_rw(w)
}

func update(w http.ResponseWriter, req *http.Request) {
	log.Println("Calling the update function")
	body := json.NewDecoder(req.Body)
	body.UseNumber()
	var man string_manipulation
	err := body.Decode(&man)
	if err != nil {
		log.Println(err.Error())
		return
	}
	old_value := strings.TrimSpace(man.Old_string) // this is the value to be replaced
	new_value := strings.TrimSpace(man.New_string) // this is the value that will replace the old value
	old_string := return_final_output()            // this is the old string
	new_string, err := update_resource.Update_resource(old_string, old_value, new_value)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	// let's set the value of the map to replace value
	set_final_output(new_string)
	// writing the value back to response writer
	write_to_rw(w)
}

func delete(w http.ResponseWriter, req *http.Request) {
	log.Println("Calling the delete function")
	body := json.NewDecoder(req.Body)
	body.UseNumber()
	var man string_manipulation
	err := body.Decode(&man)
	if err != nil {
		log.Println(err.Error())
		return
	}
	old_value := strings.TrimSpace(man.Old_string) // this is the value to be deleted
	old_string := return_final_output()            // this is the old string
	new_string, err := delete_resource.Delete_resource(old_string, old_value)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	// let's set the value of the map to replace value
	set_final_output(new_string)
	// writing the value back to response writer
	write_to_rw(w)
}

func main() {

	// using the golang http package to implement mapping function for each type of API call
	// they are corresponding to the CRUD operations
	http.HandleFunc("/resource/create", create)
	http.HandleFunc("/resource/read", read)
	http.HandleFunc("/resource/update", update)
	http.HandleFunc("/resource/delete", delete)
	http.HandleFunc("/resource/reset", reset)
	http.HandleFunc("/", read)

	fmt.Println("Launching application on port 8080")
	http.ListenAndServe(":8080", nil)
}
