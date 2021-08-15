//the package is a test for gorm and pkg/errors
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)
//the data for my database
const driverName string = "mysql"
const dataSourceName string = "root:Tangruilin123@tcp(127.0.0.1:3306)/go_error"
//User is a struct for data
type User struct {
	Uid int
	Name string
	Phone string
}
//the db and the user
var db *sql.DB
var user = &User{
	Uid: 0,
	Name: " ",
	Phone: " ",
}

//the main func
func main()  {
//	open the database
	err := openSql(driverName, dataSourceName)
	if err != nil {
		fmt.Printf("the error: \n %+v \n", err)
		return
	}
	err = queryRow()
	if err != nil {
		fmt.Printf("the error: \n %+v \n", err)
		return
	}
	fmt.Println(user.Uid, user.Name, user.Phone)
}


func openSql(driverName, dataSourceName string) error {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return errors.Wrap(err, "the connect error")
	}
	err = db.Ping()
	if err != nil {
		return errors.Wrap(err, "ping error")
	}
	return nil
}

// the func query only a row
func queryRow() error {
	err := db.QueryRow("select uid, name, phone from `user` where uid = ?", 0).Scan(&user.Uid, &user.Name, &user.Phone)
	if err != nil {
		return errors.Wrap(err, "The query error")
	}
	return nil
}