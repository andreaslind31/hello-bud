package controller

import (
	context "context"
)

type Controller struct {
	// Dependencies...
}

// Post struct
type Post struct {
	// Fields...
}

// Index of posts
// GET /
func (c *Controller) Index(ctx context.Context) (posts []*Post, err error) {
	return []*Post{}, nil
}

// Show post
// GET /:id
func (c *Controller) Show(ctx context.Context, id int) (post *Post, err error) {
	return &Post{}, nil
}
