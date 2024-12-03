package repositories

// import (
// 	"errors"
// 	"tech-challenge-fase-1/internal/core/entities"
// 	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
// 	"tech-challenge-fase-1/internal/tests/mocks"
// 	"testing"
//
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )
//
// func TestInsertOrder(t *testing.T) {
//
// 	conn := mocks.NewMockConnectionDB(t)
//
// 	repo := NewOrderRepositoryDB(conn)
//
// 	order := entities.RestoreOrder(uuid.NewString(), nil, []*valueobjects.OrderItem{},
// 		entities.ORDER_PAYMENT_PAID, entities.ORDER_PREPARATION_FINISHED,
// 	)
//
// 	conn.On("Exec", mock.Anything, order.GetId(),
// 		order.GetCustomerId(),
// 		newOrderItemHelperList(order.GetItems()),
// 		order.GetPaymentStatus().String(),
// 		order.GetPreparationStatus().String(),
// 	).Return(nil).Once()
//
// 	err := repo.Insert(order)
//
// 	assert.Nil(t, err)
// }
//
// func TestUpdateOrder(t *testing.T) {
//
// 	conn := mocks.NewMockConnectionDB(t)
//
// 	repo := NewOrderRepositoryDB(conn)
//
// 	order := entities.RestoreOrder(uuid.NewString(), nil, []*valueobjects.OrderItem{},
// 		entities.ORDER_PAYMENT_PAID, entities.ORDER_PREPARATION_FINISHED,
// 	)
//
// 	conn.On("Exec", mock.Anything,
// 		order.GetCustomerId(),
// 		newOrderItemHelperList(order.GetItems()),
// 		order.GetPaymentStatus().String(),
// 		order.GetPreparationStatus().String(),
// 		order.GetId(),
// 	).Return(nil).Once()
//
// 	err := repo.Update(order)
//
// 	assert.Nil(t, err)
// }
//
// func TestFindOrderByID(t *testing.T) {
//
// 	conn := mocks.NewMockConnectionDB(t)
//
// 	row := mocks.NewMockRowDB(t)
//
// 	repo := NewOrderRepositoryDB(conn)
//
// 	order := entities.RestoreOrder(uuid.NewString(), nil, []*valueobjects.OrderItem{},
// 		entities.ORDER_PAYMENT_PAID, entities.ORDER_PREPARATION_FINISHED,
// 	)
//
// 	conn.On("QueryRow", mock.Anything, order.GetId()).Return(row, nil).Once()
// 	row.On("Scan", mock.Anything, mock.Anything, mock.Anything,
// 		mock.Anything, mock.Anything).Return(nil).Once()
//
// 	_, err := repo.FindOrderByID(order.GetId())
//
// 	assert.Nil(t, err)
// }
//
// func TestFindOrderByIDWithErrOrderNotFound(t *testing.T) {
//
// 	conn := mocks.NewMockConnectionDB(t)
//
// 	row := mocks.NewMockRowDB(t)
//
// 	repo := NewOrderRepositoryDB(conn)
//
// 	order := entities.RestoreOrder(uuid.NewString(), nil, []*valueobjects.OrderItem{},
// 		entities.ORDER_PAYMENT_PAID, entities.ORDER_PREPARATION_FINISHED,
// 	)
//
// 	conn.On("QueryRow", mock.Anything, order.GetId()).Return(row, nil).Once()
// 	row.On("Scan", mock.Anything, mock.Anything, mock.Anything,
// 		mock.Anything, mock.Anything).Return(errors.New("no rows in result set")).Once()
//
// 	_, err := repo.FindOrderByID(order.GetId())
//
// 	assert.EqualError(t, err, ErrOrderNotFound.Error())
// }
//
// func TestFindOrderByIDWithError(t *testing.T) {
//
// 	conn := mocks.NewMockConnectionDB(t)
//
// 	row := mocks.NewMockRowDB(t)
//
// 	repo := NewOrderRepositoryDB(conn)
//
// 	order := entities.RestoreOrder(uuid.NewString(), nil, []*valueobjects.OrderItem{},
// 		entities.ORDER_PAYMENT_PAID, entities.ORDER_PREPARATION_FINISHED,
// 	)
//
// 	conn.On("QueryRow", mock.Anything, order.GetId()).Return(row, nil).Once()
// 	row.On("Scan", mock.Anything, mock.Anything, mock.Anything,
// 		mock.Anything, mock.Anything).Return(errors.New("error")).Once()
//
// 	_, err := repo.FindOrderByID(order.GetId())
//
// 	assert.EqualError(t, err, "error")
// }
