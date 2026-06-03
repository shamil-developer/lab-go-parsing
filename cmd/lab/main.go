package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	ID       int      `json:"id"`
	Email    string   `json:"email"`
	FullName FullName `json:"fullname"`
}

type FullName struct {
	FirstName string
	LastName  string
}

func (f *FullName) UnmarshalJSON(data []byte) error {
	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	parts := strings.Split(value, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if strings.HasPrefix(part, "firstName:") {
			f.FirstName = strings.TrimSpace(
				strings.TrimPrefix(part, "firstName:"),
			)
		}

		if strings.HasPrefix(part, "lastName:") {
			f.LastName = strings.TrimSpace(
				strings.TrimPrefix(part, "lastName:"),
			)
		}
	}

	return nil
}

func main() {
	data := []byte(`
{
	"id": 123,
	"email": "alibekov@example.com",
	"fullname": "lastName: alibekov, firstName: shamil"
}
`)

	var user User

	if err := json.Unmarshal(data, &user); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

	fmt.Println("FirstName:", user.FullName.FirstName)
	fmt.Println("LastName :", user.FullName.LastName)
}
