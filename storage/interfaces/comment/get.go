package comment

import (
	"database/sql"
	"errors"
	"time"

	"RTF/storage"
	"RTF/storage/interfaces/likes"
	"RTF/storage/interfaces/user"
	"RTF/types"

	"github.com/gofrs/uuid"
)

const QueryPostComments = "SELECT comm_id, user_id, SUBSTR(comment_date, '%Y-%m-%d') AS Date, comment FROM comments WHERE post_id = ?"

// function to return an array of post comments by id
func GetPostCommentsByID(req_user *types.User, postid uuid.UUID) ([]types.Comment, error) {
	stmt, err := storage.DB_Conn.Prepare(QueryPostComments)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(postid.String())
	if err != nil {
		return nil, errors.Join(types.ErrExec, err)
	}

	var comments []types.Comment

	for rows.Next() {
		c := &types.Comment{}
		var comment_id string
		var user_id string
		var comment_date string

		if err := rows.Scan(&comment_id, &user_id, &comment_date, &c.Content); err != nil {
			return nil, errors.Join(types.ErrScan, err)
		}

		u, err := user.GetSingleUser("user_id", user_id)
		if err != nil {
			return nil, errors.Join(types.ErrGetCommentDetails, err)
		}
		partial_user := types.PartialUser{
			ID:       u.ID,
			Username: u.Username,
		}

		// check if the user liked the comment
		liked, err := likes.CheckUserCommentLike(uuid.FromStringOrNil(comment_id), req_user.ID)
		if err != nil {
			return nil, errors.Join(types.ErrGetCommentDetails, err)
		}

		// set the liked value
		c.Liked = liked

		//! MIGHT BE BUGGY FROM HERE ON OUT
		if c.Likes, err = GetCommentLikes(uuid.FromStringOrNil(comment_id)); err != nil {
			return nil, errors.Join(types.ErrGetCommentDetails, err)
		}

		c.ID = uuid.FromStringOrNil(comment_id)
		c.Post_ID = postid
		c.User = partial_user
		if c.CreationDate, err = time.Parse("YYYY-MM-DD", comment_date); err != nil {
			return nil, errors.Join(types.ErrGetCommentDetails, err)
		}
		comments = append(comments, *c)
	}

	return comments, nil
}

// gets comment likes by id
func GetCommentLikes(commentID uuid.UUID) (int64, error) {
	query := "SELECT COUNT(*) FROM comment_likes WHERE comment_id = ?"

	row := storage.DB_Conn.QueryRow(query, commentID)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Join(types.ErrGetLikes, err)
	}

	return count, nil
}

// GetCommentsCount retrieves the count of comments for a given post ID.
func GetCommentsCount(postID string) (int, error) {
	var count int

	// Prepare the SQL query
	query := "SELECT COUNT(*) FROM comments WHERE post_id = ?"
	row := storage.DB_Conn.QueryRow(query, postID)

	// Scan the result into the count variable
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetCommentByID(commentID uuid.UUID) (*types.Comment, error) {
	query := `
        SELECT c.comm_id, c.user_id, c.post_id, c.comment_date, c.comment, 
               u.user_name, COUNT(cl.user_id) AS likes, 
               EXISTS(SELECT 1 FROM comment_likes cl WHERE cl.comment_id = c.comm_id AND cl.user_id = ?) AS liked
        FROM comments c
        JOIN users u ON c.user_id = u.user_id
        LEFT JOIN comment_likes cl ON c.comm_id = cl.comment_id
        WHERE c.comm_id = ?
        GROUP BY c.comm_id, c.user_id, c.post_id, c.comment_date, c.comment, u.user_name
    `

	var comment types.Comment
	var userID, postID, commentDate string

	err := storage.DB_Conn.QueryRow(query, commentID, commentID).Scan(
		&commentID, &userID, &postID, &commentDate, &comment.Content,
		&comment.User.Username, &comment.Likes, &comment.Liked,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Join(types.ErrScan, err)
	}

	comment.ID = uuid.FromStringOrNil(commentID.String())
	comment.User.ID = uuid.FromStringOrNil(userID)
	comment.Post_ID = uuid.FromStringOrNil(postID)

	if comment.CreationDate, err = time.Parse("YYYY-MM-DD", commentDate); err != nil {
		return nil, errors.Join(types.ErrGetCommentDetails, err)
	}

	return &comment, nil
}
