package entity

import "time"

type User struct {
	ID         int
	RoleID     string
	RoleName   string
	Name       string
	Email      string
	Password   string
	LastAccess time.Time
}

type Role struct {
	ID   string
	Name string
}

type RoleRight struct {
	ID      string
	RoleID  string
	Section string
	Route   string
	RCreate bool
	RRead   bool
	RUpdate bool
	RDelete bool
}
