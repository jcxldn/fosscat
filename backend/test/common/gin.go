package common

import (
	"net/http"
	"net/http/httptest"
	"os"

	jwt "github.com/jcxldn/fosscat/backend/util/jwt"

	"github.com/gin-gonic/gin"
)

type GinTestSuite struct {
	UserDatabaseTestSuite
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
}

func (s *GinTestSuite) SetupSuite() {
	// Run inherited SetupSuite
	s.UserDatabaseTestSuite.SetupSuite()

	// Setup JWT
	os.Setenv("JWT_PK_FILE", os.TempDir()+"/fosscat-test.key")
	jwt.SetupKey()

	// Set gin to test mode (supresses some output)
	gin.SetMode(gin.TestMode)
}

func (s *GinTestSuite) SetupTest() {
	s.Recorder = httptest.NewRecorder()
	s.Context, _ = gin.CreateTestContext(s.Recorder)

	// Populate s.Context.Request
	// avoids nil pointer dereference
	// eg when calling c.Request.Header.Get(...) in auth_middleware.go
	req, _ := http.NewRequest(http.MethodGet, "/invalid", nil)
	s.Context.Request = req

}

func (s *GinTestSuite) TeardownTest() {
	s.Recorder = nil
	s.Context = nil
}
