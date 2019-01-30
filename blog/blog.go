package blog

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/davswo/BlogServices/repository"
	"io/ioutil"
	"net/http"
)

type BlogService struct {
	repository repository.BlogRepository
}

func NewBlogService(repository repository.BlogRepository) BlogService {
	return BlogService{repository}
}

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (blogService BlogService) InsertBlogPost(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Error parsing request.", err)
		respondWithCodeAndMessage(http.StatusInternalServerError, "Internal error.", w)
		return
	}

	defer r.Body.Close()
	var blogPost repository.BlogPost
	err = json.Unmarshal(b, &blogPost)
	if err != nil || blogPost.BlogId == "" || blogPost.Title == "" || blogPost.Text == "" || blogPost.Author == "" {
		respondWithCodeAndMessage(http.StatusBadRequest, "Invalid request body, blogId / title / text / author can not be empty.", w)
		return
	}

	log.Debugf("Inserting blogPost: '%+v'.", blogPost)

	err = blogService.repository.InsertBlogPost(blogPost)

	switch err {
	case nil:
		w.WriteHeader(http.StatusCreated)
	case repository.ErrDuplicateKey:
		respondWithCodeAndMessage(http.StatusConflict, fmt.Sprintf("BlogPost %s already exists.", blogPost.BlogId), w)
	default:
		log.Error(fmt.Sprintf("Error inserting blogPost: '%+v'", blogPost), err)
		respondWithCodeAndMessage(http.StatusInternalServerError, "Internal error.", w)
	}
}

func (blogService BlogService) GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	log.Debug("Retrieving blog posts")

	blogPosts, err := blogService.repository.GetBlogPosts()
	if err != nil {
		log.Error("Error retrieving blogPosts.", err)
		respondWithCodeAndMessage(http.StatusInternalServerError, "Internal error.", w)
		return
	}

	if err = respondBlogPosts(blogPosts, w); err != nil {
		log.Error("Error sending blogPosts response.", err)
		respondWithCodeAndMessage(http.StatusInternalServerError, "Internal error.", w)
		return
	}
}

func respondBlogPosts(blogPosts []repository.BlogPost, w http.ResponseWriter) error {
	if len(blogPosts) == 0 {
		blogPosts = getStaticBlogEntries()
	}
	body, err := json.Marshal(blogPosts)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(body); err != nil {
		return err
	}
	return nil
}

func respondWithCodeAndMessage(code int, msg string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(code)
	response := errorResponse{code, msg}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error("Error sending response", err)
	}
}

func getStaticBlogEntries() []repository.BlogPost {
	return []repository.BlogPost{
		{Title: "O’zapft is!",
			Text: "Bavaria ipsum dolor sit amet Zidern so, " +
				"iabaroi Mamalad da in da. Gwihss mehra barfuaßat großherzig singd a am acht’n Tag schuf Gott des Bia. " +
				"Do des back mas, sog i schnacksln. Hallelujah sog i, luja Wiesn bittschön a fescha Bua sodala auf der Oim, " +
				"da gibt’s koa Sünd is ma Wuascht kimmt i Biaschlegl: Wann griagd ma nacha wos z’dringa umma nomoi mei Weiznglasl " +
				"jo mei is des schee aba des basd scho ognudelt mechad: Kuaschwanz Habedehre hawadere midananda obandln, " +
				"des wiad a Mordsgaudi kumm geh! Ned Xaver a ganze Hoiwe des is a gmahde Wiesn nix Gwiass woass ma ned unbandig Watschnbaam" +
				" is des liab Wiesn mi. Gscheckate ned is ma Wuascht, Steckerleis. Obandln Baamwach hogg di hera am acht’n Tag schuf " +
				"Gott des Bia nia is, ognudelt pfiad de nackata Wurschtsolod. Allerweil Watschnbaam Schbozal auszutzeln nois Bussal von Greichats. " +
				"Oba da jo mei is des schee, .",
			Author: "Huaba Bua",
		},
		{Title: "Ja kurze fix!",
			Text: "Bavaria ipsum dolor sit amet fias Bradwurschtsemmal, i mechad dee Schwoanshaxn. " +
				"Nia need Stubn fensdaln a ganze, Bussal pfenningguat. " +
				"I daad wuid fei Biawambn mogsd a Bussal Goaßmaß scheans naa aasgem i.",
			Author: "Goaßl Sepp",
		},
	}
}
