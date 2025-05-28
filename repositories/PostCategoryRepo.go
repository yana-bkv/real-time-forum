package repositories

import (
	"fmt"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"strings"
)

type postCategoryRepository struct{}

func NewPostCategoryRepository() PostCategoryRepository {
	return &postCategoryRepository{}
}

func (r *postCategoryRepository) Assign(postCategory *models.PostCategory) error {
	sqlStmt := database.SqlPostCategoryDb("assignCategoryToPost")

	stmt, err := database.DB.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, catID := range postCategory.CategoryIDs {
		_, err := stmt.Exec(postCategory.PostID, catID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *postCategoryRepository) Remove(postId, categoryId int) error {
	return nil
}

func (r *postCategoryRepository) GetTagByPostId(postId int) ([]int, error) {
	sqlStmt := database.SqlPostCategoryDb("getTagByPostId")
	rows, err := database.DB.Query(sqlStmt, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categoryIDs []int
	for rows.Next() {
		var categoryID int
		if err := rows.Scan(&categoryID); err != nil {
			return nil, err
		}
		categoryIDs = append(categoryIDs, categoryID)
	}

	return categoryIDs, nil
}

func (r *postCategoryRepository) GetCategoryNamesByIDs(ids []int) (map[int]string, error) {
	if len(ids) == 0 {
		return map[int]string{}, nil
	}

	// Подготовим плейсхолдеры: $1, $2, $3 и т.д.
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`SELECT id, text FROM categories WHERE id IN (%s)`, strings.Join(placeholders, ","))
	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]string)
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		result[id] = name
	}

	return result, nil
}
