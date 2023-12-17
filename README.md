
 ![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

# customer management system api using gofr framework

API for customer management system using gofr with crud operations

## Getting Started

Download the zip by clicking code button on the upper left side of repository home page and extract the file 

Run this command to run the project -> 

    go run main.go 
    

To run the tests -> 

    go test 

To see the coverage -> 
 
    go test -v --cover 

### Prerequisites

install Go on your system

install Gofr -> go get gofr.dev

install docker on your system 

### Install mysql on docker

install docker on your system then refer to gofr documentation for mysql intallation and connection -> https://gofr.dev/docs/v1/quick-start/connecting-mysql

create your own table and db on mysql - name it as you want

**use these commands to access the mysql database on docker ->**

    docker exec -it gofr-mysql bash

bash-4.4# mysql -u root -proot123 dont add space after -p write it like this -> -proot123

    mysql -u root -p<password>
    
Now u can acces the mysql 


## To Run project
 
use this command 
    
    go run main.go

to see the data go to localhost:3000/customer

to see data by id go to localhost:3000/customer/2

use postman to make post, update and delete request 

Postman collection - https://documenter.getpostman.com/view/31714271/2s9Ykn92Fy


## To Run tests

 for testing -> install sql-mock using this command 
 
    go get gopkg.in/DATA-DOG/go-sqlmock.v1 

 then run these commands in root directory ->
 
    go test
 
 To see coverage -> 

    go test -v --cover


## Built With

* Gofr - The go lang framework
* GORM - The Object-Relational Mapping (ORM) framework, acts as a bridge between Go objects and relational databases.
* sqlmock - used to simulate any sql driver behavior in tests, without needing a real database connection. 
* mockgen - for mocking the datastore layer


## Versioning

go version go1.21.4 windows/amd64

Docker version 24.0.7, build afdd53b

mysql:8.0.30 IN USE

## Diagrams

Use Case Diagram

![usecase](https://github.com/chetankush/customer-management-system/assets/78559285/55686791-5ca6-416b-9b86-faaeca89622f)


Sequence Diagram
![sequence](https://github.com/chetankush/customer-management-system/assets/78559285/084a0d05-523a-4b05-96fd-b7b30563a12b)


## Author

* **Chetan Kushwah** 



