package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TopicInfoRecordV2 topic info record v2
// swagger:model TopicInfoRecordV2
type TopicInfoRecordV2 struct {

	// A unique identifier for each topic.
	ID string `json:"id,omitempty"`

	// A summarizing word or phrase for the topic.
	KeyPhrase string `json:"keyPhrase,omitempty"`

	// Count of documents assigned to topic.
	Score float64 `json:"score,omitempty"`
}

// Validate validates this topic info record v2
func (m *TopicInfoRecordV2) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}