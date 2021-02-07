package main

import (
	arango "github.com/joselitofilho/gondm/driver/arango"
	gondm "github.com/joselitofilho/gondm/internal"
)

func main() {
	db, err := gondm.Open(arango.Open(&arango.Config{}), &gondm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	type User struct {
		gondm.Model
		Name string
		Age  int
	}

	// Create
	db.Create(&User{Name: "111", Age: 80})

	// Read
	var user User
	db.First(&user, &gondm.Options{}) // find first user with options

	// Update - update all user's age to 30
	db.Update(&User{}, "Age", 30)
	// Update - update user's age to 30
	db.Update(&user, "Age", 30)
	// Update - update multiple fields
	db.Updates(&user, User{Age: 85, Name: "222"}) // non-zero fields
	db.Updates(&user, map[string]interface{}{"Age": 85, "Id": "222"})
	db.Updates(&User{}, User{Age: 85, Name: "222"}) // all

	// Delete - delete user
	db.Delete(&user, &gondm.Options{})
}
