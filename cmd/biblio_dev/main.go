package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yandakiY/biblio_dev_go/internal/controller"
	"github.com/yandakiY/biblio_dev_go/internal/repository"
	"github.com/yandakiY/biblio_dev_go/internal/service"
)

var (
	// Injection Repositories
	auteurRepo repository.AuteurRepository = repository.NewAuteurRepostitory()
	livreRepo repository.LivreRepository = repository.NewLivreRepository()

	// injection Services
	auteurService service.AuteurService = service.NewAuteurService(auteurRepo)
	livreService service.LivreService = service.NewLivreService(livreRepo)

	// injection controller 
	auteurController controller.AuteurController =  controller.NewAuteurController(auteurService)
	livreController controller.LivreController = controller.NewLivreController(livreService)

)

func main(){
	fmt.Println("Biblio_dev Application...")

	server := gin.New()

	routerAuteur := server.Group("/api/auteur")
	{	
		// Auteur - Endpoint
		routerAuteur.GET("" , func(ctx *gin.Context) {
			livres := auteurController.Get()
			if len(livres) == 0 {
				ctx.JSON(http.StatusOK , gin.H{
					"status": http.StatusOK,
					"items":nil,
				})
			} else {
				ctx.JSON(http.StatusOK , gin.H{
					"status": http.StatusOK,
					"items": livres,
				})
			}
		})

		routerAuteur.GET("/:id" , func(ctx *gin.Context) {

			id, _ := strconv.ParseUint(ctx.Param("id"), 10 ,32)

			livre , err := auteurController.FindById(uint(id))
			if err != nil {
				ctx.JSON(http.StatusNotFound , gin.H{
					"status": http.StatusNotFound,
					"error":err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK , gin.H{
					"status": http.StatusOK,
					"items": livre,
				})
			}
		})

		routerAuteur.POST("" , func(ctx *gin.Context) {

			_ , err := auteurController.Create(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":http.StatusBadRequest,
					"message":err.Error(),
					"items":nil,
				})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{
					"status":http.StatusCreated,
					"message":"Operation succesfull",
				})
			}
		})

		routerAuteur.PUT("/:id", func(ctx *gin.Context) {
			id , _ := strconv.ParseUint(ctx.Param("id"), 10 , 32)

			res , err := auteurController.Update(uint(id) , ctx)

			if err != nil && err.Error() == "Auteur a modifier inexistant" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":http.StatusNotFound,
					"message":err.Error(),
				})
			} else if err != nil && err.Error() != "Auteur a modifier inexistant" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":http.StatusBadRequest,
					"message":err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"status":http.StatusOK,
					"message":"Operation succesfull",
					"items": res,
				})
			}
		})

		routerAuteur.DELETE("/:id", func(ctx *gin.Context) {
			
			id , _ := strconv.ParseUint(ctx.Param("id"), 10 , 32)

			err := auteurController.Delete(uint(id))

			if err != nil && err.Error() == "Auteur inexistant" {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status":http.StatusNotFound,
					"message":err.Error(),
				})
			} else if err != nil && err.Error() != "Auteur inexistant" {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status":http.StatusInternalServerError,
					"message":err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"status":http.StatusOK,
					"message":"Operation succesfull",
				})
			}
		})
	}

	routerLivre := server.Group("/api/livre")
	{	
		// livre - Endpoint

		routerLivre.GET("" , func(ctx *gin.Context) {
			livres := livreController.Get()

			if len(livres) != 0{
				ctx.JSON(http.StatusOK , gin.H{
					"items":livres,
					"status":http.StatusOK,
				})
			} else {
				ctx.JSON(http.StatusOK , gin.H{
					"items":nil,
					"status":http.StatusOK,
				})
			}
		})

		routerLivre.GET("/:id" , func(ctx *gin.Context) {

			id , _ := strconv.ParseUint(ctx.Param("id"), 10 , 32)
			res , err := livreController.FindById(uint(id))

			if err != nil {
				ctx.JSON(http.StatusNotFound , gin.H{
					"error":err.Error(),
					"status":http.StatusNotFound,
				})
			} else {
				ctx.JSON(http.StatusOK , gin.H{
					"items":res,
					"status":http.StatusOK,
				})
			}
		})

		routerLivre.POST("" , func(ctx *gin.Context) {

			_ , err := livreController.Create(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message":err.Error(),
					"status":http.StatusBadRequest,
				})
			} else {
				ctx.JSON(http.StatusOK , gin.H{
					"message":"Operation succesfull",
					"status":http.StatusOK,
				})
			}
		})

		routerLivre.PUT("/:id", func(ctx *gin.Context) {

			print("Update Livre with Id")

			id , _ := strconv.ParseUint(ctx.Param("id") , 10, 32)
			
			res , err := livreController.Update(uint(id) , ctx)
			
			if err != nil && err.Error() == "Livre a modifier inexistant" {
				ctx.JSON(http.StatusNotFound, gin.H{
					"err": err.Error(),
					"status":http.StatusNotFound,
				})
			} else if err != nil && err.Error() != "Livre a modifier inexistant" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":err.Error(),
					"status":http.StatusBadRequest,
				})
			}else {
				ctx.JSON(http.StatusOK, gin.H{
					"message":"Operation succesfull",
					"status":http.StatusOK,
					"items": res,
				})
			}
		})

		routerLivre.DELETE("/:id", func(ctx *gin.Context) {

			id , err := strconv.ParseUint(ctx.Param("id"), 10 ,32)

			err = livreController.Delete(uint(id))
			
			if err != nil && err.Error() == "Livre inexistant" {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
					"status":http.StatusNotFound,
				})
			} else if err != nil && err.Error() != "Livre inexistant" {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
					"status":http.StatusInternalServerError,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message":"Operation Successfull",
					"status":http.StatusOK,
				})
			}
		})
	}

	server.Run(":8080")

}