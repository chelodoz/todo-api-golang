package note

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateNoteHandlerTestSuite struct {
	suite.Suite
	router          *mux.Router
	noteHandlerTest Handler
	mockService     *MockService
}

func TestCreateNoteHandlerTestSuite(t *testing.T) {
	suite.Run(t, &CreateNoteHandlerTestSuite{})
}

func (suite *CreateNoteHandlerTestSuite) SetupSuite() {
	suite.mockService = NewMockService(suite.T())
	suite.noteHandlerTest = NewHandler(suite.mockService)
	suite.router = mux.NewRouter()
	suite.router.HandleFunc("/notes", suite.noteHandlerTest.Create).Methods(http.MethodPost)
}

func (suite *CreateNoteHandlerTestSuite) TestCreateNote_ValidTestInput_ShouldReturn201Created() {
	// Arrange
	var request = strings.NewReader(`{
		"name": "test",
		"description": "description"
	}`)

	suite.mockService.On("Create", mock.Anything, mock.Anything).Return(&Note{
		ID:          uuid.New(),
		Name:        "test",
		Description: "description",
		CreatedAt:   time.Now().UTC(),
	}, nil)

	// Act
	req, _ := http.NewRequest(http.MethodPost, "/notes", request)
	res := suite.ExecuteRequest(req, suite.router)

	// Assert

	suite.Equal(http.StatusCreated, res.StatusCode)
}

func (suite *CreateNoteHandlerTestSuite) TestCreateNote_InvalidInputWithMissingDescription_ShouldReturn400BadRequest() {
	// Arrange
	var request = strings.NewReader(`{
		"name": "test"
	}`)

	// Act
	req, _ := http.NewRequest(http.MethodPost, "/notes", request)
	res := suite.ExecuteRequest(req, suite.router)

	// Assert

	suite.Equal(http.StatusBadRequest, res.StatusCode)
}

func (suite *CreateNoteHandlerTestSuite) TestCreateNote_InvalidInputWithWrongJsonFormat_ShouldReturn422UnprocessableEntity() {
	// Arrange
	var request = strings.NewReader(`{
		"name": ["test"]
	}`)

	// Act
	req, _ := http.NewRequest(http.MethodPost, "/notes", request)
	res := suite.ExecuteRequest(req, suite.router)

	// Assert

	suite.Equal(http.StatusUnprocessableEntity, res.StatusCode)
}

func (suite *CreateNoteHandlerTestSuite) ExecuteRequest(req *http.Request, router *mux.Router) *http.Response {
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	res := recorder.Result()
	return res
}
