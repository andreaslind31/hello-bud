package user

type Controller struct{}

type User struct {
	ID    int
	Name  string
	Email int
}

func (c *Controller) Index() ([]*User, error) {
	var users []*User

	return users, nil
}

func (c *Controller) New() {}

func (c *Controller) Create(name, email string) (*User, error) {
	user := User{}

	return &user, nil
}

func (c *Controller) Show(id int) (*User, error) {
	user := User{}

	return &user, nil
}

// Update a user
// PATCH /users/:id
func (c *Controller) Update(id int, name, email string) error {
	return nil
}

func (c *Controller) Delete(id int) error {
	return nil
}

// Edit user page
// GET /users/:id/edit
func (c *Controller) Edit(id int) (*User, error) {
	user := User{}

	return &user, nil
}
