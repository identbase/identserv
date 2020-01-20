package account

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

/*
User defines a user within the identity server. */
type User struct {
	key   string
	Token string
}

/*
Key implements the Keyer interface so that a User can be stored. */
func (u *User) Key() string {
	return u.key
}

/*
GenerateToken creates a new (access) Token for a User. */
func (u *User) GenerateToken() error {
	t := time.Now()
	h, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s%s", u.key, t)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	hasher := md5.New()
	hasher.Write(h)
	u.Token = hex.EncodeToString(hasher.Sum(nil))
	return nil
}

/*
New creates a new User object. */
func New(id string) *User {
	u := User{
		key: id,
	}

	return &u
}
