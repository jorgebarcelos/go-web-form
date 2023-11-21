package controllers

import "web-form/db"


type Subscription struct {
	Name string
	Email string
}

func Create(name string, email string) error {
	s := Subscription{name, email}

	return db.Insert("newsletter", s)
}