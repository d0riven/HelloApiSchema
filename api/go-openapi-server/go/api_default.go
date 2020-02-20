/*
 * HelloApiSchema
 *
 * Practice api schema
 *
 * API version: 1.0.0
 * Contact: doriven@example.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"AddUser",
			strings.ToUpper("Post"),
			"/users",
			c.AddUser,
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

// AddUser - 
func (c *DefaultApiController) AddUser(w http.ResponseWriter, r *http.Request) { 
	createUserInput := &CreateUserInput{}
	if err := json.NewDecoder(r.Body).Decode(&createUserInput); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.AddUser(*createUserInput)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// GetUser - 
func (c *DefaultApiController) GetUser(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	id, err := parseIntParameter(params["id"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.GetUser(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// UpdateUser - 
func (c *DefaultApiController) UpdateUser(w http.ResponseWriter, r *http.Request) { 
	updateUserInput := &UpdateUserInput{}
	if err := json.NewDecoder(r.Body).Decode(&updateUserInput); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.UpdateUser(*updateUserInput)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}