![ci workflow](https://github.com/khihadysucahyo/go-echo-boilerplate/actions/workflows/ci.yml/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/b55f79729435efcc46f9/maintainability)](https://codeclimate.com/github/khihadysucahyo/go-echo-boilerplate/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b55f79729435efcc46f9/test_coverage)](https://codeclimate.com/github/khihadysucahyo/go-echo-boilerplate/test_coverage)

# Overview
It's an API project based on Echo framework.

## What's inside:

- Registration
- Authentication with JWT
- CRUD API for posts
- Migrations
- Request validation
- Swagger docs
- Environment configuration
- Docker development environment

## Usage
1. Copy .env.dist to .env and set the environment variables. There are examples for all the environment variables except COMPOSE_USER_ID, COMPOSE_GROUP_ID which are used by the linter. To get the current user ID, run in terminal:
    
    `echo $UID`
    
    In the .env file set these variables:

    `COMPOSE_USER_ID="username in current system"` - your username in system

    `COMPOSE_GROUP_ID="user uid"` - the user ID which you got in the terminal

2. Run your application using the command in the terminal:

    `docker-compose up`
3. Browse to {HOST}:{PORT}/swagger/index.html. You will see Swagger 2.0 API documents.
4. Using the API documentation, make requests to register a user (if necessary) and login.
5. After the successful login, copy a token from the response, then click "Authorize" and in a popup that opened, enter the value for "apiKey" in a form:
"Bearer {token}". For example:


    Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk0NDA5NjYsIm9yaWdfaWF0IjoxNTg5NDM5OTY2LCJ1c2VyX2lkIjo1fQ.f8dSG3NxFLHwyA5-XIYALT5GtXm4eiH-motqtqAUBOI 

   
Then, click "Authorize" and close the popup.
Now, you are able to make requests which require authentication.

## Directories
1. **/cmd** entry points.

2. **/config** has structures which contains service config.

3. **/db** has seeders and method for connecting to the database.

4. **/migrations** has files for run migrations.

5. **/models** includes structures describing data models.

6. **/repositories** contains methods for accesing entities from the database.

7 **/requests** has structures describing the parameters of incoming requests, and validator.

8. **/responses** includes structures describing the parameters of outgoing response.

9. **/server** is the main project folder. This folder contains the executable server.go.

10. **/server/builders** contains builders for initializing entities.

11. **/server/handlers** contains request handlers.

12. **/server/routes** has a file for configuring routes.

### How to use
Linter tool wrapped to docker-compose and first of all need to build container with linters

- `make run`

To run project by makeFile

- `make migrate`

## Libraries
Migrations - https://github.com/ShkrutDenis/go-migrations

Jwt - https://github.com/dgrijalva/jwt-go

Swagger - https://github.com/swaggo/echo-swagger

Gomock - https://github.com/golang/mock

Orm - https://github.com/jinzhu/gorm
