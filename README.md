# CoderSchool Blockchain Prework!
![B](http://www.coderschool.vn/assets/courses/banner-square-blockchain-41db44e92f0c89b8b53cb0b5b7dc076bba22f1937db4fc73abe6482b1146ac67.png)

Hey there! Welcome to the Blockchain prework.

You can view instructions for this here: http://www.coderschool.vn/prework/blockchain.

* Finish and passed all test cases

* Optional features: 

1. Make your code into a library that can be imported into another project. 

- I published github.com/Thanghand/blockchain to github with https://github.com/Thanghand/blockchain

- In folder /restapi/rest.go, It is using blockchain library. To run it, you have to run with command:  go get github.com/Thanghand/blockchain (it will donwload into $GOPATH/src/github)

2. Create api with github.com/gorilla/mux 

* How to run it ? :  

- cd /restapi 

- go build

- ./restapi

* Get last block chain api: 
    
- Method GET: http://localhost:8069/blockchain

- Header: Content-Type/application/json

- Response example: 

{
    "message": "Get Last Block successfully",
    "hashBlock": "/sbJ2x/ONwGHjcH+LPZ9+NCsUbTeeogg1mbjGvwg05I="
}

* Add new block api: 

- Method POST: http://localhost:8069/blockchain 

- Header: Content-Type/application/json

- Request: 

{
	"text": "Hi My name is THang Cao",
	"oldHash": "/sbJ2x/ONwGHjcH+LPZ9+NCsUbTeeogg1mbjGvwg05I="
}

- Response success: 

{
    "message": "Add new block successfully with text : Hi My name is THang Cao",
    "hashBlock": "+tYhzzgzXhIT9h+CzVVw9EmAIkLZYmAI1okR9noGwOA="
}

- Response error: When we use invalid hash

{
    "message": "Add new block error"
}