// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUsersUserIDHandlerFunc turns a function with the right signature into a get users user ID handler
type GetUsersUserIDHandlerFunc func(GetUsersUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersUserIDHandlerFunc) Handle(params GetUsersUserIDParams) middleware.Responder {
	return fn(params)
}

// GetUsersUserIDHandler interface for that can handle valid get users user ID params
type GetUsersUserIDHandler interface {
	Handle(GetUsersUserIDParams) middleware.Responder
}

// NewGetUsersUserID creates a new http.Handler for the get users user ID operation
func NewGetUsersUserID(ctx *middleware.Context, handler GetUsersUserIDHandler) *GetUsersUserID {
	return &GetUsersUserID{Context: ctx, Handler: handler}
}

/*GetUsersUserID swagger:route GET /users/{userId} getUsersUserId

Returns a user by ID.

*/
type GetUsersUserID struct {
	Context *middleware.Context
	Handler GetUsersUserIDHandler
}

func (o *GetUsersUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersUserIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}