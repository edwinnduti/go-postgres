# postgres and Golang API.

[![Build Status](https://travis-ci.com/edwinnduti/go-postgres.svg?branch=master)](https://travis-ci.com/edwinnduti/go-postgres)
![License: MIT](https://img.shields.io/badge/Language-Golang-blue.svg)
[![Build GO workflow](https://github.com/edwinnduti/go-postgres/actions/workflows/deploy.yaml/badge.svg?branch=master)](https://github.com/edwinnduti/go-postgres/actions/workflows/deploy.yaml)
![License: MIT](https://img.shields.io/badge/Database-POSTGRESQL-darkblue.svg)


### Requirements:
* Postgresql
* Golang

 ```
 $ git clone https://github.com/edwinnduti/go-postgres.git 
 $ cd go-postgres
 $ go mod download
 $ go run main.go
 ```

Available :

| function              |   path                    |   method  |
|   ----                |   ----                    |   ----    |
| Create user           |   /api			|	POST    |
| Get single user       |   /api/{user_id}			|	GET     |
| Get All users         |   /api/users                  |	GET     |
| Delete single user    |   /api/user/{user_id}		|	DELETE  |
| update single user    |   /api/user/{user_id}		|	UPDATE  |


 Have a nice day!

