package main

import (
    "fmt"
    "golang.org/x/crypto/ssh/terminal"
    "syscall"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Print("Enter Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err == nil {
        fmt.Println(bytePassword)
    } else {
        fmt.Println("kiri shod")
    }
}
