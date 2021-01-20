package delete_resource

import "testing"

func TestDelete_resource(t *testing.T){
    // test for a valid case
    new_string , _ := Delete_resource("Welcome to Terraform" , "Terraform")
    if new_string != "Welcome to " {
        t.Errorf("failed, expected  %v , got %v ","Welcome to " , new_string)
    }else{
        t.Logf("success, expected  %v , got %v ","Welcome to " , new_string)
    }

    // test for empty string values
    new_string , _ = Delete_resource("Welcome to Terraform" , "")
    if new_string != "Welcome to Terraform" {
        t.Errorf("failed, expected  %v , got %v ","Welcome to Terraform" , new_string)
    }else{
        t.Logf("success, expected  %v , got %v ","Welcome to Terraform" , new_string)
    }

}