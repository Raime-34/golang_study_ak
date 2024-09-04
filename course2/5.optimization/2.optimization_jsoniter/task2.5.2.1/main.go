package task2_5_2_1

import (
	"encoding/json"

	gofakeit "github.com/brianvoe/gofakeit/v7"
	jsoniter "github.com/json-iterator/go"
	easyjson "github.com/mailru/easyjson"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

func EasyJSON(users []User) {
	for _, user := range users {
		_, _ = easyjson.Marshal(&user)
	}
}

func JSON(users []User) {
	for _, user := range users {
		_, _ = json.Marshal(&user)
	}
}

func JSONiter(users []User) {
	for _, user := range users {
		_, _ = jsoniter.Marshal(&user)
	}
}

func GenerateUSER(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			ID:       i,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 14),
			Age:      gofakeit.Number(18, 100),
			Email:    gofakeit.Email(),
		}
	}

	return users
}
