# Auth SKM Service

This is the project for Jody Almaida and Turasman undergraduated thesis

I use existing libs :

>Chi Router
>Ozzo Validation, for input request validation
>Godotenv, for env loader
>Gorm, for ORM

# For setup after cloning the repo:

cd survey-skm go mod tidy

# to do a unit test :

go to the package you want to testing then run a command "go test" you can see the coverage testing in each package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package

# for db table :

in folder db, there is a .sql file with the create table command and insert command. I use postgresql for this case. you can run the command in your sql editor page.
the endpoint

# postman link for the endpoint test

https://www.getpostman.com/collections/d385aa0f4dd0c8daa1dd