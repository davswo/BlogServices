package repository

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"github.com/vrischmann/envconfig"
	"io"

	dbPkg "github.com/kyma-project/examples/http-db-service/internal/mssqldb"
)

const (
	insertQuery = "INSERT INTO %s (blog_id, namespace, total) VALUES (?, ?, ?)"
	getQuery    = "SELECT * FROM %s"
	deleteQuery = "DELETE FROM %s"
)

type blogRepositorySQL struct {
	database      dbQuerier
	blogTableName string
}

//go:generate mockery -name dbQuerier -inpkg
type dbQuerier interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	io.Closer
}

func NewBlogRepositoryDb() (BlogRepository, error) {
	var (
		dbCfg    dbPkg.Config
		database dbQuerier
		err      error
	)
	if err = envconfig.Init(&dbCfg); err != nil {
		return nil, errors.Wrap(err, "Error loading db configuration %v.")
	}
	if database, err = dbPkg.InitDb(dbCfg); err != nil {
		return nil, errors.Wrap(err, "Error loading db configuration %v.")
	}
	return &blogRepositorySQL{database, dbCfg.DbBlogTableName}, nil
}

type sqlError interface {
	sqlErrorNumber() int32
}

func (repository *blogRepositorySQL) InsertBlogPost(blog BlogPost) error {
	q := fmt.Sprintf(insertQuery, dbPkg.SanitizeSQLArg(repository.blogTableName))
	log.Debugf("Running insert blog post query: '%q'.", q)
	_, err := repository.database.Exec(q, blog.BlogId, blog.Title, blog.Text, blog.Author)

	if errorWithNumber, ok := err.(sqlError); ok {
		if errorWithNumber.sqlErrorNumber() == dbPkg.PrimaryKeyViolation {
			return ErrDuplicateKey
		}
	}

	return errors.Wrap(err, "while inserting blog post")
}

func (repository *blogRepositorySQL) GetBlogPosts() ([]BlogPost, error) {
	q := fmt.Sprintf(getQuery, dbPkg.SanitizeSQLArg(repository.blogTableName))
	log.Debugf("Quering blog posts: '%q'.", q)
	rows, err := repository.database.Query(q)

	if err != nil {
		return nil, errors.Wrap(err, "while reading blog posts from DB")
	}

	defer rows.Close()
	return readFromResult(rows)
}

func readFromResult(rows *sql.Rows) ([]BlogPost, error) {
	blogPosts := make([]BlogPost, 0)
	for rows.Next() {
		blogPost := BlogPost{}
		if err := rows.Scan(&blogPost.BlogId, &blogPost.Title, &blogPost.Text, &blogPost.Author); err != nil {
			return []BlogPost{}, err
		}
		blogPosts = append(blogPosts, blogPost)
	}
	return blogPosts, nil
}
