// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NullableSignalsAPITargetEntity struct {
	ID         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Type       *string `json:"type,omitempty"`
	TeamID     *string `json:"team_id,omitempty"`
	IsPageable *bool   `json:"is_pageable,omitempty"`
}

func (o *NullableSignalsAPITargetEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NullableSignalsAPITargetEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *NullableSignalsAPITargetEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *NullableSignalsAPITargetEntity) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *NullableSignalsAPITargetEntity) GetIsPageable() *bool {
	if o == nil {
		return nil
	}
	return o.IsPageable
}
