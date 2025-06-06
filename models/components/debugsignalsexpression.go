// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Annotations struct {
}

type Image struct {
	Src *string `json:"src,omitempty"`
	Alt *string `json:"alt,omitempty"`
}

func (o *Image) GetSrc() *string {
	if o == nil {
		return nil
	}
	return o.Src
}

func (o *Image) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

type DebugSignalsExpressionLink struct {
	Href *string `json:"href,omitempty"`
	Text *string `json:"text,omitempty"`
}

func (o *DebugSignalsExpressionLink) GetHref() *string {
	if o == nil {
		return nil
	}
	return o.Href
}

func (o *DebugSignalsExpressionLink) GetText() *string {
	if o == nil {
		return nil
	}
	return o.Text
}

type Signal struct {
	ID             *string                      `json:"id,omitempty"`
	OrganizationID *string                      `json:"organization_id,omitempty"`
	Summary        *string                      `json:"summary,omitempty"`
	Body           *string                      `json:"body,omitempty"`
	Level          *string                      `json:"level,omitempty"`
	Annotations    *Annotations                 `json:"annotations,omitempty"`
	Tags           []string                     `json:"tags,omitempty"`
	Images         []Image                      `json:"images,omitempty"`
	Links          []DebugSignalsExpressionLink `json:"links,omitempty"`
}

func (o *Signal) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Signal) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *Signal) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *Signal) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Signal) GetLevel() *string {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *Signal) GetAnnotations() *Annotations {
	if o == nil {
		return nil
	}
	return o.Annotations
}

func (o *Signal) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Signal) GetImages() []Image {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Signal) GetLinks() []DebugSignalsExpressionLink {
	if o == nil {
		return nil
	}
	return o.Links
}

// DebugSignalsExpression - Debug Signals expressions
type DebugSignalsExpression struct {
	// CEL expression
	Expression string `json:"expression"`
	// List of signals to evaluate rule expression against
	Signals []Signal `json:"signals"`
}

func (o *DebugSignalsExpression) GetExpression() string {
	if o == nil {
		return ""
	}
	return o.Expression
}

func (o *DebugSignalsExpression) GetSignals() []Signal {
	if o == nil {
		return []Signal{}
	}
	return o.Signals
}
