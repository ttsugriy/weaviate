//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IndexStatusList The status of all the indexes of a Class
//
// swagger:model IndexStatusList
type IndexStatusList struct {

	// Name of the class
	ClassName string `json:"className,omitempty"`

	// A list of indexes for this class
	Indexes []*IndexStatus `json:"indexes"`

	// Number of shards with indexes for this class
	ShardCount int64 `json:"shardCount,omitempty"`

	// Total number of indexes for this class
	Total int64 `json:"total,omitempty"`
}

// Validate validates this index status list
func (m *IndexStatusList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndexes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IndexStatusList) validateIndexes(formats strfmt.Registry) error {

	if swag.IsZero(m.Indexes) { // not required
		return nil
	}

	for i := 0; i < len(m.Indexes); i++ {
		if swag.IsZero(m.Indexes[i]) { // not required
			continue
		}

		if m.Indexes[i] != nil {
			if err := m.Indexes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("indexes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IndexStatusList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexStatusList) UnmarshalBinary(b []byte) error {
	var res IndexStatusList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
