// Code generated by go-swagger; DO NOT EDIT.

package update

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateDNSEntryHandlerFunc turns a function with the right signature into a update dns entry handler
type UpdateDNSEntryHandlerFunc func(UpdateDNSEntryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateDNSEntryHandlerFunc) Handle(params UpdateDNSEntryParams) middleware.Responder {
	return fn(params)
}

// UpdateDNSEntryHandler interface for that can handle valid update dns entry params
type UpdateDNSEntryHandler interface {
	Handle(UpdateDNSEntryParams) middleware.Responder
}

// NewUpdateDNSEntry creates a new http.Handler for the update dns entry operation
func NewUpdateDNSEntry(ctx *middleware.Context, handler UpdateDNSEntryHandler) *UpdateDNSEntry {
	return &UpdateDNSEntry{Context: ctx, Handler: handler}
}

/*
	UpdateDNSEntry swagger:route PUT /dns update updateDnsEntry

Update dns entry
*/
type UpdateDNSEntry struct {
	Context *middleware.Context
	Handler UpdateDNSEntryHandler
}

func (o *UpdateDNSEntry) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateDNSEntryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
