package handlers

import (
	"net/http"
	"strconv"

	s "github.com/khihadysucahyo/go-echo-boilerplate/server"

	"github.com/khihadysucahyo/go-echo-boilerplate/responses"

	"github.com/khihadysucahyo/go-echo-boilerplate/requests"

	"github.com/khihadysucahyo/go-echo-boilerplate/repositories"

	"github.com/khihadysucahyo/go-echo-boilerplate/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type PostHandlers struct {
	server *s.Server
}

func NewPostHandlers(server *s.Server) *PostHandlers {
	return &PostHandlers{server: server}
}

func (p *PostHandlers) CreatePost(c echo.Context) error {
	createPostRequest := new(requests.CreatePostRequest)

	if err := c.Bind(createPostRequest); err != nil {
		return err
	}

	if err := createPostRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*repositories.JwtCustomClaims)
	id := claims.ID

	post := models.Post{
		Title:   createPostRequest.Title,
		Content: createPostRequest.Content,
		UserID:  id,
	}

	postRepository := repositories.NewPostRepository(p.server.DB)

	postRepository.Store(&post)

	return responses.MessageResponse(c, http.StatusCreated, "Post successfully created")
}

func (p *PostHandlers) DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	post := models.Post{}

	postRepository := repositories.NewPostRepository(p.server.DB)
	postRepository.GetPost(&post, id)

	if post.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
	}

	postRepository.Delete(&post)

	return responses.MessageResponse(c, http.StatusNoContent, "Post deleted successfully")
}

func (p *PostHandlers) GetPosts(c echo.Context) error {
	var posts []models.Post

	postRepository := repositories.NewPostRepository(p.server.DB)
	postRepository.GetPosts(&posts)

	for i := 0; i < len(posts); i++ {
		p.server.DB.Model(&posts[i]).Related(&posts[i].User)
	}

	response := responses.NewPostResponse(posts)
	return responses.Response(c, http.StatusOK, response)
}

func (p *PostHandlers) UpdatePost(c echo.Context) error {
	updatePostRequest := new(requests.UpdatePostRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(updatePostRequest); err != nil {
		return err
	}

	if err := updatePostRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	post := models.Post{}

	postRepository := repositories.NewPostRepository(p.server.DB)
	postRepository.GetPost(&post, id)

	if post.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
	}

	postRepository.Update(&post, updatePostRequest)

	return responses.MessageResponse(c, http.StatusOK, "Post successfully updated")
}
