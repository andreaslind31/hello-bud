package users

import (
	context "context"
)

type Controller struct {
	// Dependencies...
}

// User struct
type User struct {
	// Fields...
}

// Index of users
// GET /users
func (c *Controller) Index(ctx context.Context) (users []*User, err error) {
	return []*User{}, nil
}

// Show user
// GET /users/:id
func (c *Controller) Show(ctx context.Context, id int) (user *User, err error) {
	return &User{}, nil
}
