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
func GetPostByID(req_user *types.User, id uuid.UUID) (types.Post, error) {
	// prepare the query
	stmt, err := storage.DB_Conn.Prepare(PostByIDQuery)
	if err != nil {
		return (types.Post{}), errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

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
		Gender:   creator.Gender,
	}

	// get post likes
	if p.Likes, err = likes.GetPostLikes(id); err != nil {
		return (types.Post{}), err
	}

	// get post comments
	if p.Comments, err = comment.GetPostCommentsByID(req_user, id); err != nil {
		return (types.Post{}), err
	}

	if p.Number_of_comments, err = comment.GetCommentsCount(id.String()); err != nil {
		return (types.Post{}), errors.Join(types.ErrAppendPost, err)
	}

	// Get post categories
	post_cats, _ := categories.GetPostCategories(p.ID)

	p.User = partial_creator
	p.Category = post_cats

	// Check if the user has liked the post
	liked, err := likes.CheckUserPostLike(p.ID, req_user.ID)
	if err != nil {
		return types.Post{}, errors.Join(types.ErrAppendPost, err)
	}
	// Set the post liked field to true if the user has liked the post
	p.Liked = liked

	return p, nil
}

// function that gets posts that are liked by a specific user.
// takes in a user id
func GetUserLikedPosts(req_user *types.User) ([]types.Post, error) {
	// prepare the query
	stmt, err := storage.DB_Conn.Prepare(GETUSERLIKEDPOSTSQUERY)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	// execute the query
	posts := []types.Post{}
	rows, err := stmt.Query(req_user.ID.String())
	if err != nil {
		return nil, errors.Join(types.ErrExec, err)
	}
	// close the rows after the query is done
	defer rows.Close()

	// scan the results
	for rows.Next() {
		// get the post id
		var post_id string
		if err := rows.Scan(&post_id); err != nil {
			return nil, errors.Join(errors.New("error scanning to post_id"), err)
		}

		// get the post
		var p types.Post
		if p, err = GetPostByID(req_user, uuid.FromStringOrNil(post_id)); err != nil {
			return nil, errors.Join(errors.New("error getting post"), err)
		}
		// append the post to the slice
		posts = append(posts, p)
	}

	return posts, nil
}

// Invokes a query to get all posts sorted by creation_date.
func AllPostsFromDB(req_user *types.User) ([]types.Post, error) {
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
		// Get the post id
		var str_post_id string
		if err := rows.Scan(&str_post_id); err != nil {
			return nil, errors.Join(types.ErrScan, err)
		}

		// Get the post from the DB
		post, err := GetPostByID(req_user, uuid.FromStringOrNil(str_post_id))
		post.Number_of_comments, _ = comment.GetCommentsCount(post.ID.String())
		if err != nil {
			return nil, errors.Join(types.ErrAppendPost, err)
		}

		// Check if the user has liked the post
		liked, err := likes.CheckUserPostLike(post.ID, req_user.ID)
		if err != nil {
			return nil, errors.Join(types.ErrAppendPost, err)
		}
		// Set the post liked field to true if the user has liked the post
		post.Liked = liked
		// Append the post to the list of posts
		all_posts = append(all_posts, post)
	}

	return all_posts, nil
}
