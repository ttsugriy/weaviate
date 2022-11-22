//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// BackupsRestoreReader is a Reader for the BackupsRestore structure.
type BackupsRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupsRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupsRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBackupsRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupsRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupsRestoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewBackupsRestoreUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupsRestoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBackupsRestoreOK creates a BackupsRestoreOK with default headers values
func NewBackupsRestoreOK() *BackupsRestoreOK {
	return &BackupsRestoreOK{}
}

/*
BackupsRestoreOK handles this case with default header values.

Backup restoration process successfully started.
*/
type BackupsRestoreOK struct {
	Payload *models.BackupRestoreResponse
}

func (o *BackupsRestoreOK) Error() string {
	return fmt.Sprintf("[POST /backups/{backend}/{id}/restore][%d] backupsRestoreOK  %+v", 200, o.Payload)
}

func (o *BackupsRestoreOK) GetPayload() *models.BackupRestoreResponse {
	return o.Payload
}

func (o *BackupsRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.BackupRestoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsRestoreUnauthorized creates a BackupsRestoreUnauthorized with default headers values
func NewBackupsRestoreUnauthorized() *BackupsRestoreUnauthorized {
	return &BackupsRestoreUnauthorized{}
}

/*
BackupsRestoreUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type BackupsRestoreUnauthorized struct{}

func (o *BackupsRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[POST /backups/{backend}/{id}/restore][%d] backupsRestoreUnauthorized ", 401)
}

func (o *BackupsRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewBackupsRestoreForbidden creates a BackupsRestoreForbidden with default headers values
func NewBackupsRestoreForbidden() *BackupsRestoreForbidden {
	return &BackupsRestoreForbidden{}
}

/*
BackupsRestoreForbidden handles this case with default header values.

Forbidden
*/
type BackupsRestoreForbidden struct {
	Payload *models.ErrorResponse
}

func (o *BackupsRestoreForbidden) Error() string {
	return fmt.Sprintf("[POST /backups/{backend}/{id}/restore][%d] backupsRestoreForbidden  %+v", 403, o.Payload)
}

func (o *BackupsRestoreForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsRestoreNotFound creates a BackupsRestoreNotFound with default headers values
func NewBackupsRestoreNotFound() *BackupsRestoreNotFound {
	return &BackupsRestoreNotFound{}
}

/*
BackupsRestoreNotFound handles this case with default header values.

Not Found - Backup does not exist
*/
type BackupsRestoreNotFound struct {
	Payload *models.ErrorResponse
}

func (o *BackupsRestoreNotFound) Error() string {
	return fmt.Sprintf("[POST /backups/{backend}/{id}/restore][%d] backupsRestoreNotFound  %+v", 404, o.Payload)
}

func (o *BackupsRestoreNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsRestoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsRestoreUnprocessableEntity creates a BackupsRestoreUnprocessableEntity with default headers values
func NewBackupsRestoreUnprocessableEntity() *BackupsRestoreUnprocessableEntity {
	return &BackupsRestoreUnprocessableEntity{}
}

/*
BackupsRestoreUnprocessableEntity handles this case with default header values.

Invalid backup restoration attempt.
*/
type BackupsRestoreUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *BackupsRestoreUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /backups/{backend}/{id}/restore][%d] backupsRestoreUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BackupsRestoreUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsRestoreUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsRestoreInternalServerError creates a BackupsRestoreInternalServerError with default headers values
func NewBackupsRestoreInternalServerError() *BackupsRestoreInternalServerError {
	return &BackupsRestoreInternalServerError{}
}

/*
BackupsRestoreInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type BackupsRestoreInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *BackupsRestoreInternalServerError) Error() string {
	return fmt.Sprintf("[POST /backups/{backend}/{id}/restore][%d] backupsRestoreInternalServerError  %+v", 500, o.Payload)
}

func (o *BackupsRestoreInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsRestoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
