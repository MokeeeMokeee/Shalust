package model

import (
	"shalust/api/pkg/db"
)

func GetAllIllustratio(data *[]ContentData) error {

	client, err := db.Init_mysql()
	if err != nil {
		return err
	}

	client.
		From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, content_data.image_url, content_data.image_index, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN content_data ON content_data.content_id = content_handling.content_id").
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", "testid").
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", "testid").
		Where("illustratio = 1").
		Scan(data)

	defer client.Close()
	return err
}

func GetUserIllustratio(user_id string, data *[]ContentData) error {
	client, err := db.Init_mysql()
	if err != nil {

		return err
	}

	client.From("content_handling").
		Select(`content_handling.*, user.user_name, user.icon_url, content_data.image_url, content_data.image_index, likes.user_liked, bookmarks.user_bookmarked`).
		Join("JOIN content_data ON content_data.content_id = content_handling.content_id").
		Join("JOIN user ON user.user_id = content_handling.user_id").
		Join("LEFT JOIN likes ON likes.content_id = content_handling.content_id AND likes.user_id = ?", "testid").
		Join("LEFT JOIN bookmarks ON bookmarks.content_id = content_handling.content_id AND bookmarks.user_id = ?", "testid").
		Where("content_handling.user_id = ?", user_id).
		Where("illustratio = 1").
		Scan(data)

	defer client.Close()
	return err
}
