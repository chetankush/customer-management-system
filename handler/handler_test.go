package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"car-service/datastore"
	"car-service/model"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
)

func initializeHandlerTest(t *testing.T) (*datastore.MockCustomer, Handler, *gofr.Gofr) {
	ctrl := gomock.NewController(t)

	mockStore := datastore.NewMockCustomer(ctrl)
	h := NewHandler(mockStore)
	app := gofr.New()

	return mockStore, h, app
}

func TestCreate(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	input := `{
		"id": 1,
		"name": "Chetan Kushwah",
		"email": "chetankushwah929@gmail.com",
		"phone": "1234567890",
		"address": "123 Main St",
		"city": "Anytown",
		"date_of_birth": "1990-01-01",
		"is_active": true
	}`

	expResp := &model.Customer{
		ID:          1,
		Name:        "Chetan Kushwah",
		Email:       "chetankushwah929@gmail.com",
		Phone:       "1234567890",
		Address:     "123 Main St",
		City:        "Anytown",
		DateOfBirth: "1990-01-01",
		IsActive:    true,
	}

	in := strings.NewReader(input)
	req := httptest.NewRequest(http.MethodPost, "/customer", in)
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	var customer model.Customer

	_ = ctx.Bind(&customer)

	mockStore.EXPECT().Create(ctx, &customer).Return(expResp, nil).MaxTimes(1)

	resp, err := h.Create(ctx)

	assert.Nil(t, err, "TestCreate: failed - success case")
	assert.Equal(t, expResp, resp, "TestCreate: failed - success case")
}

func TestCreate_Error(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	tests := []struct {
		desc    string
		input   string
		expResp interface{}
		err     error
	}{
		{
			desc: "create with invalid body",
			input: `{
				"id": 1,
				"name": "Chetan Kushwah",
				"email": "chetankushwah929@gmail.com",
				"phone": "1234567890",
				"address": "123 Main St",
				"city": "Anytown",
				"date_of_birth": "1990-01-01",
				"is_active": true
			}`,
			expResp: &model.Customer{},
			err:     errors.InvalidParam{Param: []string{"body"}},
		},
		{
			desc:    "create with invalid body",
			input:   `{}`,
			expResp: &model.Customer{},
			err:     errors.InvalidParam{Param: []string{"body"}},
		},
	
	}

	for i, tc := range tests {
		in := strings.NewReader(tc.input)
		req := httptest.NewRequest(http.MethodPost, "/customer", in)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		var customer model.Customer

		_ = ctx.Bind(&customer)

		mockStore.EXPECT().Create(ctx, &customer).Return(tc.expResp.(*model.Customer), tc.err).MaxTimes(1)

		resp, err := h.Create(ctx)

		assert.Equal(t, tc.err, err, "TestCreate_Error[%d]: failed.\n%s", i, tc.desc)
		assert.Nil(t, resp, "TestCreate_Error[%d]: failed.\n%s", i, tc.desc)
	}
}

func TestGetAll(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	mockData := []*model.Customer{
		{
			ID:          1,
			Name:        "Chetan Kushwah",
			Email:       "chetankushwah929@gmail.com",
			Phone:       "1234567890",
			Address:     "123 Main St",
			City:        "Anytown",
			DateOfBirth: "1990-01-01",
			IsActive:    true,
		},
	}

	mockStore.EXPECT().GetAll(gomock.Any()).Return(mockData, nil).MaxTimes(1)

	resp, err := h.GetAll(gofr.NewContext(nil, nil, app))

	assert.Nil(t, err, "TestGetAll: failed - success case")

	assert.Equal(t, mockData, resp, "TestGetAll: failed - success case")
}

func TestGetByID(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	mockData := &model.Customer{
		ID:          1,
		Name:        "Chetan Kushwah",
		Email:       "chetankushwah929@gmail.com",
		Phone:       "1234567890",
		Address:     "123 Main St",
		City:        "Anytown",
		DateOfBirth: "1890-01-01",
		IsActive:    true,
	}

	mockStore.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(mockData, nil).MaxTimes(1)

	resp, err := h.GetByID(gofr.NewContext(nil, nil, app))

	assert.Nil(t, err, "TestGetByID: failed - success case")

	assert.Equal(t, mockData, resp, "TestGetByID: failed - success case")
}

func TestUpdate(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	mockData := &model.Customer{
		ID:          1,
		Name:        "Updated Name",
		Email:       "chetankushwah999@gmail.com",
		Phone:       "1234567890",
		Address:     "123 Main St",
		City:        "Anytown",
		DateOfBirth: "1900-01-01",
		IsActive:    true,
	}

	mockStore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mockData, nil).MaxTimes(1)

	resp, err := h.Update(gofr.NewContext(nil, nil, app))

	assert.Nil(t, err, "TestUpdate: failed - success case")

	assert.Equal(t, mockData, resp, "TestUpdate: failed - success case")
}

func TestDelete(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	mockData := "Data Deleted successfully"

	mockStore.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).MaxTimes(1)

    	resp, err := h.Delete(gofr.NewContext(nil, nil, app))

	assert.Nil(t, err, "TestDelete: failed - success case")


	assert.Equal(t, mockData, resp, "TestDelete: failed - success case")
}
