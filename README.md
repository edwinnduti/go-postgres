# postgres and Golang API.

### Requirements:
* Postgresql
* Golang

 ```
 $ git clone https://github.com/edwinnduti/go-postgres.git 
 $ cd go-postgres
 $ go run main.go
 ```

Available :

| function              |   path                    |   method  |
|   ----                |   ----                    |   ----    |
| Create user           |   /api/newuser		|	POST    |
| Get single user       |   /api/{user}			|	GET     |
| Get All users         |   /api/users                  |	GET     |
| Delete single user    |   /api/user/{id}		|	DELETE  |
| update single user    |   /api/user/{id}		|	UPDATE  |


 Have a nice day!

