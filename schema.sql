
-- query to create database

CREATE DATABASE test_db;


-- query to create table in test_db database

CREATE TABLE customers (

    id              INT AUTO_INCREMENT PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    phone           VARCHAR(255) NOT NULL,
    address         VARCHAR(255) NOT NULL,
    city            VARCHAR(255) NOT NULL,
    date_of_birth   VARCHAR(255) NOT NULL,
    is_active       BOOLEAN      NOT NULL

);

