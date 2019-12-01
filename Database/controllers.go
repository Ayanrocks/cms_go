package Database

import (
	"github.com/cms/types"
	"github.com/kataras/golog"
)

// INSERT function
func Insert(data types.Post) error {
	golog.Debug(DB, data.ImageLinks)
	_, err := DB.Query(`INSERT INTO posts.posts(name, content, imagesLink) VALUES ($1, $2, $3)`, data.Name, data.Content, data.ImageLinks)

	if err != nil {
		golog.Error(err)
		return err
	}

	return nil
}

// Read Function
func Read(id string) (error, types.Post) {
	golog.Debug("reading DataBase")
	Post, err := DB.Query(`SELECT * FROM posts.posts WHERE id=$1`, id)
	if err != nil {
		golog.Error(err)
		return err, types.Post{}

	}
	data, err := Post.Columns()

	return nil, data
}
