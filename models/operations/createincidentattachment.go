// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"firehydrant/internal/utils"
	"firehydrant/models/components"
	"fmt"
	"time"
)

type File struct {
	FileName string `multipartForm:"name=fileName"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *File) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *File) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type VoteDirection string

const (
	VoteDirectionUp   VoteDirection = "up"
	VoteDirectionDown VoteDirection = "down"
)

func (e VoteDirection) ToPointer() *VoteDirection {
	return &e
}
func (e *VoteDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "up":
		fallthrough
	case "down":
		*e = VoteDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VoteDirection: %v", v)
	}
}

type CreateIncidentAttachmentRequestBody struct {
	File          File           `multipartForm:"file,name=file"`
	Description   *string        `multipartForm:"name=description"`
	OccurredAt    *time.Time     `multipartForm:"name=occurred_at"`
	VoteDirection *VoteDirection `multipartForm:"name=vote_direction"`
}

func (c CreateIncidentAttachmentRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateIncidentAttachmentRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateIncidentAttachmentRequestBody) GetFile() File {
	if o == nil {
		return File{}
	}
	return o.File
}

func (o *CreateIncidentAttachmentRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateIncidentAttachmentRequestBody) GetOccurredAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.OccurredAt
}

func (o *CreateIncidentAttachmentRequestBody) GetVoteDirection() *VoteDirection {
	if o == nil {
		return nil
	}
	return o.VoteDirection
}

type CreateIncidentAttachmentRequest struct {
	IncidentID  string                              `pathParam:"style=simple,explode=false,name=incident_id"`
	RequestBody CreateIncidentAttachmentRequestBody `request:"mediaType=multipart/form-data"`
}

func (o *CreateIncidentAttachmentRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentAttachmentRequest) GetRequestBody() CreateIncidentAttachmentRequestBody {
	if o == nil {
		return CreateIncidentAttachmentRequestBody{}
	}
	return o.RequestBody
}

type CreateIncidentAttachmentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Allows adding image attachments to an incident
	IncidentAttachmentEntity *components.IncidentAttachmentEntity
}

func (o *CreateIncidentAttachmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateIncidentAttachmentResponse) GetIncidentAttachmentEntity() *components.IncidentAttachmentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentAttachmentEntity
}
