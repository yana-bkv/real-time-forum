package repositories

type postCategoryRepository struct{}

func NewPostCategoryRepository() PostCategoryRepository {
	return &postCategoryRepository{}
}

func (r *postCategoryRepository) Assign(postId, categoryId int) error {
	return nil
}

func (r *postCategoryRepository) Remove(postId, categoryId int) error {
	return nil
}
