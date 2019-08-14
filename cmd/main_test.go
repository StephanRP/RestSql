package main

import (
	"RestSQL/pkg/config"
	"database/sql"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	db, err := sql.Open(config.SQL, config.EndPoint)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}
	if sqlRes, err := db.Exec("SELECT 1 FROM members LIMIT 1"); err != nil {
		t.Errorf("Expected %s, got %s", sqlRes, err)
	}

	/*sqlRes, err := db.Exec("SELECT 1 FROM members LIMIT 1")
	assert.Equal(t, err != nil, sqlRes)*/
}
