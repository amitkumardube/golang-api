package update_resource

import "testing"

func TestUpdate_resource(t *testing.T){
    // test for a valid case
    new_string , _ := Update_resource("Welcome to Terraform" , "Terraform" , "API Application")
    if new_string != "Welcome to API Application" {
        t.Errorf("failed, expected  %v , got %v ","Welcome to API Application" , new_string)
    }else{
        t.Logf("success, expected  %v , got %v ","Welcome to API Application" , new_string)
    }

    // test for empty string values
    new_string , _ = Update_resource("Welcome to Terraform" , "" , "")
    if new_string != "Welcome to Terraform" {
        t.Errorf("failed, expected  %v , got %v ","Welcome to Terraform" , new_string)
    }else{
        t.Logf("success, expected  %v , got %v ","Welcome to Terraform" , new_string)
    }

    // test for one empty string value
    new_string , _ = Update_resource("Welcome to Terraform" , "API Application" , "")
    if new_string != "Welcome to Terraform" {
        t.Errorf("failed, expected  %v , got %v ","Welcome to Terraform" , new_string)
    }else{
        t.Logf("success, expected  %v , got %v ","Welcome to Terraform" , new_string)
    }

    // test for one empty string value
    new_string , _ = Update_resource("Welcome to Terraform" , "" , "API Application")
    if new_string != "Welcome to Terraform" {
        t.Errorf("failed, expected  %v , got %v ","Welcome to Terraform" , new_string)
    }else{
        t.Logf("success, expected  %v , got %v ","Welcome to Terraform" , new_string)
    }
}