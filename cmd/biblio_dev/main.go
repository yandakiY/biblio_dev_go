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

		routerAuteur.POST("" , func(ctx *gin.Context) {

			res , err := auteurController.Create(ctx)

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
					"items": res,
				})
			}
		})

		routerAuteur.PUT("/:id", func(ctx *gin.Context) {
			id , _ := strconv.ParseUint(ctx.Param("id"), 10 , 32)

			res , err := auteurController.Update(uint(id) , ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":http.StatusBadRequest,
					"message":err.Error(),
					"items":nil,
				})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{
					"status":http.StatusOK,
					"message":"Operation succesfull",
					"items": res,
				})
			}
		})

		routerAuteur.DELETE("/:id", func(ctx *gin.Context) {
			
			id , _ := strconv.ParseUint(ctx.Param("id"), 10 , 32)

			err := auteurController.Delete(uint(id))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":http.StatusBadRequest,
					"message":err.Error(),
					"items":nil,
				})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{
					"status":http.StatusOK,
					"message":"Operation succesfull",
					"items": nil,
				})
			}
		})
	}

	routerLivre := server.Group("/api/livre")
	{	
		// Auteur - Endpoint
		routerLivre.GET("" , func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK , gin.H{
				"message":"Get Endpoint /api",
			})
		})

		routerLivre.POST("" , func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message":"Post Endpoint /api",
			})
		})

		routerLivre.PUT("/:id", func(ctx *gin.Context) {

			id , _ := strconv.ParseUint(ctx.Param("id") , 10, 32)
			print(id)
			ctx.JSON(http.StatusOK, gin.H{
				"message":"Update Endpoint /api",
			})
		})

		routerLivre.DELETE("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message":"Delete Endpoint /api",
			})
		})
	}

	server.Run(":8080")

}