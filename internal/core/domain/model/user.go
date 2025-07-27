package domain

import (
	domain_builder "github.com/alfariiizi/vandor/internal/core/domain/builder"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/utils"
)

type User struct {
	*db.User
	client *db.Client
}

func NewUserDomain(client *db.Client) domain_builder.Domain[*db.User, *User] {
	return domain_builder.NewDomain(
		func(e *db.User, c *db.Client) *User {
			return &User{
				User:   e,
				client: c,
			}
		}, client)
}

func (u *User) CanLogin() bool {
	return u.Email != ""
}

func (u *User) IsPasswordMatches(password string) bool {
	// passwordHash, err := utils.HashPassword(password)
	// if err != nil {
	// 	return false
	// }
	// log.Printf("Comparing password hash: %s with user password hash: %s", *passwordHash, u.PasswordHash)
	return utils.VerifyPassword(password, u.PasswordHash)
}
