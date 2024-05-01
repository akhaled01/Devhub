package post

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"

	"RTF/storage"
	"RTF/storage/interfaces/categories"
	"RTF/types"
	"RTF/utils"
)

const NEWPOSTQUERY = `
    INSERT INTO posts (post_id, user_id, creation_date, post_content, post_image_path)
  VALUES ($1, $2, $3, $4, $5)
  `

// This function saves a post object to the DB
func SavePostInDB(p types.Post, post_categories []int) error {
	stmt, err := storage.DB_Conn.Prepare(NEWPOSTQUERY)
	if err != nil {
		return errors.Join(errors.New("error preparing SavePostInDB query"), err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(
		p.ID,
		p.User.ID.String(),
		p.CreationDate,
		p.Content,
		p.Image_Path,
	); err != nil {
		return errors.Join(errors.New("error executing SavePostInDB query"), err)
	}

	if err = categories.AssignPostCategory(p.ID, post_categories); err != nil {
		return errors.Join(errors.New("error assigning categories to post"), err)
	}

	utils.InfoConsoleLog(fmt.Sprintf("New post created with ID: %s", p.ID))
	return nil
}

// This function construct a new post with default stats in
// in order to facilitate communication with the DB.
func ConstructNewPostFromRequest(r types.PostCreationRequest) (types.Post, error) {
	new_pid, err := uuid.NewV7()
	if err != nil {
		return (types.Post{}), errors.Join(types.ErrUUID, err)
	}

	author_session, ok := types.Sessions[uuid.FromStringOrNil(r.Session_id)]
	if !ok {
		return (types.Post{}), errors.Join(types.ErrSessionNotFound, err)
	}

	post_author := author_session.User

	var categories_as_string []string
	for cat_idx := range r.Post_category {
		category_as_string, err := categories.GetFullCategory(r.Post_category[cat_idx])
		if err != nil {
			return (types.Post{}), errors.Join(types.ErrCats, err)
		}

		categories_as_string = append(categories_as_string, category_as_string.Name)
	}

	partial_post_author := types.PartialUser{
		ID:       post_author.ID,
		Username: post_author.Username,
	}
	var image_path string

	if r.Post_image_base64 == "" {
		image_path = ""
	} else {
		image_path, err = utils.SaveImage(r.Post_image_base64, "post")
		if err != nil {
			return (types.Post{}), errors.Join(types.ErrImage, err)
		}
	}

	return types.Post{
		ID:           new_pid,
		User:         partial_post_author,
		Content:      r.Post_text,
		CreationDate: time.Now(),
		Image_Path:   image_path,
		Category:     categories_as_string,
	}, nil
}
