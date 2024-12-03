package queries

import (
	"errors"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)
	rows := mocks.NewMockRowsDB(t)

	query := NewOrderDisplayListQueryDB(conn)

	conn.On("Query", mock.Anything).Return(rows, nil).Once()
	rows.On("Next").Return(true).Once()
	rows.On("Next").Return(false).Once()
	rows.On("Scan", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything,
	).Return(nil).Once()

	_, err := query.Execute()

	assert.Nil(t, err)
}

func TestExecuteWithErrorQuery(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)

	query := NewOrderDisplayListQueryDB(conn)

	conn.On("Query", mock.Anything).Return(nil, errors.New("error")).Once()

	_, err := query.Execute()

	assert.Error(t, err)
}
