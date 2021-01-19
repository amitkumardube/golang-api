package update_resource

import (
    "fmt"
    "strings"
)

func Update_resource(old_string string, old_value string, new_value string) string{
    fmt.Println("Replacing the string : "+old_value+" with string : "+new_value)
    strings.Replace(old_string,old_value,new_value,-1)
    return old_string
}