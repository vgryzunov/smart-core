// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"smart-core/swagger/models"
)

// GetHostnameOKCode is the HTTP code returned for type GetHostnameOK
const GetHostnameOKCode int = 200

/*GetHostnameOK returns the hostname of the machine

swagger:response getHostnameOK
*/
type GetHostnameOK struct {

	/*the hostname of the machine
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetHostnameOK creates GetHostnameOK with default headers values
func NewGetHostnameOK() *GetHostnameOK {

	return &GetHostnameOK{}
}

// WithPayload adds the payload to the get hostname o k response
func (o *GetHostnameOK) WithPayload(payload string) *GetHostnameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hostname o k response
func (o *GetHostnameOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHostnameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetHostnameDefault error

swagger:response getHostnameDefault
*/
type GetHostnameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHostnameDefault creates GetHostnameDefault with default headers values
func NewGetHostnameDefault(code int) *GetHostnameDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHostnameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get hostname default response
func (o *GetHostnameDefault) WithStatusCode(code int) *GetHostnameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get hostname default response
func (o *GetHostnameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get hostname default response
func (o *GetHostnameDefault) WithPayload(payload *models.Error) *GetHostnameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hostname default response
func (o *GetHostnameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHostnameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
