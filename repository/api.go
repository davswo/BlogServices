package repository

import "errors"

type BlogPost struct {
	BlogId int    `json: blogId`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

type BlogRepository interface {
	InsertBlogPost(bp BlogPost) error
	GetBlogPosts() ([]BlogPost, error)
}

var ErrDuplicateKey = errors.New("Duplicate key")
