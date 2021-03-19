package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
  )

func TestResource(t *testing.T){
    terraformOptions := &terraform.Options{
        TerraformDir: "../",
    }

	// defer the destroy part till the CLI is init and resources are created.
    defer terraform.Destroy(t, terraformOptions)
	// this will init terraform and apply the resources
    terraform.InitAndApply(t, terraformOptions)	
}