package post

import (
	"errors"

	"RTF/storage"
	"RTF/types"
)

const NEWPOSTQUERY = `
    INSERT INTO posts (user_id, creation_date, post_content, post_image_path)
    VALUES (:UserID, :CreationDate, :Content, :Image_Path)
  `

// This function saves a entire post to the database.
//
// Interfaces are used to map bind values to queries
// in any desired order.
func SavePostInDB(p types.Post) error {
	stmt, err := storage.DB_Conn.Prepare(NEWPOSTQUERY)
	if err != nil {
		return errors.Join(errors.New("error preparing SavePostInDB query"), err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(
		map[string]interface{}{
			":UserID":       p.User.ID.String(),
			":CreationDate": p.CreationDate,
			":Content":      p.Content,
			":Image_Path":   p.Image_Path,
		},
	); err != nil {
		return errors.Join(errors.New("error executing SavePostInDB query"), err)
	}

	return nil
}
