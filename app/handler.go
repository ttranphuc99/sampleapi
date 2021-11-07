package app

import (
	"fmt"
	"log"
	"net/http"
	"sampleapi/app/models"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Call API from '/'. Process in IndexHandler()")
		fmt.Fprint(w, "Welcome to SampleAPI")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, nil, http.StatusInternalServerError)
			return
		}

		// create the post
		p := &models.Post{
			Title:   req.Title,
			Author:  req.Author,
			Content: req.Content,
		}

		// save in db
		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("Cannot save post to db. err=%v \n", err)
			sendResponse(w, nil, http.StatusInternalServerError)
			return
		}

		resp := mapToJSON(p)
		sendResponse(w, resp, http.StatusOK)
	}
}

func (a *App) GetPostsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := a.DB.GetPosts()

		if err != nil {
			log.Printf("Cannot get posts, err=%v \n", err)
			sendResponse(w, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, len(posts))
		log.Println(len(posts))
		for idx, post := range posts {
			resp[idx] = mapToJSON(post)
		}

		sendResponse(w, resp, http.StatusOK)
	}
}
