package main

import (
	"fmt"

	"github.com/Merphyy/go-rest/book"
	"github.com/Merphyy/go-rest/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func helloWorld(c *fiber.Ctx){
	c.Send("hello world")
}

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Fail to connect database")
	}
	fmt.Println("Database successfully connected!")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated!")
	  
}

func main(){
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	
	app.Listen(3000)
}