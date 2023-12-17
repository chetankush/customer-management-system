![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

# Customer Management System API using Gofr Framework

API for a customer management system using Gofr with CRUD operations.

## Getting Started

1. **Download the zip file** by clicking the "Code" button on the upper left side of the repository home page and extract the files.

2. **Run the project:**


    ```Go
    go run main.go
    ```

3. **To run tests:**

    ```Go
    go test
    ```

4. **To see the coverage:**

    ```Go
    go test -v --cover
    ```



**step by step instructions**

### Prerequisites

- Install Go on your system.
- Install Gofr: `go get gofr.dev`.
- Install Docker on your system.

### Install MySQL on Docker

1. **Install Docker on your system.**
2. **Refer to the Gofr documentation for MySQL installation and connection:** [Gofr MySQL Documentation](https://gofr.dev/docs/v1/quick-start/connecting-mysql).

3. **Use these commands to access the MySQL database on Docker:**

    ```bash
    docker exec -it gofr-mysql bash
    ```

    ```bash
    mysql -u root -proot123
    ```

    Now you can access MySQL.

## To Run the Project

Use this command:

   ```Go
    go run main.go
   ```


to see the data go to localhost:3000/customer

to see data by id go to localhost:3000/customer/2

use postman to make post, update and delete request 

Postman collection - https://documenter.getpostman.com/view/31714271/2s9Ykn92Fy


## To Run tests

 run these commands in root directory ->
 
   ```Go
    go test
   ```
 To see coverage -> 
 
   ```Go
    go test -v --cover
   ```

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



