package user

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"test/gomock/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)

	var id int64 = 1
	mockMale := mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(errors.New("test error")),
	)
	user := NewUser(mockMale)
	err := user.GetUserInfo(id)
	assert.Equal(t, errors.New("test error"), err)
}
