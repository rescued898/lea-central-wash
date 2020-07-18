// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SaveCollectionReader is a Reader for the SaveCollection structure.
type SaveCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewSaveCollectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSaveCollectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSaveCollectionNoContent creates a SaveCollectionNoContent with default headers values
func NewSaveCollectionNoContent() *SaveCollectionNoContent {
	return &SaveCollectionNoContent{}
}

/*SaveCollectionNoContent handles this case with default header values.

OK
*/
type SaveCollectionNoContent struct {
}

func (o *SaveCollectionNoContent) Error() string {
	return fmt.Sprintf("[POST /save-collection][%d] saveCollectionNoContent ", 204)
}

func (o *SaveCollectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveCollectionInternalServerError creates a SaveCollectionInternalServerError with default headers values
func NewSaveCollectionInternalServerError() *SaveCollectionInternalServerError {
	return &SaveCollectionInternalServerError{}
}

/*SaveCollectionInternalServerError handles this case with default header values.

internal error
*/
type SaveCollectionInternalServerError struct {
}

func (o *SaveCollectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /save-collection][%d] saveCollectionInternalServerError ", 500)
}

func (o *SaveCollectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
