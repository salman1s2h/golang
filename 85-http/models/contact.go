package models

import (
	"encoding/json"
	"errors"
)

func init() {
	println("init is in models-1")
}

func init() {
	println("init is in models-2")
}

type Contact struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	Lastmodified int64  `json:"last_modified"`
}

func (c *Contact) Validate() error {
	if c.Id == 0 {
		return errors.New("invalid id field")
	}
	if c.Name == "" {
		return errors.New("invalid name field")
	}

	if c.Email == "" {
		return errors.New("invalid email field")
	}
	return nil
}

func (c *Contact) ToBytes() ([]byte, error) {
	return json.Marshal(c)
}
