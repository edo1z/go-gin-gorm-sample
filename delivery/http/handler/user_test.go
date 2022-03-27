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
	mockUserList := []*domain.User{&mockUser}
	u := new(usecase.MockUserUsecase)
	u.On("GetAll").Return(mockUserList, nil)
	h := NewUserHandler(u)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/users", h.GetAll)
	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUserGetById(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	u := new(usecase.MockUserUsecase)
	u.On("GetById").Return(&mockUser, nil)
	h := NewUserHandler(u)

	cases := map[string]int{"1": 200, "a": 400}
	for id, code := range cases {
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)
		r.GET("/users/view/:id", h.GetById)
		req, err := http.NewRequest("GET", "/users/view/"+id, nil)
		assert.NoError(t, err)
		r.ServeHTTP(w, req)
		assert.Equal(t, code, w.Code)
	}
}
