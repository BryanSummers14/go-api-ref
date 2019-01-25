package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type User struct {
	Username string
	Password string
	Email    string
}

func main() {
	users := []User{
		{Username: "Jane Doe", Password: "Change me", Email: "janedoe@msn.com"},
		{Username: "Jane Doe", Password: "Change me", Email: "janedoe@msn.com"},
		{Username: "Jane Doe", Password: "Change me", Email: "janedoe@msn.com"},
	}
	//fmt.Println(user)

	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(users)
	io.Copy(os.Stdout, buf)
}
