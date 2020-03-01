package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/GIT_USER_ID/GIT_REPO_ID/pkg/openapi"
	"github.com/gorilla/mux"
)

type MockApiController struct {
	service openapi.DefaultApiServicer
}

func NewMockApiController(s openapi.DefaultApiServicer) openapi.Router {
	return &MockApiController{ service: s }
}

func (c *MockApiController) Routes() openapi.Routes {
	return openapi.Routes{
		{
			"CreateUser",
			strings.ToUpper("Post"),
			"/users",
			c.CreateUser,
		},
		{
			"GetUser",
			strings.ToUpper("Get"),
			"/users/{id}",
			c.GetUser,
		},
		{
			"UpdateUser",
			strings.ToUpper("Put"),
			"/users",
			c.UpdateUser,
		},
	}
}

func (c *MockApiController) CreateUser(w http.ResponseWriter, r *http.Request) {
	createUserInput := &openapi.CreateUserInput{}
	if err := json.NewDecoder(r.Body).Decode(&createUserInput); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.CreateUser(*createUserInput)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	openapi.EncodeJSONResponse(result, nil, w)
}

// GetUser -
func (c *MockApiController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.GetUser(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	openapi.EncodeJSONResponse(result, nil, w)
}

// UpdateUser -
func (c *MockApiController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateUserInput := &openapi.UpdateUserInput{}
	if err := json.NewDecoder(r.Body).Decode(&updateUserInput); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.UpdateUser(*updateUserInput)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	openapi.EncodeJSONResponse(result, nil, w)
}
