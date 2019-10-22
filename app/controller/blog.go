package controller

import (
	"../../config"
	"../model"
	"encoding/json"
	"log"
	"net/http"
)

// BlogAll return all blog
func BlogAll(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	var arrBlog []model.Blog
	var response model.BlogResponse

	db := config.DBConnect()
	defer db.Close()

	rows, err := db.Query("Select id, title, text from blog")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Text); err != nil {
			log.Fatal(err)
		} else {
			arrBlog = append(arrBlog, blog)
		}
	}

	response.Status = "success"
	response.Data = arrBlog

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
