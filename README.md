# Secure-API-Gateway-in-GoLang ->  You're building a secure REST API in GoLang using the Fiber framework and JWT (JSON Web Tokens) to:

Issue a token when user logs in
Protect routes like /admin and /user
Allow access only if JWT is valid and role matches.

What I Learnt:

-> How to build a web server in Go using Fiber
-> How to generate and validate JWT tokens
-> How to use middleware to protect routes
-> How to enforce role-based access
-> How to test real-world APIs using Postman
-> How modern cloud systems manage stateless authentication

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

-> Testing JWT-Based Secure API with Postman:
This guide shows how to test a real-world secure API using GoLang, JWT, and Fiber in Postman.
It includes explanations of each step, why we use certain HTTP methods like POST, how headers work, and how to test protected routes securely.

Why Use Postman?
Postman lets you simulate how a frontend app or client talks to a backend API.

In real life, systems like:
Cisco Secure Access
 AWS IAM
Google APIs
-> use secure HTTP requests with tokens, headers, and JSON — and Postman helps you test that as a developer.

we're testing backend API

Route	Method	Description
/login	POST	Accepts user info, returns JWT
/admin	GET	Protected admin route
/user	GET	Protected user route

STEP 1: Test the /login Route (to get JWT Token)
Goal:
Send login info → receive JWT token in response.

Why POST?
POST is used when you're sending data to the server.
In this case, we send a JSON body with username and role.

How To Do It in Postman
Setting	Value
Method	POST
URL	http://localhost:3000/login
Body type	raw
Content format	JSON (application/json)
Body content	Paste this:

{
  "username": "komal",
  "role": "admin"
}

Why JSON?
JSON is the standard format used by most APIs today.
It's structured, lightweight, and easy for Go (c.BodyParser) to read.
The key-value pairs (username, role) are mapped to a struct in Go:

Note : JSON must match the field names.

Header Setup
Go to the Headers tab and ensure this is added:
Key	Value
Content-Type	application/json
Why? -> This tells the server: “The body I’m sending is JSON. Please parse it as JSON.

If successful, you’ll get the token

STEP 2: Test the /admin Route (using the token)
Goal:
Access a protected route using the token you got from /login

How To Do It
Setting	Value
Method	GET
URL	http://localhost:3000/admin

Headers
Go to Headers tab and set:
Key	Value
Authorization	Bearer <your_token_here>

Note: Make sure there is a space between Bearer and your token.

Why Use Bearer Token in Header?
This simulates how real browsers and mobile apps send tokens.
Authorization: Bearer <token> is the standard header for secure API access.

Your middleware in Go reads this header like:
authHeader := c.Get("Authorization")

Response: 
If your token has "role": "admin":
Welcome, admin!

If your token has "role": "user":
403 Forbidden: role mismatch

This proves your role-based access logic is working perfectly.

STEP 3: Try Wrong Role
Go back to /login
Change role to "user" in body:

{
  "username": "komal",
  "role": "user"
}

Note: Use the new token on /admin with GET
Result : 403 Forbidden: role mismatch

This shows you’ve built secure APIs that only allow access if the JWT token has the correct role.











 

 
