package likes

import (
	"database/sql"
	"errors"

	"RTF/storage"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const GETINTERACTIONSQUERY = `SELECT action_type FROM posts_interactions 
	WHERE post_id = ? AND user_id = ?
`

// Takes in a user's id and a post's id and checks if the user liked the post
func CheckUserPostLike(postid uuid.UUID, userid uuid.UUID, like int) (bool, error) {
	stmt, err := storage.DB_Conn.Prepare(GETINTERACTIONSQUERY)
	if err != nil {
		return false, errors.Join(errors.New("error preparing CheckUserPostLike query"), err)
	}
	defer stmt.Close()

	var isLiked bool

	if err := stmt.QueryRow(postid.String(), userid.String()).Scan(&isLiked); err != nil {
		return false, errors.Join(errors.New("error executing CheckUserPostLike query"), err)
	}

	return isLiked, nil
}

// function that gets post likes by id
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

// function that gets posts that are liked by a specific user
//
// takes in a user id
func GetUserLikedPosts(id uuid.UUID) ([]types.Post, error) {
	posts := []types.Post{}
	return posts, nil
}
