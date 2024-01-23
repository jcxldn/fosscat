package common

import (
	"context"
	"fmt"
	"log"

	"github.com/jcxldn/fosscat/backend/database"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define the test suite, absorbing the testify basic suite
type DatabaseTestSuite struct {
	suite.Suite
	DB          *gorm.DB
	dbContainer testcontainers.Container
	dbCtx       context.Context
}

// SetupSuite is called before any tests run
func (s *DatabaseTestSuite) SetupSuite() {
	s.dbCtx = context.Background()
	// Create a container request for a container running the "postgres" image
	req := testcontainers.ContainerRequest{
		Image:        "postgres",                                                                      // container image to use
		ExposedPorts: []string{"5432"},                                                                // ports to expose to host
		WaitingFor:   wait.ForLog("database system is ready to accept connections").WithOccurrence(2), // trigger to define when container has started up
		Env: map[string]string{ // environment variables
			"POSTGRES_DB":       "fosscat",
			"POSTGRES_USER":     "fosscat",
			"POSTGRES_PASSWORD": "fosscat",
		},
	}

	// Note that we do not define any persistent storage so the database will start from scratch every time it is created.

	// Create the container **and** wait for it to start up
	dbContainer, err := testcontainers.GenericContainer(s.dbCtx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,  // specify the container request
		Started:          true, // automatically start once created
	})

	s.dbContainer = dbContainer

	if err != nil {
		panic(err)
	}

	// Determine the IP address of the container
	ctrIp, _ := dbContainer.ContainerIP(s.dbCtx)

	// Log that the database is now available
	log.Default().Printf("[test/common/database]: ephemeral db available on %s:5432", ctrIp)

	// Define connection details for the database
	dsn := fmt.Sprintf(
		"host=%s user=fosscat password=fosscat dbname=fosscat port=5432", ctrIp,
	)

	// Attempt to connect to the db
	db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err2 != nil {
		panic(err)
	}

	// Call migrate function (defined in backend/database/database.go) to create tables
	database.Migrate(db)

	// make the GORM instance available to tests
	s.DB = db

}

// TeardownSuite is called when all tests have finished
func (s *DatabaseTestSuite) TeardownSuite() {
	log.Default().Printf("[test/common/database]: teardown in progress")
	// Attempt to stop (terminate) the container, panicking if any errors are encountered.
	if err := s.dbContainer.Terminate(s.dbCtx); err != nil {
		panic(err)
	}
}
