package handler

/*import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/p12s/wildberries-http-api"
	"github.com/p12s/wildberries-http-api/pkg/service"
	service_mocks "github.com/p12s/wildberries-http-api/pkg/service/mocks"
	"net/http/httptest"
	"testing"
)*/

// TODO дописать тесты, замокать работу с БД
/*
func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockAuthorization, user common.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            common.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "email": "test@test.ru", "name": "Test Name", "password": "qwerty"}`,
			inputUser: common.User{
				Username: "username",
				Email:    "test@test.ru",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user todo.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Empty required fields",
			inputBody:            `{"username": "username"}`,
			inputUser:            todo.User{},
			mockBehavior:         func(r *service_mocks.MockAuthorization, user todo.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error", // поведение handler-а при ошибке сервиса
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: todo.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user todo.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t) // иниц. контроллер и финишируем его
			defer c.Finish()

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
*/