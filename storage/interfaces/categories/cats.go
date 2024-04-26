package categories

import (
	"errors"

	"github.com/gofrs/uuid"

	"RTF/storage"
	"RTF/types"
)

const (
	GET_POST_CATEGORY = `SELECT cat_id FROM post_categories WHERE post_id = ?`
	GET_CATEGORY      = `SELECT * FROM category WHERE cat_id = ?`
	ADD_THREAD        = `INSERT INTO post_categories (post_id, cat_id) VALUES (?, ?)`
	CHECK_CAT_EXISTS  = `SELECT EXISTS(SELECT 1 FROM category WHERE cat_id = ?)`
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
func GetPostCategories(postid uuid.UUID) (*types.Category, error) {
	stmt, err := storage.DB_Conn.Prepare(GET_POST_CATEGORY)
	if err != nil {
		return nil, errors.Join(errors.New("error preparing GetPostCategories query"), err)
	}
	defer stmt.Close()
	var category_id int

	if err := stmt.QueryRow(postid.String()).Scan(category_id); err != nil {
		return nil, errors.Join(errors.New("error executing GetPostCategories query"), err)
	}

	cat, err := GetFullCategory(category_id)
	if err != nil {
		return nil, errors.Join(errors.New("error getting full category"), err)
	}

	return cat, nil
}

// invokes ADD_THREAD that adds a new category-post mapping record into the post_categories table.
func AssignPostCategory(postid uuid.UUID, catid int) error {
	stmt, err := storage.DB_Conn.Prepare(ADD_THREAD)
	if err != nil {
		return errors.Join(errors.New("error preparing AssignPostCategory query"), err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(postid.String(), catid); err != nil {
		return errors.Join(errors.New("error executing AssignPostCategory query"), err)
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
