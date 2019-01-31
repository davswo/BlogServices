package repository

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"time"
)

type blogRepositoryMongo struct {
	client                  *mongo.Client
	blogPostsCollectionName string
	databaseName            string
}

// NewOrderRepositoryMongo is used to instantiate and return the DB implementation of the OrderRepository.
func NewOrderRepositoryMongo() (BlogRepository, error) {
	rand.Seed(time.Now().UnixNano())

	client, err := mongo.Connect(context.TODO(), "mongodb://mongod-0.default.svc.cluster.local:27017")
	if err != nil {
		return nil, errors.Wrap(err, "Error connecting to MongoDB %v.")
	}

	return &blogRepositoryMongo{client, "blog_posts", "blog_app"}, nil
}

func (repository *blogRepositoryMongo) InsertBlogPost(blogPost BlogPost) error {
	collection := repository.client.Database(repository.databaseName).Collection(repository.blogPostsCollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, blogPost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted new BlogPost: ", res.InsertedID)
	return err
}

func (repository *blogRepositoryMongo) GetBlogPosts() ([]BlogPost, error) {
	collection := repository.client.Database(repository.databaseName).Collection(repository.blogPostsCollectionName)
	options := options.Find()
	var results []BlogPost
	filter := bson.D{}
	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem BlogPost
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results, nil
}
