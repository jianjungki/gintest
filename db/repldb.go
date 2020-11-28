package db


import (
	"context"
	"fmt"
	"net/url"
  "github.com/replit/database-go"
)



type ReplDB struct {}

func (m *ReplDB) Query(key string) (string,error) {
  return database.Get(key)
}

func (m *ReplDB) Delete(key string) error {
  return database.Delete(key)
}

func (m *ReplDB) Update(key , value string) error {
	return database.Set(key, value)
}

func (m *ReplDB) Insert(key, value string) *sql.Row {
	return database.Set(key, value)
}