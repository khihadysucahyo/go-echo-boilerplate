package repositories

import (
	"github.com/khihadysucahyo/go-echo-boilerplate/models"
	"github.com/khihadysucahyo/go-echo-boilerplate/requests"

	"github.com/jinzhu/gorm"
)

type PostRepositoryQ interface {
	GetPosts(posts *[]models.Post)
	GetPost(post *models.Post, id int)
	Store(post *models.Post)
	Delete(post *models.Post)
	Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest)
}
type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (postRepository *PostRepository) GetPosts(posts *[]models.Post) {
	postRepository.DB.Find(posts)
}

func (postRepository *PostRepository) GetPost(post *models.Post, id int) {
	postRepository.DB.Where("id = ? ", id).Find(post)
}

func (postRepository *PostRepository) Store(post *models.Post) {
	postRepository.DB.Create(post)
}

func (postRepository *PostRepository) Delete(post *models.Post) {
	postRepository.DB.Delete(post)
}

func (postRepository *PostRepository) Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest) {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title
	postRepository.DB.Save(post)
}
