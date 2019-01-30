package repository

import (
	"math/rand"
	"time"
)

type blogRepositoryMemory struct {
	BlogPost map[int]BlogPost
}

// NewOrderRepositoryMemory is used to instantiate and return the DB implementation of the OrderRepository.
func NewOrderRepositoryMemory() BlogRepository {
	rand.Seed(time.Now().UnixNano())
	return &blogRepositoryMemory{BlogPost: make(map[int]BlogPost)}
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

func mapID(repository *blogRepositoryMemory) int {
	uid := rand.Intn(10000)
	if _, exists := repository.BlogPost[uid]; exists {
		return mapID(repository)
	}
	return uid
}
