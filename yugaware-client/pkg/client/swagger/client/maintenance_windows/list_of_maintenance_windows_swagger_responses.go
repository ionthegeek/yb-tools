// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// ListOfMaintenanceWindowsReader is a Reader for the ListOfMaintenanceWindows structure.
type ListOfMaintenanceWindowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOfMaintenanceWindowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOfMaintenanceWindowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListOfMaintenanceWindowsOK creates a ListOfMaintenanceWindowsOK with default headers values
func NewListOfMaintenanceWindowsOK() *ListOfMaintenanceWindowsOK {
	return &ListOfMaintenanceWindowsOK{}
}

/* ListOfMaintenanceWindowsOK describes a response with status code 200, with default header values.

successful operation
*/
type ListOfMaintenanceWindowsOK struct {
	Payload []*models.MaintenanceWindow
}

func (o *ListOfMaintenanceWindowsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/customers/{cUUID}/maintenance_windows/list][%d] listOfMaintenanceWindowsOK  %+v", 200, o.Payload)
}
func (o *ListOfMaintenanceWindowsOK) GetPayload() []*models.MaintenanceWindow {
	return o.Payload
}

func (o *ListOfMaintenanceWindowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}