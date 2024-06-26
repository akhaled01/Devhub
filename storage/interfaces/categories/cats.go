package categories

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"

	"RTF/storage"
	"RTF/types"
)

const (
	GET_POST_CATEGORY = `SELECT cat_id FROM post_categories WHERE post_id = ?`
	GET_CATEGORY      = `SELECT * FROM category WHERE cat_id = ?`
	ADD_THREAD        = `INSERT INTO post_categories (post_id, cat_id) VALUES (?, ?)`
	CHECK_CAT_EXISTS  = `SELECT EXISTS(SELECT 1 FROM category WHERE cat_id = ?)`
	GET_FULL_CATS     = `SELECT cat_id, category FROM category`
)

// Invokes GET_CATEGORY query that returns a full category struct (name and id)
func GetFullCategory(categoryid int) (*types.Category, error) {
	stmt, err := storage.DB_Conn.Prepare(GET_CATEGORY)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()
	var category types.Category

	if err := stmt.QueryRow(categoryid).Scan(&category.Id, &category.Name); err != nil {
		return nil, errors.Join(types.ErrScan, err)
	}

	return &category, nil
}

// Invokes GET_POST_CATEGORY that returns full category that is assigned to a specific post
func GetPostCategories(postid uuid.UUID) ([]string, error) {
	query := `
	SELECT category
	FROM post_categories
	JOIN category ON post_categories.cat_id = category.cat_id
	WHERE post_id = ?
`

	rows, err := storage.DB_Conn.Query(query, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// invokes ADD_THREAD that adds a new category-post mapping record into the post_categories table.
func AssignPostCategory(postid uuid.UUID, catids []int) error {
	// Inserting the post category into the database
	query := "INSERT INTO post_categories (post_id, cat_id) VALUES (?, ?)"
	for _, catID := range catids {
		if _, err := storage.DB_Conn.Exec(query, postid, catID); err != nil {
			return fmt.Errorf("failed to insert the post category")
		}
	}

	// If no categories in the post, insert it into the General category
	if len(catids) == 0 {
		query = "INSERT INTO post_categories (post_id, cat_id) VALUES (?, ?)"
		if _, err := storage.DB_Conn.Exec(query, postid, 1); err != nil {
			return fmt.Errorf("failed to insert the post General category")
		}
	}

	return nil
}

// Invokes CHECK_CAT_EXISTS that checks if a category by an id exists
func CheckCategoryExists(id int) (bool, error) {
	stmt, err := storage.DB_Conn.Prepare(CHECK_CAT_EXISTS)
	if err != nil {
		return false, errors.Join(errors.New("error preparing CheckCategoryExists query"), err)
	}
	defer stmt.Close()

	var is_exist bool

	if err := stmt.QueryRow(id).Scan(&is_exist); err != nil {
		return false, errors.Join(errors.New("error executing CheckCategoryExists query"), err)
	}

	return is_exist, nil
}

func GetAllCategoryInfo() ([]types.Category, error) {
	var container []types.Category

	rows, err := storage.DB_Conn.Query(GET_FULL_CATS)
	if err != nil {
		return container, err
	}
	defer rows.Close()

	for rows.Next() {
		var result types.Category
		if err := rows.Scan(&result.Id, &result.Name); err != nil {
			return container, err
		}
		container = append(container, result)
	}

	return container, nil
}
