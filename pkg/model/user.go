package model

import (
	"fmt"
	"time"
)

// User represents the the security principal for this application.
// swagger:model
// title: User
// description: User object
type User struct {
	// User identifier, same as Github ID
	// example: 5869473
	ID int64 `pg:",pk" json:"id"`

	// Unique username, same as Github login
	// example: spalt08
	Username string `pg:",unique" json:"username"`

	// User full name
	// example: Konstantin Darutkin
	Name string `json:"name"`

	// Link to avatar picture
	// example: https://avatars3.githubusercontent.com/u/5869473?s=60
	AvatarURL *string `json:"avatar_url"`

	Company   *string   `json:"-"`
	Email     *string   `pg:",unique" json:"-"`
	CreatedAt time.Time `pg:",type:timestamp(0),default:now()" json:"-"`
}

func (m *User) String() string {
	return fmt.Sprintf("User <id: %d, username: %s>", m.ID, m.Username)
}
