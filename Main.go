package main

import (
    "fmt"
    "golang.org/x/crypto/ssh/terminal"
    "syscall"
    "golang.org/x/crypto/sha3"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
// we ask for the code here.
    fmt.Print("Enter the code: ")
    byteCode, err := terminal.ReadPassword(int(syscall.Stdin))
    if err == nil {
        fmt.Println(byteCode)
    } else {
        fmt.Println("codoo kiri shod")
    }
// we open a connection to our DB here to retrieve the data in the next step.
    db, err := sql.Open("mysql","Schenker:schenker@tcp(127.0.0.1:3306)/Schenker")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
    err = db.Ping()
    if err != nil {
	    log.Fatal(err)
    }
// we search for the given code and try to retrieve the order.
    var decryptedOrder string
    rows, err := db.Query("select order from users where id = ?", string(sha3.Sum256(byteCode)))
    if err != nil {
	    log.Fatal(err)
    }
    for rows.Next() {
	    err := rows.Scan(&decryptedOrder)
	    if err != nil {
		    log.Fatal(err)
	    }
	    log.Println(decryptedOrder)
    }
    err = rows.Err()
    if err != nil {
	    log.Fatal(err)
    }
    rows.Close()
// we decrypt the decryptedOrder here.
    
}
