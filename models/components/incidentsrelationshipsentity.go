// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IncidentsRelationshipsEntity - Incidents_RelationshipsEntity model
type IncidentsRelationshipsEntity struct {
	Parent *PublicAPIV1IncidentsSuccinctEntity `json:"parent,omitempty"`
	// The root incident's child incidents.
	Children []PublicAPIV1IncidentsSuccinctEntity `json:"children,omitempty"`
	// A list of incidents that share the root incident's parent.
	Siblings []PublicAPIV1IncidentsSuccinctEntity `json:"siblings,omitempty"`
}

func (o *IncidentsRelationshipsEntity) GetParent() *PublicAPIV1IncidentsSuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Parent
}

func (o *IncidentsRelationshipsEntity) GetChildren() []PublicAPIV1IncidentsSuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Children
}

func (o *IncidentsRelationshipsEntity) GetSiblings() []PublicAPIV1IncidentsSuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Siblings
}