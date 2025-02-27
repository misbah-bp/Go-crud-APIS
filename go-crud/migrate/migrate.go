package main

import (
	"github.com/misbah-bp/go-crud/initializers"
	"github.com/misbah-bp/go-crud/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {

	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Blog{})
}
