package main

import (
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/graph/resolver"
)

// Define the Graphql route handler
// based on https://gqlgen.com/recipes/gin/
func graphqlHandler() gin.HandlerFunc {
	// Connect to the database and place the handle in the graphql resolver
	// So that is accessible when executing graphql requests in ctx.
	resolver := &resolver.Resolver{}
	resolver.UpdateDb(database.Connect())

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
	// Create a gin engine instance
	r := gin.Default()

	// Define routes
	r.POST("/graphql", graphqlHandler())              // GraphQL query endpoint
	r.GET("/graphql/playground", playgroundHandler()) // GraphiQL playground
	r.GET("/ping", pingHandler)                       // Ping/pong endpoint (for healthcheck)

	r.Run() // Run on 0.0.0.0:8080
}
