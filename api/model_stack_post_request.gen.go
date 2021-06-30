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

// StackPostRequest struct for StackPostRequest
type StackPostRequest struct {
	OrgID       string   `json:"orgID" yaml:"orgID"`
	Name        string   `json:"name" yaml:"name"`
	Description *string  `json:"description,omitempty" yaml:"description,omitempty"`
	Urls        []string `json:"urls" yaml:"urls"`
}

// NewStackPostRequest instantiates a new StackPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackPostRequest(orgID string, name string, urls []string) *StackPostRequest {
	this := StackPostRequest{}
	this.OrgID = orgID
	this.Name = name
	this.Urls = urls
	return &this
}

// NewStackPostRequestWithDefaults instantiates a new StackPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackPostRequestWithDefaults() *StackPostRequest {
	this := StackPostRequest{}
	return &this
}

// GetOrgID returns the OrgID field value
func (o *StackPostRequest) GetOrgID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value
// and a boolean to check if the value has been set.
func (o *StackPostRequest) GetOrgIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgID, true
}

// SetOrgID sets field value
func (o *StackPostRequest) SetOrgID(v string) {
	o.OrgID = v
}

// GetName returns the Name field value
func (o *StackPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StackPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StackPostRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StackPostRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StackPostRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StackPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetUrls returns the Urls field value
func (o *StackPostRequest) GetUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *StackPostRequest) GetUrlsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *StackPostRequest) SetUrls(v []string) {
	o.Urls = v
}

func (o StackPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orgID"] = o.OrgID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["urls"] = o.Urls
	}
	return json.Marshal(toSerialize)
}

type NullableStackPostRequest struct {
	value *StackPostRequest
	isSet bool
}

func (v NullableStackPostRequest) Get() *StackPostRequest {
	return v.value
}

func (v *NullableStackPostRequest) Set(val *StackPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStackPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStackPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackPostRequest(val *StackPostRequest) *NullableStackPostRequest {
	return &NullableStackPostRequest{value: val, isSet: true}
}

func (v NullableStackPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
