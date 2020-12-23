package middlewares

import(
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/edwinnduti/postgres-login/models"
	"net/http"
	"os"
	"strconv"
)

// returned response
type Response struct {
	ID		uint8			`json:"id,omitempty"`
	Message string			`json:"message,omitempty"`
	user	models.User 	`json:"user"`
}
// create connection with postgres DB
func CreateConnection() *sql.DB {
	// open connection
	db,err := sql.Open("postgres",os.Getenv("POSTGRES_URI"))
	Check(err)

	// Check connection
	err = db.Ping()
	Check(err)
	fmt.Println("Connected Postgres Successfully!")

	return db
}

// create a user in postgres db
func CreateUser(w http.ResponseWriter,r *http.Request) {
	// set header content to type x-www-form-urlencoded
	// Allow all origin to handle cors
	w.Header().Set("Content-Type","application/x-www-forn-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","POST")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")

	// create new user
	var user models.User

	// decode incoming values to user
	err := json.NewDecoder(r.Body).Decode(user)
	Check(err)

	// insert user
	insertID := InsertUser(user)

	// response to be return
	response := Response{
		ID: insertID,
		Message: "New user created successfully",
		user: user,
	}

	json.NewEncoder(w).Encode(response)
}

// Get a single user
func GetUser(w http.ResponseWriter,r *http.Request) {
	// set header content to type x-www-form-urlencoded
	// Allow all origin to handle cors
	w.Header().Set("Content-Type","application/x-www-forn-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")

	// get id from url
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	Check(err)

	// call get user function
	user,err := RetrieveUser(uint8(id))
	Check(err)

	json.NewEncoder(w).Encode(user)
}

// Get all user
func GetAllUser(w http.ResponseWriter,r *http.Request) {
	// set header content to type x-www-form-urlencoded
	// Allow all origin to handle cors
	w.Header().Set("Content-Type","application/x-www-forn-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")

	// call get user function
	users,err := RetrieveAllUsers()
	Check(err)

	json.NewEncoder(w).Encode(users)
}

// Update single user
func UpdateUser(w http.ResponseWriter,r *http.Request) {
	// set header content to type x-www-form-urlencoded
	// Allow all origin to handle cors
	w.Header().Set("Content-Type","application/x-www-forn-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","PUT")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")

	// get id from url
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	Check(err)

	// create new user
	var user models.User

	// decode incoming values to user
	err = json.NewDecoder(r.Body).Decode(user)
	Check(err)

	// call get user function
	updatedRows := AmendUser(uint8(id),user)

	// message to be returned
	message := fmt.Sprintf("User updated successfully.Total rows/records affected %v",updatedRows)

	// the returned response
	response := Response{
		ID: uint8(id),
		Message: message,
		user: user,
	}

	json.NewEncoder(w).Encode(response)
}

// Delete Single user
func DeleteUser(w http.ResponseWriter,r *http.Request) {
	// set header content to type x-www-form-urlencoded
	// Allow all origin to handle cors
	w.Header().Set("Content-Type","application/x-www-forn-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","DELETE")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")

	// get id from url
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	Check(err)

	// call get user function
	deletedRows := RemoveUser(uint8(id))

	// message to be returned
	message := fmt.Sprintf("User deleted successfully.Total rows/records affected %v",deletedRows)

	// the returned response
	response := Response{
		ID: uint8(id),
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

/*	[*] handler functions [*]	*/

// insert user
func InsertUser(user models.User) uint8 {
	// create postgres db
	db := CreateConnection()
	defer db.Close()

	// sql insert query
	sqlInsertStatement := `INSERT INTO users (name,age,location) VALUES ($1,$2,$3) RETURNING userid`

	// returned id will be stored here
	var id uint8

	// execute sql statement
	err := db.QueryRow(sqlInsertStatement,user.Name,user.Age,user.Location).Scan(&id)
	Check(err)

	fmt.Println("Inserted a single record",id)
	return id
}

// get single user
func RetrieveUser(id uint8) (models.User,error) {
	// create postgres db
	db := CreateConnection()
	defer db.Close()

	// sql insert query
	sqlGetSingleUserStatement := `SELECT * FROM users WHERE userid=$1`

	// returned user will be stored here
	var user models.User

	// execute sql statement
	row := db.QueryRow(sqlGetSingleUserStatement,id)

	// unmarshal the row object to user model
	err := row.Scan(&user.ID,&user.Name,&user.Age,&user.Location)

	switch err {
	case sql.ErrNoRows:
		log.Fatal("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("unable to scan the row :%v", err)
	}

	// return empty user and error
	return user,nil
}

// get all users
func RetrieveAllUsers() ([]models.User,error) {
	// create postgres db
	db := CreateConnection()
	defer db.Close()

	// sql insert query
	sqlGetAllUsersStatement := `SELECT * FROM users`

	// returned users will be stored here
	var users []models.User

	// execute sql statement
	rows,err := db.Query(sqlGetAllUsersStatement)
	Check(err)

	defer rows.Close()

	// unmarshal the row object to user model
	for rows.Next(){
		// empty user
		var user models.User

		// deposit each row object to empty user model
		err = rows.Scan(&user.ID,&user.Name,&user.Age,&user.Location)
		Check(err)

		// add user to users slice
		users = append(users,user)
	}

	// return empty user and error
	return users,nil
}

// update user
func AmendUser(id uint8,user models.User) (uint8) {
	// create postgres db
	db := CreateConnection()
	defer db.Close()

	// sql update query
	sqlUpdateUserStatement := `UPDATE users SET name=$2,age=$3,location=$4 WHERE userid=$1`

	// execute sql statement
	res,err := db.Exec(sqlUpdateUserStatement,id,user.Name,user.Location,user.Age)
	Check(err)

	// returned affected rows
	affectedRows,err := res.RowsAffected()
	Check(err)

	return uint8(affectedRows)
}

// delete user
func RemoveUser(id uint8) (uint8) {
	// create postgres db
	db := CreateConnection()
	defer db.Close()

	// sql delete query
	sqlDeleteUserStatement := `DELETE FROM users WHERE userid=$1`

	// execute sql statement
	res,err := db.Exec(sqlDeleteUserStatement,id)
	Check(err)

	// returned affected rows
	affectedRows,err := res.RowsAffected()
	Check(err)

	return uint8(affectedRows)
}

// handle errors
func Check(err error){
	if err != nil{
		log.Print(err)
	}
}

