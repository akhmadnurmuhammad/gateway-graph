package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	cf "xti-gateway-graph-go/config"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"xti-gateway-graph-go/graph"
	"xti-gateway-graph-go/graph/resolver"
)

var (
	conf = cf.Config{}
	d    struct{}
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	// setting up configuration
	conf.Init()
	// middleware here

	server := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"Status":    fiber.StatusInternalServerError,
					"IsSuccess": false,
					"Message":   err.Error(),
					"Data":      d,
				})
			},
			AppName: conf.Service.Name,
		},
	)

	// register handler
	Handler(server)

	port := "8080"
	if conf.Service.Port != 0 {
		port = strconv.Itoa(conf.Service.Port)
	}
	return server.Listen(":" + port)
}

func Handler(route *fiber.App) {

	// middleware
	route.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} [${time}] ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))

	route.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType",
	}))

	// health check
	route.Get("health-check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"IsSuccess": true,
			"Message":   fmt.Sprintf("SVC RUNNING %s - Version %s", conf.Service.Name, conf.Service.Version),
			"Data":      d,
			"Status":    "0",
		})
	})

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &resolver.Resolver{},
			},
		),
	)

	gqlHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})

	playground := playground.Handler("GraphQL playground", "/query")

	route.All("/query", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(gqlHandler)(c.Context())
		return nil
	})

	route.All("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playground)(c.Context())
		return nil
	})

}
