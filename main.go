package main

import (
    "github.com/capybara-brain346/go-jwt/controllers"
    "github.com/capybara-brain346/go-jwt/initializers"
    "github.com/gin-gonic/gin"
)

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
    initializers.SyncDatabase()
}

func main() {
    r := gin.Default()

    r.POST("/auth/signup", controllers.SignUp)
    r.POST("/auth/login", controllers.Login)

    auth := r.Group("/api")
    auth.Use(controllers.RequireAuth)
    {
        auth.GET("/user", controllers.GetUser)
        auth.PUT("/user", controllers.UpdateUser)
        auth.DELETE("/user", controllers.DeleteUser)
        auth.POST("/logout", controllers.Logout)
    }

    r.Run()
}
