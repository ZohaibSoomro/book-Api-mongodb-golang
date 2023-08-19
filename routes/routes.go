package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zohaibsoomro/book-server-mongodb/controller"
)

func RegisterRoutes(r *httprouter.Router) {
	r.GET("/books", controller.GetAllBooks)
	r.GET("/books/:id", controller.GetBookById)
	r.POST("/book", controller.CreateBook)
	r.PUT("/books/:id", controller.UpdateBook)
	r.DELETE("/book/:id", controller.DeleteBook)
}
