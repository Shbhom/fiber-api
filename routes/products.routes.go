package routes

import "github.com/shbhom/fiber-api/models"

type responseProduct struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Price uint   `json:"price"`
}

func createResponseProduct(prod models.Product) {

}
