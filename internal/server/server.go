package server

import (
	"fmt"
	"seen/internal/controllers"
	"seen/internal/database"
	"seen/internal/pkg"
	"seen/internal/repository"
	"seen/internal/routes"
	"seen/internal/services"
	"seen/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RunServer() {
	// Define variables for server configuration obtained from environment variables.
	var (
		listenAddress = utils.String("listen-address", "Server listen address")
		dbHost        = utils.String("db-host", "Database host address") // Define a variable for the database host address.
		dbPort        = utils.Int("db-port", "Database port")            // Define a variable for the database port.
		dbName        = utils.String("db-name", "Database name")         // Define a variable for the database name.
		dbUsername    = utils.String("db-username", "Database username") // Define a variable for the database username.
		dbPassword    = utils.String("db-password", "Database password")
	)

	// Handle configuration errors.
	if err := utils.Parse(); err != nil {
		utils.PanicMissingEnvParams(err.Error()) // Log an error if there are missing environment parameters.
	}

	// Connect to the PostgreSQL database.
	db, err := database.ConnectToPostgres(*dbHost, *dbPort, *dbName, *dbUsername, *dbPassword)
	if err != nil {
		utils.PanicDBConnectionFailed(err.Error()) // Log an error if the database connection fails.
	}

	var (
		repository repository.SeenRepository = repository.NewSeenRepository(db)

		jwtService pkg.JWTService = pkg.NewJWTService()

		authService services.AuthService = services.NewAuthService(repository, jwtService)

		authController controllers.AuthController = controllers.NewAuthController(authService)
	)

	// Create a new Fiber instance.
	app := fiber.New()

	// Use CORS middleware for handling cross-origin requests.
	app.Use(cors.New())

	v1 := app.Group("/api")

	routes.AuthGroup(v1, authController)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})

	// Start the Fiber server and log any errors encountered during startup.
	err = app.Listen(*listenAddress)
	if err != nil {
		fmt.Println(err)
	}
}
