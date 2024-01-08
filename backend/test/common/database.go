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

// Mock GORM
// based on https://github.com/go-gorm/gorm/issues/1525#issuecomment-376164189
// and (updated ver): https://addshore.com/2023/09/a-copy-paste-go-sql-mock-for-gorm/
func (s *DatabaseTestSuite) SetupSuite() {
	s.dbCtx = context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{"5432"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
		Env: map[string]string{
			"POSTGRES_DB":       "fosscat",
			"POSTGRES_USER":     "fosscat",
			"POSTGRES_PASSWORD": "fosscat",
		},
	}

	dbContainer, err := testcontainers.GenericContainer(s.dbCtx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	s.dbContainer = dbContainer

	if err != nil {
		panic(err)
	}

	ctrIp, _ := dbContainer.ContainerIP(s.dbCtx)
	log.Default().Printf("[test/common/database]: ephemeral db available on %s:5432", ctrIp)
	dsn := fmt.Sprintf(
		"host=%s user=fosscat password=fosscat dbname=fosscat port=5432", ctrIp,
	)

	// Attempt to connect to the db
	db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err2 != nil {
		panic(err)
	}

	database.Migrate(db)

	s.DB = db

}

func (s *DatabaseTestSuite) TeardownSuite() {
	log.Default().Printf("[test/common/database]: teardown in progress")
	if err := s.dbContainer.Terminate(s.dbCtx); err != nil {
		panic(err)
	}
}
