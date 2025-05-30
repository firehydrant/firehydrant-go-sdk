// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// NullableRulesRuleEntityLogic - An unstructured object of key/value pairs describing the logic for applying the rule.
type NullableRulesRuleEntityLogic struct {
}

type NullableRulesRuleEntity struct {
	// An unstructured object of key/value pairs describing the logic for applying the rule.
	Logic    *NullableRulesRuleEntityLogic `json:"logic,omitempty"`
	UserData *NullableFHTypesGenericEntity `json:"user_data,omitempty"`
}

func (o *NullableRulesRuleEntity) GetLogic() *NullableRulesRuleEntityLogic {
	if o == nil {
		return nil
	}
	return o.Logic
}

func (o *NullableRulesRuleEntity) GetUserData() *NullableFHTypesGenericEntity {
	if o == nil {
		return nil
	}
	return o.UserData
}
