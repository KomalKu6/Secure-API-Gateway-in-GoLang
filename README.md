# Secure-API-Gateway-in-GoLang ->  You're building a secure REST API in GoLang using the Fiber framework and JWT (JSON Web Tokens) to:

Issue a token when user logs in
Protect routes like /admin and /user
Allow access only if JWT is valid and role matches.

✅ Why?
go mod init: Initializes Go module for dependency management

go get: Downloads Fiber (web framework) & JWT lib


---------------------------------

🎯 Goal of the Project

--------------------------------

We are building a secure backend API that:
Accepts login input (username, role)
Generates a JWT token
Has protected routes like /user, /admin
Only allows access based on the role in the token

-------------------------------

🔧 Tools we’ll Use
Tool	Why
------------------------------
GoLang	Backend language
Fiber	Fast web framework (like Express for Go)
JWT (via golang-jwt)	To generate and validate tokens
Postman/curl	For testing your endpoints.

------------------------------------

🧠 Project Logic:

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

🧱 Features You'll Build:
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
 

 
