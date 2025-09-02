package posts

import (
	"fmt"
	"socio/internals/dto"
	"socio/internals/notifications"
	"socio/internals/validator"
	"socio/services/posts"
	"socio/services/users"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Add godoc
// @Summary Create a new post
// @Description Create a new post for a specific user
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param post body dto.PostCreate true "Post creation data"
// @Success 201 {object} dto.Post
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /socio/users/{id}/posts [post]
func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var post dto.PostCreate

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	if err := validator.Payload(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	ps := posts.New()
	ps.Post = &dto.Post{}

	ps.Post.UserID = userID
	ps.Post.Content = post.Content

	ps.Create(ctx)

	msg := fmt.Sprintf("Hello, your friend %v has created a new post", us.User.Name)
	notifications.NotifyUsers(ctx, userID, msg)

	return c.Status(fiber.StatusCreated).JSON(ps.Post)
}

// Get godoc
// @Summary Get user posts
// @Description Get all posts for a specific user
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} dto.Post
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User not found"
// @Router /socio/users/{id}/posts [get]
func Get(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	ps := posts.New()
	ps.Posts = &dto.Posts{}
	ps.UserID = userID

	ps.GetAll(ctx)

	return c.Status(fiber.StatusOK).JSON(ps.Posts)
}

// Delete godoc
// @Summary Delete a post
// @Description Delete a specific post by user ID and post ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param post_id path string true "Post ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User not found"
// @Router /socio/users/{id}/posts/{post_id} [delete]
func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	postID, err := uuid.Parse(c.Params("post_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	ps := posts.New()
	ps.Posts = &dto.Posts{}

	ps.UserID = userID
	ps.ID = postID

	ps.Delete(ctx)

	return c.SendStatus(fiber.StatusNoContent)
}
