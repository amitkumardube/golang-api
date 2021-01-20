package update_resource

import (
    "fmt"
    "strings"
    "errors"
)

func Update_resource(old_string string, old_value string, new_value string) (string , error){
    if old_value == "" || new_value == "" {
        return old_string , errors.New("Error : please pass parameters old_string and new_string in the json body")
    }
    fmt.Println("Replacing the string : "+old_value+" with string : "+new_value)
    return strings.Replace(old_string,old_value,new_value,-1) , nil
}