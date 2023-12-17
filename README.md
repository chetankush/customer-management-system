
 ![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

# customer management system api using gofr framework

API for customer management system using gofr with crud operations

postman collection - https://documenter.getpostman.com/view/31714271/2s9Ykn92Fy

## Diagrams are in the end

## Getting Started

download the zip by clicking code button on the upper left side of repository home page and extract the file to use it 

run this command -> go run main.go (use it in project root directory terminal on your local machine for development and testing purposes.)

to run the tests -> go test (use it in project root directory terminal on your local machine for testing purposes.)

to see the coverage -> go test -v --cover (to see the coverage of this project)



### Prerequisites

install Go on your system

install Gofr -> go get gofr.dev

install docker on your system then refer this gofr documentation for mysql intallation -> https://gofr.dev/docs/v1/quick-start/connecting-mysql
  

### Install and Run project

A step by step series of examples that tell you how to get a development env running

download the project zip and extract it on the desktop to access easily

install docker on your system then refer to gofr documentation for mysql intallation and connection
  
create your own table and db on mysql - name it as you want

**use these commands to access the mysql on docker ->**

docker exec -it gofr-mysql bash

bash-4.4# mysql -u root -p<your mysql password for docker image> dont add space after -p write it like this -> -proot123

show databases (to see all databases)

use <your database name> that you created using gofr-mysql documentation
 
then run **go run main.go**

to see the data go to localhost:3000/customer

to see data by id go to localhost:3000/customer/2


use postman to make post, update and delete request 


## Running the tests

 for testing -> install sql-mock if not present 

 then run these commands in root directory ->**go test**
                           To see coverage ->**go test -v --cover**


## Built With

* Gofr - The go lang framework
* GORM - The Object-Relational Mapping (ORM) framework, acts as a bridge between Go objects and relational databases.
* sqlmock - go get gopkg.in/DATA-DOG/go-sqlmock.v1
* mockgen - for mocking the datastore layer


## Versioning

go version go1.21.4 windows/amd64

Docker version 24.0.7, build afdd53b

mysql:8.0.30 IN USE

## Diagrams

Use Case Diagram
![WhatsApp Image 2023-12-17 at 14 21 19](https://github.com/chetankush/customer-management-system/assets/78559285/932fddf5-6af8-4b33-a7af-e3279bdc136e)


Sequence Diagram
![WhatsApp Image 2023-12-17 at 14 21 37](https://github.com/chetankush/customer-management-system/assets/78559285/5c8dd528-5fbc-4b15-9d85-19749f72aaf7)


## Author

* **Chetan Kushwah** 

