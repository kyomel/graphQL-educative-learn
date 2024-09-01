package utils

import (
	"time"

	"github.com/bxcodec/faker/v3"
)

type UserFaker struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Username  string    `json:"username" bson:"username" faker:"username"`
	Email     string    `json:"email" bson:"email" faker:"email"`
	Password  string    `json:"password" bson:"password" faker:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type BlogFaker struct {
	ID        string     `json:"id" bson:"_id, omitempty"`
	Title     string     `json:"title" bson:"title" faker:"title"`
	Content   string     `json:"content" bson:"content" faker:"content"`
	Author    *UserFaker `json:"author" bson:"author"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
}

func CreateFaker[T any]() (T, error) {
	var fakerData *T = new(T)
	err := faker.FakeData(fakerData)
	if err != nil {
		return *fakerData, err
	}

	return *fakerData, nil
}
