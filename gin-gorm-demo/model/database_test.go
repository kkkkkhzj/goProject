package model

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_DATABASE_FILE = "test.db"
)

func TestInsertUserByHand(t *testing.T) {
	db, err := SetupDatabase(TEST_DATABASE_FILE)
	assert.Nil(t, err)
	defer os.Remove(TEST_DATABASE_FILE)

	db.engine.Create(&User{Username: "alex", Password: "test", Email: "alex@alex.com"})

	// Read
	var aUser User
	db.engine.First(&aUser, 1) // find product with integer primary key

	log.Println(aUser)

	t.Fatal("xxx")
}

func TestInsertAndQueryUsers(t *testing.T) {
	db, err := SetupDatabase(TEST_DATABASE_FILE)
	assert.Nil(t, err)
	defer os.Remove(TEST_DATABASE_FILE)
	err = db.InsertAUser("test", "test", "test@test.com")
	assert.Nil(t, err)
	users, err := db.QueryAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, "test", users[0].Username, "username not equal")
}
