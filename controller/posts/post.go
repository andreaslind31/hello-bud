package posts

type Controller struct{}

type Post struct {
	ID      int
	Subject string
}

func (c *Controller) Index() ([]*Post, error) {
	var posts []*Post

	return posts, nil
}

func (c *Controller) New() {}

func (c *Controller) Create(name, subject string) (*Post, error) {
	post := Post{}

	return &post, nil
}

func (c *Controller) Show(id int) (*Post, error) {
	post := Post{}

	return &post, nil
}

// Update a Post
// PATCH /Posts/:id
func (c *Controller) Update(id int, subject string) error {
	return nil
}

func (c *Controller) Delete(id int) error {
	return nil
}

// Edit Post page
// GET /Posts/:id/edit
func (c *Controller) Edit(id int) (*Post, error) {
	post := Post{}

	return &post, nil
}
