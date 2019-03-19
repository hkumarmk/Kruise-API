// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateDeploymentHandlerFunc turns a function with the right signature into a create deployment handler
type CreateDeploymentHandlerFunc func(CreateDeploymentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateDeploymentHandlerFunc) Handle(params CreateDeploymentParams) middleware.Responder {
	return fn(params)
}

// CreateDeploymentHandler interface for that can handle valid create deployment params
type CreateDeploymentHandler interface {
	Handle(CreateDeploymentParams) middleware.Responder
}

// NewCreateDeployment creates a new http.Handler for the create deployment operation
func NewCreateDeployment(ctx *middleware.Context, handler CreateDeploymentHandler) *CreateDeployment {
	return &CreateDeployment{Context: ctx, Handler: handler}
}

/*CreateDeployment swagger:route POST /deployment deployments createDeployment

Generates a new Kruise deployment

*/
type CreateDeployment struct {
	Context *middleware.Context
	Handler CreateDeploymentHandler
}

func (o *CreateDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateDeploymentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
