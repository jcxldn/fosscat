package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/graph/resolver"
	"github.com/jcxldn/fosscat/backend/util/jwt"
)

// Define the Graphql route handler
// based on https://gqlgen.com/recipes/gin/
func graphqlHandler(resolver *resolver.Resolver) gin.HandlerFunc {

	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Define the Playground route handler
// based on https://gqlgen.com/recipes/gin/
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Define the Ping route handler
func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Main function
func main() {
	// Connect to the database and place the handle in the graphql resolver
	// So that is accessible when executing graphql requests in ctx.
	resolver := &resolver.Resolver{}
	resolver.UpdateDb(database.Connect())

	// Setup JWT keys (load from file)
	jwt.SetupKey()

	// Create a gin engine instance
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type"},
	}))

	r.Use(AuthMiddleware(resolver.DB))

	// Define routes
	r.POST("/graphql", graphqlHandler(resolver))      // GraphQL query endpoint
	r.GET("/graphql/playground", playgroundHandler()) // GraphiQL playground
	r.GET("/ping", pingHandler)                       // Ping/pong endpoint (for healthcheck)

	r.Run() // Run on 0.0.0.0:8080
}
