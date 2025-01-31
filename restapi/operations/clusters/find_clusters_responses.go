package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/radanalyticsio/oshinko-rest/models"
)

/*FindClustersOK Clusters response

swagger:response findClustersOK
*/
type FindClustersOK struct {

	// In: body
	Payload FindClustersOKBodyBody `json:"body,omitempty"`
}

// NewFindClustersOK creates FindClustersOK with default headers values
func NewFindClustersOK() *FindClustersOK {
	return &FindClustersOK{}
}

// WithPayload adds the payload to the find clusters o k response
func (o *FindClustersOK) WithPayload(payload FindClustersOKBodyBody) *FindClustersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find clusters o k response
func (o *FindClustersOK) SetPayload(payload FindClustersOKBodyBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindClustersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindClustersDefault Unexpected error

swagger:response findClustersDefault
*/
type FindClustersDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFindClustersDefault creates FindClustersDefault with default headers values
func NewFindClustersDefault(code int) *FindClustersDefault {
	if code <= 0 {
		code = 500
	}

	return &FindClustersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find clusters default response
func (o *FindClustersDefault) WithStatusCode(code int) *FindClustersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find clusters default response
func (o *FindClustersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find clusters default response
func (o *FindClustersDefault) WithPayload(payload *models.ErrorResponse) *FindClustersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find clusters default response
func (o *FindClustersDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindClustersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
