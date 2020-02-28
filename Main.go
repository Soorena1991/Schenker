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
    var encryptedOrder string
    rows, err := db.Query("select order from users where id = ?", string(sha3.Sum256(byteCode)))
    if err != nil {
	    log.Fatal(err)
    }
    for rows.Next() {
	    err := rows.Scan(&encryptedOrder)
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
// we decrypt the encryptedOrder to plainOrder here.
	aesBlock, _ = NewCipher(sha3.Sum256(append(byte["Schenker"], byteCode...)))
	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	nonce, ciphertext := []byte(encryptedOrder)[:gcm.NonceSize()], data[gcm.NonceSize():]
	plainOrder, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
// we read the plainOrder and execute it here.
	
}
