package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/voideus/go-rest-api/controller"
	"gitlab.com/voideus/go-rest-api/middlewares"
	"gitlab.com/voideus/go-rest-api/repository"
	"gitlab.com/voideus/go-rest-api/service"
)

var (
	videoRepository      repository.VideoRepository      = repository.NewVideoRepository()
	articleRepository    repository.ArticleRepository    = repository.NewArticleRepository()
	creditCardRepository repository.CreditCardRepository = repository.NewCreditCardRepo()
	personCardRepository repository.PersonRepo           = repository.NewPersonRepo()

	videoService      service.VideoService      = service.New(videoRepository)
	articleService    service.ArticleService    = service.NewArticleRepoService(articleRepository)
	creditCardService service.CreditCardService = service.NewCreditCardService(creditCardRepository)
	personService     service.PersonService     = service.NewPersonService(personCardRepository)
	loginService      service.LoginService      = service.NewLoginService()
	jwtService        service.JWTService        = service.NewJWTService()

	videoController      controller.VideoController      = controller.New(videoService)
	articleController    controller.ArticleController    = controller.NewArticleController(articleService)
	creditCardController controller.CreditCardController = controller.NewCreditCardController(creditCardService)
	personController     controller.PersonController     = controller.NewPersonController(personService)
	loginController      controller.LoginController      = controller.NewLoginController(loginService, jwtService)
)

func setLogOutput() {
	f, err := os.Create("gin.log")
	fmt.Println(err)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	defer videoRepository.CloseDB()
	setLogOutput()

	server := gin.New()

	// Serve static resources
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	// Middlewares
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	// Login: Authentication + Token Creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid"})
			}
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video updated successfully"})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video deleted successfully"})
			}
		})

		apiRoutes.POST("/articles", func(ctx *gin.Context) {
			err := articleController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Article added success."})
			}
		})

		apiRoutes.GET("/articles", func(ctx *gin.Context) {
			ctx.JSON(200, articleController.FindAll())
		})

		apiRoutes.GET("/articles/:id", articleController.FindArticle)
		apiRoutes.POST("/articles/:id/comments", articleController.AddCommentToArticle)

		apiRoutes.GET("/credit-cards", creditCardController.FindAll)
		apiRoutes.POST("/credit-cards", creditCardController.Save)

		apiRoutes.GET("/people", personController.FindAll)
		apiRoutes.GET("/people/:id/languages", personController.GetLanguagesOfPerson)
		apiRoutes.POST("/people/:id/languages", personController.AddLanguagesToPerson)
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", func(ctx *gin.Context) {
			videoController.ShowAll(ctx)
		})
	}

	server.Run()
}
