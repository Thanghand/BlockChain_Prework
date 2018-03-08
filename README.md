# CoderSchool Blockchain Prework!

#### * Finished and passed all test cases

#### * Optional features: 

## 1. Make your code into a library that can be imported into another project. 

- I published github.com/Thanghand/blockchain to github with https://github.com/Thanghand/blockchain

- In folder /restapi/rest.go, It is using blockchain library. 

- To run it, you have to run with command:  

```
$: go get github.com/Thanghand/blockchain 

```
(it will donwload into $GOPATH/src/github)

## 2. Create api with github.com/gorilla/mux 

# How to run it ? :  

```
$: cd /restapi 

$: go build

$: ./restapi

```

# Get last block chain api: 
    
* Method GET: http://localhost:8069/blockchain

* Header: Content-Type/application/json

* Response example: 

```
{
    "message": "Get Last Block successfully",
    "hashBlock": "/sbJ2x/ONwGHjcH+LPZ9+NCsUbTeeogg1mbjGvwg05I="
}
```

# Add new block api: 

* Method POST: http://localhost:8069/blockchain 

* Header: Content-Type/application/json

* Request: 

```
{
	"text": "Hi My name is THang Cao",
	"oldHash": "/sbJ2x/ONwGHjcH+LPZ9+NCsUbTeeogg1mbjGvwg05I="
}
```

* Response success: 

```
{
    "message": "Add new block successfully with text : Hi My name is THang Cao",
    "hashBlock": "+tYhzzgzXhIT9h+CzVVw9EmAIkLZYmAI1okR9noGwOA="
}
```

* Response error: When we use invalid hash

```
{
    "message": "Add new block error"
}
```

* POST MAN: 
<img width="600" alt="screen shot 2018-03-08 at 8 20 45 pm" src="https://user-images.githubusercontent.com/8011261/37153190-76d7ecf4-230e-11e8-9fc2-9894eed3a733.png">


<img width="600" alt="screen shot 2018-03-08 at 8 21 06 pm" src="https://user-images.githubusercontent.com/8011261/37153205-8194054c-230e-11e8-85b4-e35645cd3eaf.png">


<img width="600" alt="screen shot 2018-03-08 at 8 21 19 pm" src="https://user-images.githubusercontent.com/8011261/37153222-92b87b14-230e-11e8-86a2-a1c7efbc9222.png">

