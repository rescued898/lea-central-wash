// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SaveCollectionHandlerFunc turns a function with the right signature into a save collection handler
type SaveCollectionHandlerFunc func(SaveCollectionParams) SaveCollectionResponder

// Handle executing the request and returning a response
func (fn SaveCollectionHandlerFunc) Handle(params SaveCollectionParams) SaveCollectionResponder {
	return fn(params)
}

// SaveCollectionHandler interface for that can handle valid save collection params
type SaveCollectionHandler interface {
	Handle(SaveCollectionParams) SaveCollectionResponder
}

// NewSaveCollection creates a new http.Handler for the save collection operation
func NewSaveCollection(ctx *middleware.Context, handler SaveCollectionHandler) *SaveCollection {
	return &SaveCollection{Context: ctx, Handler: handler}
}

/*SaveCollection swagger:route POST /save-collection saveCollection

SaveCollection save collection API

*/
type SaveCollection struct {
	Context *middleware.Context
	Handler SaveCollectionHandler
}

func (o *SaveCollection) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSaveCollectionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}