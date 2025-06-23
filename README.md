# Secure-API-Gateway-in-GoLang ->  You're building a secure REST API in GoLang using the Fiber framework and JWT (JSON Web Tokens) to:

Issue a token when user logs in
Protect routes like /admin and /user
Allow access only if JWT is valid and role matches.

 Why?
go mod init: Initializes Go module for dependency management

go get: Downloads Fiber (web framework) & JWT lib


---------------------------------

 Goal of the Project

--------------------------------

We are building a secure backend API that:
Accepts login input (username, role)
Generates a JWT token
Has protected routes like /user, /admin
Only allows access based on the role in the token

-------------------------------

 Tools we’ll Use
Tool	Why
------------------------------
GoLang	Backend language
Fiber	Fast web framework (like Express for Go)
JWT (via golang-jwt)	To generate and validate tokens
Postman/curl	For testing your endpoints.

------------------------------------

 Project Logic:

------------------------------------

1. User sends login request with username and role
2. Server creates a JWT token with that info (username, role, expiry)
3. Server returns the token
4. When user accesses /user or /admin, they must send token
5. Middleware checks the token:

Is it valid?
Is the role allowed?

If YES → allow
If NO → block

------------------------------------

 Features You'll Build:
Feature	Purpose
----------------------------------

/login endpoint	Takes username & role → returns JWT
/admin endpoint	Only accessible if token has role = admin
/user endpoint	Accessible by all roles
JWT middleware	Validates token, extracts claims
Code structure	Modular: routes, handlers, middleware

Let's start to build:
 Step 1: Create Folder and init project : mkdir secure-api-gateway-go
                         cd secure-api-gateway-go
                         go mod init secure-api-gateway-go

 Step 2: Install Dependencies: 
 go get github.com/gofiber/fiber/v2
 go get github.com/golang-jwt/jwt/v5

Real-world uses:
This is how Google, Cisco, Amazon, Microsoft, etc. build stateless secure backend APIs.

Code Part Covered:

package main
import ("github.com/gofiber/fiber/v2"
	"secure-api-gateway-go/handlers"
	"secure-api-gateway-go/middleware")

Import	    What it gives you	               Why it matters
fiber/v2	  Web framework	                 To create routes and run the server
handlers	  Login + page handlers	         Clean separation of login + routes
middleware	Token + role check logic	    Secure each route with role-based access

Step-by-Step Explanation:

Code Line	  What it Does	                            Why it Matters
func main()	The entry point of the Go application	Every Go executable must start here
app := fiber.New()	Initializes the web server using Fiber	Just like Express in Node.js
app.Post("/login", handlers.Login)	Defines a route to accept login data	Accepts username & role, and returns a JWT
app.Get("/user", middleware.JWTMiddleware("user"), handlers.UserPage)	Defines a protected route for users	Token & role "user" required
app.Get("/admin", middleware.JWTMiddleware("admin"), handlers.AdminPage)	Protected route for admin access	Token & role "admin" required
app.Listen(":3000")	Starts the web server on port 3000	App runs at http://localhost:3000



 

 
