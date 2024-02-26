package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/graph/resolver"
	"github.com/jcxldn/fosscat/backend/structs"
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

// Define the upload route handler
func uploadHandler(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		user := c.Request.Context().Value("user")
		if user != nil {
			// User context available.
			// Type assertion that "user" ctx is not nil and is of type structs.User
			userStruct := user.(*structs.User)

			// Check if upload[] was set
			fileExists, err := c.FormFile("upload[]")
			if fileExists == nil || err != nil {
				c.JSON(422, gin.H{
					"error": "upload[] not set",
				})
				return
			}

			// upload[] was set

			// Multipart form
			form, _ := c.MultipartForm()
			files := form.File["upload[]"]

			uploadedFiles, err := database.UploadFiles(db, files, *userStruct)

			var returnFiles []*structs.UploadedFile

			// Get UUIDs for each uploaded file and return those to keep things simple for the client
			// They don't need []*structs.File, they can request over GraphQL if they need it.
			for _, fileStruct := range uploadedFiles {
				uploadedFile := structs.UploadedFile{
					ID:   fileStruct.ID,
					Name: fileStruct.Name,
				}
				returnFiles = append(returnFiles, &uploadedFile)
			}

			if err == nil {
				c.JSON(200, gin.H{
					"files": returnFiles,
				})
			} else {
				c.JSON(400, gin.H{
					"error": err,
				})
			}
		} else {
			// User context not available
			c.JSON(400, gin.H{
				"error": "route requires authorization",
			})
		}
	}
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
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	r.Use(AuthMiddleware(resolver.DB))

	// Define routes
	r.POST("/graphql", graphqlHandler(resolver))      // GraphQL query endpoint
	r.GET("/graphql/playground", playgroundHandler()) // GraphiQL playground
	r.GET("/ping", pingHandler)                       // Ping/pong endpoint (for healthcheck)
	// Gin multiple file upload: based on <https://gin-gonic.com/docs/examples/upload-file/multiple-file/>
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", uploadHandler(resolver.DB))

	r.Run() // Run on 0.0.0.0:8080
}
