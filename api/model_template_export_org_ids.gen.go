/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TemplateExportOrgIDs struct for TemplateExportOrgIDs
type TemplateExportOrgIDs struct {
	OrgID           *string                        `json:"orgID,omitempty"`
	ResourceFilters *TemplateExportResourceFilters `json:"resourceFilters,omitempty"`
}

// NewTemplateExportOrgIDs instantiates a new TemplateExportOrgIDs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateExportOrgIDs() *TemplateExportOrgIDs {
	this := TemplateExportOrgIDs{}
	return &this
}

// NewTemplateExportOrgIDsWithDefaults instantiates a new TemplateExportOrgIDs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateExportOrgIDsWithDefaults() *TemplateExportOrgIDs {
	this := TemplateExportOrgIDs{}
	return &this
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *TemplateExportOrgIDs) GetOrgID() string {
	if o == nil || o.OrgID == nil {
		var ret string
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateExportOrgIDs) GetOrgIDOk() (*string, bool) {
	if o == nil || o.OrgID == nil {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *TemplateExportOrgIDs) HasOrgID() bool {
	if o != nil && o.OrgID != nil {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given string and assigns it to the OrgID field.
func (o *TemplateExportOrgIDs) SetOrgID(v string) {
	o.OrgID = &v
}

// GetResourceFilters returns the ResourceFilters field value if set, zero value otherwise.
func (o *TemplateExportOrgIDs) GetResourceFilters() TemplateExportResourceFilters {
	if o == nil || o.ResourceFilters == nil {
		var ret TemplateExportResourceFilters
		return ret
	}
	return *o.ResourceFilters
}

// GetResourceFiltersOk returns a tuple with the ResourceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateExportOrgIDs) GetResourceFiltersOk() (*TemplateExportResourceFilters, bool) {
	if o == nil || o.ResourceFilters == nil {
		return nil, false
	}
	return o.ResourceFilters, true
}

// HasResourceFilters returns a boolean if a field has been set.
func (o *TemplateExportOrgIDs) HasResourceFilters() bool {
	if o != nil && o.ResourceFilters != nil {
		return true
	}

	return false
}

// SetResourceFilters gets a reference to the given TemplateExportResourceFilters and assigns it to the ResourceFilters field.
func (o *TemplateExportOrgIDs) SetResourceFilters(v TemplateExportResourceFilters) {
	o.ResourceFilters = &v
}

func (o TemplateExportOrgIDs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgID != nil {
		toSerialize["orgID"] = o.OrgID
	}
	if o.ResourceFilters != nil {
		toSerialize["resourceFilters"] = o.ResourceFilters
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateExportOrgIDs struct {
	value *TemplateExportOrgIDs
	isSet bool
}

func (v NullableTemplateExportOrgIDs) Get() *TemplateExportOrgIDs {
	return v.value
}

func (v *NullableTemplateExportOrgIDs) Set(val *TemplateExportOrgIDs) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateExportOrgIDs) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateExportOrgIDs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateExportOrgIDs(val *TemplateExportOrgIDs) *NullableTemplateExportOrgIDs {
	return &NullableTemplateExportOrgIDs{value: val, isSet: true}
}

func (v NullableTemplateExportOrgIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateExportOrgIDs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
