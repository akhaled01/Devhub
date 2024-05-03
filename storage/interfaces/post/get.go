package post

import (
	"database/sql"
	"errors"

	"RTF/storage"
	"RTF/storage/interfaces/categories"
	"RTF/storage/interfaces/comment"
	"RTF/storage/interfaces/likes"
	"RTF/storage/interfaces/user"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const (
	PostByIDQuery          = `SELECT user_id, creation_date, post_content, post_image_path FROM posts WHERE post_id = ?`
	AllPostsQuery          = `SELECT post_id FROM posts ORDER BY creation_date desc`
	GETUSERLIKEDPOSTSQUERY = `SELECT post_id FROM post_likes WHERE user_id = ?`
)

// Function that Gets a post from a DB by its ID
func GetPostByID(id uuid.UUID) (types.Post, error) {
	stmt, err := storage.DB_Conn.Prepare(PostByIDQuery)
	if err != nil {
		return (types.Post{}), errors.Join(types.ErrPrepare, err)
	}

	p := types.Post{}
	p.ID = id
	var stringUserID string

	// scan DB results to populate the post struct
	if err := stmt.QueryRow(id).Scan(&stringUserID, &p.CreationDate, &p.Content, &p.Image_Path); err != nil {
		if err == sql.ErrNoRows {
			return (types.Post{}), types.ErrPostNotFound
		}
		return (types.Post{}), errors.Join(types.ErrScan, err)
	}

	// get the creator of the post
	creator, err := user.GetSingleUser("user_id", stringUserID)
	if err != nil {
		return (types.Post{}), errors.Join(errors.New("error getting user for this post"), err)
	}

	partial_creator := types.PartialUser{
		ID:       creator.ID,
		Username: creator.Username,
	}

	// get post likes
	if p.Likes, err = likes.GetPostLikes(id); err != nil {
		return (types.Post{}), err
	}

	// get post comments
	if p.Comments, err = comment.GetPostCommentsByID(id); err != nil {
		return (types.Post{}), err
	}

	// Get post categories
	post_cats, _ := categories.GetPostCategories(p.ID)

	p.User = partial_creator
	p.Category = post_cats
	return p, nil
}

// function that gets posts that are liked by a specific user.
// takes in a user id
func GetUserLikedPosts(userid uuid.UUID) ([]types.Post, error) {
	stmt, err := storage.DB_Conn.Prepare(GETUSERLIKEDPOSTSQUERY)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	posts := []types.Post{}
	rows, err := stmt.Query(userid.String())
	if err != nil {
		return nil, errors.Join(types.ErrExec, err)
	}
	defer rows.Close()

	for rows.Next() {
		var post_id string
		if err := rows.Scan(&post_id); err != nil {
			return nil, errors.Join(errors.New("error scanning to post_id"), err)
		}

		var p types.Post
		if p, err = GetPostByID(uuid.FromStringOrNil(post_id)); err != nil {
			return nil, errors.Join(errors.New("error getting post"), err)
		}
		posts = append(posts, p)
	}

	return posts, nil
}

// Invokes a query to get all posts sorted by creation_date.
func AllPostsFromDB() ([]types.Post, error) {
	all_posts := []types.Post{}
	stmt, err := storage.DB_Conn.Prepare(AllPostsQuery)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.Join(types.ErrExec, err)
	}
	defer rows.Close()

	for rows.Next() {
		var str_post_id string
		if err := rows.Scan(&str_post_id); err != nil {
			return nil, errors.Join(types.ErrScan, err)
		}

		post, err := GetPostByID(uuid.FromStringOrNil(str_post_id))
		post.Number_of_comments, _ = comment.GetCommentsCount(post.ID.String())
		if err != nil {
			return nil, errors.Join(types.ErrAppendPost, err)
		}

		all_posts = append(all_posts, post)
	}

	return all_posts, nil
}
