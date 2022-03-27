package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"web3ten0/go-gin-gorm-sample/domain"
	"web3ten0/go-gin-gorm-sample/domain/mock/usecase"

	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserGetAll(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(rec)
	ginContext.Request = req
	u := new(usecase.MockUserUsecase)
	mockUserList := []*domain.User{&mockUser}
	u.On("GetAll").Return(mockUserList, nil)
	h := NewUserHandler(u)
	h.GetAll(ginContext)
	assert.Equal(t, 200, rec.Code)
}

func TestUserGetById(t *testing.T) {
	if false {
		t.Error("oh my god")
	}
}
