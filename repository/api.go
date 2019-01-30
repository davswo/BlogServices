package repository

import "errors"

type BlogPost struct {
	BlogId string `json:"blogId"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

type BlogRepository interface {
	InsertBlogPost(bp BlogPost) error
	GetBlogPosts() ([]BlogPost, error)
}

// ErrDuplicateKey is thrown when there is an attempt to create an order with an BlogId which already is used.
var ErrDuplicateKey = errors.New("Duplicate key")
