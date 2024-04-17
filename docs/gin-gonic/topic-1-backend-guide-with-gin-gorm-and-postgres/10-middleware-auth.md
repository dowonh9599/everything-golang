# 10 Middleware AUTH

## 10.1 Refactor Routes

At this point with multiple routes for different controllers added, we can try to refactor routes as follows

```
- ...
- ...
- ...
- routes
|   ㄴ book.routes.go
|   ㄴ file.routes.go
|   ㄴ index.routes.go
|   ㄴ user.routes.go
- ...
- main.go
```

```go
// routes/book.routes.go
func SetupBookRouter(app *gin.Engine) {
    app.GET("/books", book_controller.GetAllBooks)
}
```

```go
// routes/file.routes.go
func SetupFileRouter(app *gin.Engine) {
    fileRoute := app.Group("file")
    
    // file controllers
    fileRoute.POST("/", file_controller.HandleUploadFile)
    fileRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
}
```

```go
// routes/user.routes.go
func SetupUserRouter(app *gin.Engine) {
    userRoute := app.Group("user")
    
    app.GET("/users", user_controller.GetAllUsers)
    userRoute.GET("/:id", user_controller.GetUserById)
    userRoute.GET("/paginate", user_controller.GetUsersPaginate) // /user/paginate?perPage=2*page=2
    userRoute.POST("/", user_controller.AddNewUser)
    userRoute.PATCH("/:id", user_controller.UpdateById)
    userRoute.DELETE("/:id", user_controller.DeleteById)
}
```

```go
func SetupRouter(app *gin.Engine) {
    app.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
    // Ping test
    app.GET("/ping", ping.GetPong)
    
    SetupUserRouter(app)
    SetupBookRouter(app)
    SetupFileRouter(app)
}
```

## 10.2 Setting up Middleware

create a new folder in the root directory `middleware`. Create a middleware for authentication at `middleware/auth.middleware.go`.

```go
// middleware/auth.middleware.go
func AuthMiddleware(c *gin.Context) {
    token := c.GetHeader("X-Token")
    
    if token == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            "message": "token must be provided",
        })
        
        return
    }
    
    c.Next()
}
```

Modify file routes in `routes/file.routes.go` so that it always goes through Authentication middleware

```go
// routes/file.routes.go
func SetupFileRouter(app *gin.Engine) {
    authRoute := app.Group("file", middleware.AuthMiddleware)
    
    // file controllers
    authRoute.POST("/", file_controller.HandleUploadFile)
    authRoute.POST("/middleware", middleware.UploadFile, file_controller.SendStatus)
    authRoute.DELETE("/:filename", file_controller.HandleRemoveFile)
}
```

### 11.3 (Optional) Rewrite HandleUploadFile feature with middleware

```go
// middleware/file.middleware.go
func UploadFile(c *gin.Context) {
    fileHeader, _ := c.FormFile("file")
    
    allowedFileTypes := []string{"image/png", "image/jpg", "application/pdf"}
    isValid := utils.ValidateFile(fileHeader, allowedFileTypes)
    if !isValid {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "file type now allowed",
        })
        
        return
    }
    
    fileExt := filepath.Ext(fileHeader.Filename)
    filename := utils.GetRandomFilename(fileExt)
    
    errSave := utils.SaveFile(c, fileHeader, &filename)
    if errSave != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "message": errSave.Error(),
        })
        
        return
    }
    
    c.Set("filename", filename)
    
    c.Next()
}
```

```go
// controllers/file_controller/index.file_controller.go
func SendStatus(c *gin.Context) {
    filename := c.MustGet("filename").(string)
    
    c.JSON(http.StatusOK, gin.H{
        "message":   "file uploaded",
        "file_name": filename,
    })
}
```

```go
// routes/file.routes.go
func SetupFileRouter(app *gin.Engine) {
    authRoute := app.Group("file", middleware.AuthMiddleware)
    
    // ...
    authRoute.POST("/middleware", middleware.UploadFile, file_controller.SendStatus)
}
```