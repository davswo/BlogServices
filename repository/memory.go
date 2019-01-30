package repository

import (
	"fmt"
)

type blogRepositoryMemory struct {
	BlogPost map[string]BlogPost
}

// NewOrderRepositoryMemory is used to instantiate and return the DB implementation of the OrderRepository.
func NewOrderRepositoryMemory() BlogRepository {
	return &blogRepositoryMemory{BlogPost: make(map[string]BlogPost)}
}

func (repository *blogRepositoryMemory) InsertBlogPost(blogPost BlogPost) error {
	id := mapID(blogPost)
	if _, exists := repository.BlogPost[id]; exists {
		return ErrDuplicateKey
	}
	repository.BlogPost[id] = blogPost
	return nil
}

func (repository *blogRepositoryMemory) GetBlogPosts() ([]BlogPost, error) {
	ret := make([]BlogPost, 0, len(repository.BlogPost))
	for _, blogPost := range repository.BlogPost {
		ret = append(ret, blogPost)
	}
	return ret, nil
}

func mapID(bp BlogPost) string {
	return fmt.Sprintf("%s", bp.BlogId)
}
