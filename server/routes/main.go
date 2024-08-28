package routes

import (
	adminRoutes "api/routes/admin"
	loginRoutes "api/routes/login"
	rankRoutes "api/routes/ranks"
	testRoutes "api/routes/test"
	userRoutes "api/routes/users"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func SetupRoutes(app *fiber.App) {

	//Preparing dotenv files
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err)
	}

	//Setup admin routes
	adminRoutes.SetupAdminRoutes(app)

	//Setup login/logout routes
	loginRoutes.SetupLoginRoutes(app)

	//Setup rank routes
	rankRoutes.SetupRankRoutes(app)

	//Setup test routes
	testRoutes.SetupTestRoutes(app)

	//Setup user routes
	userRoutes.SetupUserRoutes(app)
}
