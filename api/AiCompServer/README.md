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
│   │           ├── answer.go
│   │           ├── auth.go
│   │           ├── challenge.go
│   │           ├── notification.go
│   │           ├── ranking.go
│   │           ├── response.go
│   │           ├── setting.go
│   │           └── user.go
│   ├── db
│   │   └── gorm.go
│   ├── init.go
│   ├── models
│   │   ├── answer.go
│   │   ├── base.go
│   │   ├── challenge.go
│   │   ├── notification.go
│   │   ├── setting.go
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
// AiCompServer/app/controllers/api/v1/user.go
GET     /api/v1/user                            ApiUser.Index                    Admin
GET     /api/v1/user/:id                        ApiUser.Show                     Admin, Signed in User
PUT     /api/v1/user/:id                        ApiUser.Update                   Admin, Signed in User
DELETE  /api/v1/user/:id                        ApiUser.Delete                   Admin, Signed in User

POST    /api/v1/signup                          ApiUser.Create                   Everyone
// AiCompServer/app/controllers/api/v1/auth.go
GET     /api/v1/signin                          ApiAuth.GetSessionID             Everyone
POST    /api/v1/signin                          ApiAuth.SignIn                   Everyone
DELETE  /api/v1/signin                          ApiAuth.SignOut                  Signed in
GET     /api/v1/role                            ApiAuth.Role                     Everyone

// AiCompServer/app/controllers/api/v1/ranking.go
GET    /api/v1/ranking                          ApiChallenge.Ranking             Signed in
GET    /api/v1/ranking/:id                      ApiChallenge.RankingChallenge    Signed in

// AiCompServer/app/controllers/api/v1/challenge.go
GET    /api/v1/challenges                       ApiChallenge.Index               Signed in
POST   /api/v1/challenges                       ApiChallenge.Create              Admin
POST   /api/v1/challengesfile                  ApiChallenge.ChallengeFileUpload  Admin
GET    /api/v1/challenges/:id                   ApiChallenge.Show                Signed in
PUT    /api/v1/challenges/:id                   ApiChallenge.Update              Admin
DELETE /api/v1/challenges/:id                   ApiChallenge.Delete              Admin

// AiCompServer/app/controllers/api/v1/answer.go
GET    /api/v1/answers                          ApiAnswer.Index                  Admin
POST   /api/v1/answers                          ApiAnswer.Create                 Admin
GET    /api/v1/answers/:id                      ApiAnswer.Show                   Admin
PUT    /api/v1/answers/:id                      ApiAnswer.Update                 Admin
DELETE /api/v1/answers/:id                      ApiAnswer.Delete                 Admin

GET     /api/v1/challengeanswer/:id             ApiAnswer.UserChallengeAnswer    Signed in User

POST   /api/v1/submit                           ApiAnswer.Submit                 Signed in

// AiCompServer/app/controllers/api/v1/notification.go
GET     /api/v1/notifications                    ApiNotification.Index           Everyone
POST    /api/v1/notifications                    ApiNotification.Create          Admin
GET     /api/v1/notifications/:id                ApiNotification.Show            Everyone
PUT     /api/v1/notifications/:id                ApiNotification.Update          Admin
DELETE  /api/v1/notifications/:id                ApiNotification.Delete          Admin

// AiCompServer/app/controllers/api/v1/setting.go
GET     /api/v1/setting                         ApiSetting.Index                 Admin
POST    /api/v1/setting                         ApiSetting.Create                Admin
GET     /api/v1/setting/:id                     ApiSetting.Show                  Admin
PUT     /api/v1/setting/:id                     ApiSetting.Update                Admin
DELETE  /api/v1/setting/:id                     ApiSetting.Delete                Admin
GET     /api/v1/setting/adapt/:id               ApiSetting.Adapt                 Admin
```
