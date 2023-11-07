package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/kedarnathpc/go-auth/database"
	"github.com/kedarnathpc/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

// SecretKey is the secret key used for JWT token signing.
const SecretKey = "secret"

// Register handles user registration.
func Register(c *fiber.Ctx) error {
	// Parse the request body to retrieve user registration data.
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Hash the user's password before storing it.
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	// Create a new user model with the registration data.
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	// Create the user in the database.
	database.DB.Create(&user)

	// Return the created user as a JSON response.
	return c.JSON(user)
}

// Login handles user login and JWT token generation.
func Login(c *fiber.Ctx) error {
	// Parse the request body to retrieve login data.
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Query the database to find the user with the provided email.
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	// Check if the user exists in the database.
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Compare the provided password with the hashed password from the database.
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// Create a JWT token for the authenticated user.
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	// Check for errors while creating the token.
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	// Set the JWT token as a cookie and return a success message.
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// User retrieves user data based on the JWT token.
func User(c *fiber.Ctx) error {
	// Retrieve the JWT token from the request cookie.
	cookie := c.Cookies("jwt")

	// Parse the JWT token with the specified secret key.
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// Check for token parsing errors.
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// Extract user claims from the token.
	claims := token.Claims.(*jwt.StandardClaims)

	// Query the database to find the user based on the JWT issuer (user ID).
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	// Return the user data as a JSON response.
	return c.JSON(user)
}

// Logout clears the JWT cookie to log out the user.
func Logout(c *fiber.Ctx) error {
	// Clear the JWT cookie to log the user out.
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	// Return a success message.
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
