package main  

import (
	"database/sql"
	"log"
   _"github.com/lib/pq"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "hello"
	DB_NAME     = "postgres"
	DB_HOST     = "localhost:5432"
	SSL_MODE    = "disable"
)

var DSN string = "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + "/" + DB_NAME + "?sslmode=" + SSL_MODE  

var db *sql.DB

type table struct {
    id int64
    name string 
    region string
    numberrange int64 
    list int64 
}   

func tabque(name string) ([]table ,error) {

    var tabl []table

    rows, err := db.Query("SELECT id, name,  region, numberrange, list FROM hhhhhhh WHERE region = $1 " , name)
    if err != nil {
        return nil, fmt.Errorf("hhhhhhh %q: %v", name, err)
    }


    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var tab table
        if err := rows.Scan(&tab.id, &tab.name, &tab.region, &tab.numberrange, &tab.list); err != nil {
            return nil, fmt.Errorf("hhhhhhh %q: %v", name, err)
        }
        tabl = append(tabl, tab)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("hhhhhhh %q: %v", name, err)
    }
    return tabl, nil

}





func main() {
    
    // Get a database handle.
    var err error
    db, err = sql.Open("postgres", DSN)
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")


     
    fmt.Println(tabque("Derbyshire"))
    


}