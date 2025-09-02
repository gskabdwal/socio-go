package friendship

import (
	"fmt"
	"socio/internals/cache"
	"socio/internals/dto"
	"socio/internals/validator"
	"socio/services/friendships"
	"socio/services/users"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Add godoc
// @Summary Add a friend
// @Description Create a new friendship between users
// @Tags friendships
// @Accept json
// @Produce json
// @Param friend body dto.FriendsCrate true "Friendship data"
// @Success 201 {object} dto.Friends
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /socio/friends [post]
func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var friend dto.FriendsCrate

	if err := c.BodyParser(&friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	if err := validator.Payload(friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = friend.UserID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	us.User = &dto.User{}
	us.User.ID = friend.FriendID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	fs := friendships.New()
	fs.Friends = &dto.Friends{}

	fs.Friends.UserID = friend.UserID
	fs.Friends.FriendID = friend.FriendID

	fs.Create(ctx)

	if err := cache.Client().Del(ctx, friend.UserID.String()).Err(); err != nil {
		fmt.Printf("Error invalidating cache: %v", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fs.Friends)
}

// Get godoc
// @Summary Get user's friends
// @Description Get all friends for a specific user
// @Tags friendships
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} dto.AllFriends
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User not found"
// @Router /socio/friends/{id} [get]
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

	fs := friendships.New()
	fs.UserID = userID
	fs.GetAll(ctx)

	return c.Status(fiber.StatusOK).JSON(fs.AllFriends)
}

// Delete godoc
// @Summary Delete a friendship
// @Description Remove a friendship between two users
// @Tags friendships
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param f_id query string true "Friend ID to remove"
// @Success 204 "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User not found"
// @Router /socio/friends/{id} [delete]
func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	fid := c.Query("f_id")
	friendID, err := uuid.Parse(fid)
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

	us.User = &dto.User{}
	us.User.ID = friendID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	fs := friendships.New()
	fs.UserID = userID
	fs.FriendID = friendID
	fs.Delete(ctx)

	if err := cache.Client().Del(ctx, userID.String()).Err(); err != nil {
		fmt.Println("Error invalidating cache: ", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
