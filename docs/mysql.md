## Connecting to a MySQL DB
The following is an example of how to connect to a MySQL database using `Go`.
```go
package main

import (
	"fmt"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
    id int
    name string
    createdAt time.Time
}

const db_name string = "dbname"
const db_user string = "dbuser"
const db_pass string = "dbpass"

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", db_user, db_pass, db_name))
	if err != nil {
		panic(err)
	}

	// Initialize the first connection to the database,
  // to see if everything works correctly.
	ping_err := db.Ping()
	if ping_err != nil {
		panic(ping_err)
	}

	var myuser user
	stmt, err := db.Prepare("SELECT user_id, name FROM users")
	if err = stmt.QueryRow().Scan(&myuser.id, &myuser.name); err != nil {
		panic(err)
	}
	fmt.Println(myuser)

}
```
