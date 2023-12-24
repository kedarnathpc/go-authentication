# Go JWT Authentication App

This Go application demonstrates how to implement JSON Web Token (JWT) authentication for securing your web services. 
It provides user registration, login, user profile retrieval, and user logout functionalities using Fiber and GORM.

## Features

- User registration with password hashing and unique email validation.
- User login with JWT token generation.
- User profile retrieval using JWT authentication.
- User logout by clearing JWT cookies.
- Structured codebase with separate controllers, models, and routes.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed on your system.
- MySQL database set up with the necessary configurations.
- Familiarity with JWT authentication and basic Go programming.

## Installation and Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/kedarnathpc/go-authentication.git
  
2. Install required Go packages:
   
   ```bash
   go mod tidy

3. Run the application:

    ```bash
    go run main.go

4. Your Go JWT authentication app should be running on http://localhost:8000.

## Usage
User Registration
To register a new user, make a POST request to /api/register with the following JSON data:

    {
      "name": "Your Name",
      "email": "your@email.com",
      "password": "your_password"
    }

## User Login
To log in, make a POST request to /api/login with the following JSON data:

    {
      "email": "your@email.com",
      "password": "your_password"
    }

## User Profile Retrieval
To retrieve the user profile, make a GET request to /api/user. You must include a valid JWT token in the request headers for authentication.

## User Logout
To log out, make a POST request to /api/logout. This will clear the JWT token from your browser cookies.

Contributing
Fork the repository.
Create a new branch for your feature: git checkout -b feature-name.
Make your changes and commit them: git commit -m 'Add feature'.
Push to the branch: git push origin feature-name.
Create a pull request.
  
