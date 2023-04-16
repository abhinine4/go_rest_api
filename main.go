package main

import (
	"fmt"
	"go-fiber-crud/config"
	"go-fiber-crud/controller"
	"go-fiber-crud/model"
	"go-fiber-crud/repository"
	"go-fiber-crud/router"
	"go-fiber-crud/service"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Print("Start service ... ")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variable", err)
	}

	db := config.ConnectionDB(&loadConfig)

	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	noteRepository := repository.NewNoteRepositoryImpl(db)

	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	noteController := controller.NewNoteController(noteService)

	routes := router.NewRouter(noteController)

	app := fiber.New()
	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8000"))
}
