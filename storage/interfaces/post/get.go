package post

import (
	"database/sql"
	"errors"

	"RTF/storage"
	"RTF/storage/interfaces/comment"
	"RTF/storage/interfaces/likes"
	"RTF/storage/interfaces/user"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const PostByIDQuery = `SELECT user_id, creation_date, post_content, post_image_path FROM posts WHERE post_id = ?`

const GETUSERLIKEDPOSTSQUERY = `SELECT post_id FROM posts_interaction WHERE user_id = ? AND actions_type = 1`

// Function that Gets a post from a DB by its ID
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
	if p.Likes, err = likes.GetPostLikes(id); err != nil {
		return nil, err
	}

	// get post comments
	if p.Comments, err = comment.GetPostCommentsByID(id); err != nil {
		return nil, err
	}

	p.User = *creator
	return p, nil
}

// function that gets posts that are liked by a specific user.
// takes in a user id
func GetUserLikedPosts(userid uuid.UUID) ([]types.Post, error) {
	stmt, err := storage.DB_Conn.Prepare(GETUSERLIKEDPOSTSQUERY)
	if err != nil {
		return nil, errors.Join(errors.New("error preparing GetUserLikedPosts query"), err)
	}
	defer stmt.Close()

	posts := []types.Post{}
	rows, err := stmt.Query(userid.String())
	if err != nil {
		return nil, errors.Join(errors.New("error executing GetUserLikedPosts query"), err)
	}
	defer rows.Close()

	for rows.Next() {
		var post_id string
		if err := rows.Scan(&post_id); err != nil {
			return nil, errors.Join(errors.New("error scanning to post_id"), err)
		}

		p := &types.Post{}
		if p, err = GetPostByID(uuid.FromStringOrNil(post_id)); err != nil {
			return nil, errors.Join(errors.New("error getting post"), err)
		}
		posts = append(posts, *p)
	}

	return posts, nil
}
