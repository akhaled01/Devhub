package post

import (
	"database/sql"
	"errors"

	"RTF/storage"
	"RTF/storage/interfaces/comment"
	"RTF/storage/interfaces/user"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const PostByIDQuery = "SELECT user_id, creation_date, post_content, post_image_path FROM posts WHERE post_id = ?"

// returns a post by inputting its ID
func GetPostByID(id uuid.UUID) (*types.Post, error) {
	stmt, err := storage.DB_Conn.Prepare(PostByIDQuery)
	if err != nil {
		return nil, errors.Join(errors.New("error preparing GetPostByID sql statement"), err)
	}

	p := &types.Post{}
	p.ID = id
	var stringUserID string

	// scan DB results to populate the post struct
	if err := stmt.QueryRow(id).Scan(&stringUserID, &p.CreationDate, &p.Content, &p.Image_Path); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("post not found")
		}
		return nil, errors.Join(errors.New("error scanning GetPostByID rows"), err)
	}

	// get the creator of the post
	creator, err := user.GetSingleUser("user_id", stringUserID)
	if err != nil {
		return nil, errors.Join(errors.New("error getting user for this post"), err)
	}

	// get post likes
	if p.Likes, err = GetPostLikes(id); err != nil {
		return nil, err
	}

	// get post comments
	if p.Comments, err = comment.GetPostCommentsByID(id); err != nil {
		return nil, err
	}

	p.User = *creator
	return p, nil
}

//function gets post likes by id
func GetPostLikes(postID uuid.UUID) (int64, error) {
	query := "SELECT COUNT(*) FROM post_interactions WHERE post_id = ?"

	row := storage.DB_Conn.QueryRow(query, postID)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Join(errors.New("error getting post likes"), err)
	}

	return count, nil
}
