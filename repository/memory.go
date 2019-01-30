package repository

import (
	"math/rand"
	"time"
)

type blogRepositoryMemory struct {
	BlogPost map[string]BlogPost
}

// NewOrderRepositoryMemory is used to instantiate and return the DB implementation of the OrderRepository.
func NewOrderRepositoryMemory() BlogRepository {
	rand.Seed(time.Now().UnixNano())
	return &blogRepositoryMemory{BlogPost: make(map[string]BlogPost)}
}

func (repository *blogRepositoryMemory) InsertBlogPost(blogPost BlogPost) error {
	id := mapID(repository)
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

func mapID(repository *blogRepositoryMemory) string {
	source := rand.NewSource(time.Now().UnixNano())
	uid := rand.New(source)
	if _, exists := repository.BlogPost[string(uid.Int())]; exists {
		return mapID(repository)
	}
	return string(uid.Int())
}
