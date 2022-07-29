# apis
## create a user
POST /users/ <- { username, password, email }

## get all users list
GET /users -> { users: []{ user_id, username, email } }

## get a user
GET /users/:userID -> { username, email }

# 3 type of variables in a HTTP request
- uri variable: /users/:userID --> /users/1
- query variable: /users?userID=1
- body variable: /users <- { userID : 1 }

# 资源
- https://gorm.io/
- https://github.com/gin-gonic/gin