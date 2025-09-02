package users

import (
	"socio/internals/dto"
	"socio/internals/notifications"
	"socio/internals/validator"
	"socio/services/users"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Add godoc
// @Summary Create a new user
// @Description Create a new user with name and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserCreate true "User creation data"
// @Success 201 {object} dto.User
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /socio/users [post]
func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var user dto.UserCreate

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}

	if err := validator.Payload(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}
	us := users.New()
	us.User = &dto.User{}

	us.User.Name = user.Name
	us.User.Email = user.Email
	us.User.Password = user.Password

	us.Create(ctx)

	notifications.Register(us.User.ID)

	go notifications.ListenForNotifications(ctx, us.User.ID)

	return c.Status(fiber.StatusCreated).JSON(us.User)
}

// Get godoc
// @Summary Get user by ID
// @Description Get a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.User
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User not found"
// @Router /socio/users/{id} [get]
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

	return c.Status(fiber.StatusOK).JSON(us.User)
}

// Delete godoc
// @Summary Delete user by ID
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User not found"
// @Router /socio/users/{id} [delete]
func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("incorrect user id")
	}

	us := users.New()
	us.User = &dto.User{}

	us.User.ID = userID

	if err := us.Delete(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetAll godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} dto.User
// @Failure 500 {string} string "Internal Server Error"
// @Router /socio/users [get]
func GetAll(c *fiber.Ctx) error {
	ctx := c.UserContext()

	us := users.New()
	us.Users = &dto.Users{}

	if err := us.GetAll(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(us.Users)
}
