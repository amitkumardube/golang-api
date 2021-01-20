# golang - api

golang API to perform the CRUD operations

- C : create
- R : read
- U : update
- D : delete

<b> What these APIs do? </b>
- The APIs perform CRUD operations on the string
- To keep it lightweight project, the APIs don't connect to database 

<b> How to use this project? </b>
- install go binary and configure GOROOT env variable as per the location of GO binary
- create a directory structure dir_path/src. Where dir_path should be set as GOPATH env variable
- clone this code under src repo
- run go run main.go and the application will start on port 8000

<b> List of APIs </b>
- http://localhost:8000/resource/create - Create a new string and string the set value
- http://localhost:8000/resource/read - read the string value
- http://localhost:8000/resource/update - update a part of string with new value supplied
- http://localhost:8000/resource/delete - delete a part of string

<b> How to use automated testing </b>
- The go testing package supports automated testing
- In order to run test testcases , please run go test ./... 
- This will trigger test cases in all the packages ( root and sub packages )