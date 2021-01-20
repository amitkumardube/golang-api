package delete_resource

import (
    "fmt"
    "strings"
    "errors"
)

func Delete_resource(old_string string, old_value string) (string,error){
    if old_value == "" {
        return old_string , errors.New("Error : please pass parameter old_string in json body for deletion")
    }
    fmt.Println("Deleting the value : "+old_value+" from string : "+old_string)
    return strings.Replace(old_string,old_value,"",-1) , nil
}