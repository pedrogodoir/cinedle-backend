package router

import (
	"cinedle-backend/internal/config"
	classic_game_handler "cinedle-backend/internal/modules/classicGame/handlers"
	classic_game_router "cinedle-backend/internal/modules/classicGame/routes"
	movie_handler "cinedle-backend/internal/modules/movies/handlers"
	movie_router "cinedle-backend/internal/modules/movies/routes"
	movie_service "cinedle-backend/internal/modules/movies/services"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.LoadConfig()
	r := gin.Default()
	r.Use(cors.Default())

	// Serve o arquivo OpenAPI gerado para facilitar uso com Swagger UI
	// acessível em: GET /openapi.yaml
	r.GET("/openapi.yaml", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Content-Type", "application/x-yaml")
		c.File("./openapi.yaml")
	})

	// Página simples com Swagger UI (via CDN) em /docs
	r.GET("/docs", func(c *gin.Context) {
		html := `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8" />
		<title>Cinedle API Docs</title>
		<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@4/swagger-ui.css" />
		<style>body { margin:0; padding:0; }</style>
	</head>
	<body>
		<div id="swagger-ui"></div>
		<script src="https://unpkg.com/swagger-ui-dist@4/swagger-ui-bundle.js"></script>
		<script>
			window.onload = function() {
				const ui = SwaggerUIBundle({
					url: '/openapi.yaml',
					dom_id: '#swagger-ui',
					presets: [SwaggerUIBundle.presets.apis],
					layout: 'BaseLayout'
				})
			}
		</script>
	</body>
</html>`

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})
	//setup routes
	movieService := movie_service.NewMoviesService()
	movieHandler := movie_handler.NewMoviesHandler(movieService)
	movie_router.Routes(r, movieHandler)
	classicGameHandler := classic_game_handler.NewClassicGameHandler()
	classic_game_router.Routes(r, classicGameHandler)

	r.Run(":" + cfg.Port)
}
