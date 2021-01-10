package tests

import (
	"math/rand"

	"touchstone-api/external/postgres"
	"touchstone-api/pkg/middleware"
	"touchstone-api/pkg/model"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func createRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func setDefaultValue(data map[string]interface{}, key string, value interface{}) {
	if _, ok := data[key]; ok == false {
		data[key] = value
	}
}

func getStringPointer(data interface{}) *string {
	if data == nil {
		return nil
	} else {
		var strVal = data.(string)
		return &strVal
	}
}

// CreateUser in database for tests
func CreateUser(db postgres.Connection, initialData map[string]interface{}) (*model.User, string) {
	setDefaultValue(initialData, "ID", int64(rand.Intn(1000000)))
	setDefaultValue(initialData, "Username", createRandomString(15))
	setDefaultValue(initialData, "Name", createRandomString(15))
	setDefaultValue(initialData, "Email", nil)
	setDefaultValue(initialData, "AvatarURL", nil)

	user := &model.User{
		ID:        initialData["ID"].(int64),
		Username:  initialData["Username"].(string),
		Name:      initialData["Name"].(string),
		Email:     getStringPointer(initialData["Email"]),
		AvatarURL: getStringPointer(initialData["AvatarURL"]),
	}

	_, err := db.
		Model(user).
		Insert()

	if err != nil {
		panic(err)
	}

	token, _ := middleware.CreateToken(user)

	return user, token
}
