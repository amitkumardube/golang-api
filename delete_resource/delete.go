package delete_resource

import (
    "fmt"
    "strings"
)

func Delete_resource(old_string string, old_value string) string{
    fmt.Println("Deleting the value : "+old_value+" from string : "+old_string)
    strings.Replace(old_string,old_value,"",-1)
    return old_string
}