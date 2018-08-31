# Ai Competion Server

Thie System uses revel
A high-productivity web framework for the [Go language](http://www.golang.org/).
 
### This is Score Server System

```
├── README.md
├── app
│   ├── controllers
│   │   └── api
│   │       └── v1
│   │           ├── auth.go
│   │           ├── response.go
│   │           └── user.go
│   ├── db
│   │   └── gorm.go
│   ├── init.go
│   ├── models
│   │   ├── base.go
│   │   └── user.go
│   ├── routes
│   │   └── routes.go
│   ├── tmp
│   │   └── main.go
│   └── views
├── conf
│   ├── app.conf
│   └── routes
└── tests
    └── apptest.go
```

### API LIST:

```
METHOD  URL                                     Function                         Allow Role
// app/controllers/api/v1/user.go
GET     /api/v1/user                            ApiUser.Index                    Admin
GET     /api/v1/user/:id                        ApiUser.Show                     Admin, Signed in User
PUT     /api/v1/user/:id                        ApiUser.Update                   Admin, Signed in User
DELETE  /api/v1/user/:id                        ApiUser.Delete                   Admin, Signed in User

POST    /api/v1/signup                          ApiUser.Create                   Everyone
// app/controllers/api/v1/auth.go
GET     /api/v1/signin                          ApiAuth.GetSessionID             Everyone
POST    /api/v1/signin                          ApiAuth.SignIn                   Everyone
DELETE  /api/v1/signin                          ApiAuth.SignOut                  Signed in
GET     /api/v1/role                            ApiAuth.Role                     Everyone
```
