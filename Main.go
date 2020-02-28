package main

import (
  "time"
  "fmt"
  "log"
	"bytes"
  "golang.org/x/crypto/ssh/terminal"
  "syscall"
  "crypto/cipher"
  "crypto/aes"
  "golang.org/x/crypto/sha3"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func hash(data []byte) []byte {
  hashArray := sha3.Sum256(data)
  return hashArray[:]
}

func decrypt(data, pass []byte) []byte {
	aesBlock, _ := aes.NewCipher(hash(pass))
	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	nonce, cipherText := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}

func main() {
// we ask for the code here.
    fmt.Print("Enter the code: ")
    byteCode, err := terminal.ReadPassword(int(syscall.Stdin))
    if err == nil {
        fmt.Println("\n" + string(byteCode))
    } else {
        fmt.Println("\ncodoo kiri shod")
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
    rows, err := db.Query("select order from users where id = ?", string(hash(byteCode)))
    if err != nil {
	    log.Fatal(err)
    }
    for rows.Next() {
	    err := rows.Scan(&encryptedOrder)
	    if err != nil {
		    log.Fatal(err)
	    }
	    log.Println(encryptedOrder)
    }
	err = rows.Err()
    if err != nil {
	    log.Fatal(err)
    }
    rows.Close()
// we decrypt the encryptedOrder to plainOrder here.
	plainOrder := decrypt([]byte(encryptedOrder), append([]byte("Schenker"), byteCode...))
// we read the plainOrder and execute it here.
	orderSlice := bytes.Split(plainOrder, []byte{'$'})
	if len(orderSlice) % 2 == 1 {
		fmt.Println("orderoo kirie")
	}
	for i := 0; i < len(orderSlice); i += 2 {
		if bytes.Equal(orderSlice[i], []byte{'S', 'H', 'O', 'W'}) {
			fmt.Println(string(orderSlice[i+1]))
		}else if bytes.Equal(orderSlice[i], []byte{'F', 'I', 'L', 'E'}) {
      time.Sleep(time.Millisecond)
    }
  }
}
