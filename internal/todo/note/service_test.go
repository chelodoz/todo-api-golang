package note

import (
	"testing"
	"time"

	"github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateNoteServiceTestSuite struct {
	suite.Suite
	noteService        Service
	mockNoteRepository *MockRepository
}

func TestCreateNoteTestSuite(t *testing.T) {
	suite.Run(t, &CreateNoteServiceTestSuite{})
}

func (suite *CreateNoteServiceTestSuite) SetupSuite() {
	suite.mockNoteRepository = NewMockRepository(suite.T())
	suite.noteService = NewService(suite.mockNoteRepository)
}
func (suite *CreateNoteServiceTestSuite) TestCreateNote_ValidTestInput_ShouldReturnCreatedNoteWithoutError() {
	note := &Note{
		ID:          uuid.New(),
		Name:        "note",
		Description: "note description",
		CreatedAt:   time.Now().UTC(),
	}

	suite.mockNoteRepository.On("Create", mock.Anything, mock.Anything).Return(note, nil)

	newNote, err := suite.noteService.Create(note, nil)

	suite.Nil(err)
	suite.NotNil(newNote)
}
