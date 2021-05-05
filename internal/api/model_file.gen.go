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

// File Represents a source from a single file
type File struct {
	// Type of AST node
	Type *string `json:"type,omitempty"`
	// The name of the file.
	Name    *string        `json:"name,omitempty"`
	Package *PackageClause `json:"package,omitempty"`
	// A list of package imports
	Imports *[]ImportDeclaration `json:"imports,omitempty"`
	// List of Flux statements
	Body *[]Statement `json:"body,omitempty"`
}

// NewFile instantiates a new File object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFile() *File {
	this := File{}
	return &this
}

// NewFileWithDefaults instantiates a new File object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileWithDefaults() *File {
	this := File{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *File) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *File) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *File) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *File) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *File) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *File) SetName(v string) {
	o.Name = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *File) GetPackage() PackageClause {
	if o == nil || o.Package == nil {
		var ret PackageClause
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetPackageOk() (*PackageClause, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *File) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given PackageClause and assigns it to the Package field.
func (o *File) SetPackage(v PackageClause) {
	o.Package = &v
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *File) GetImports() []ImportDeclaration {
	if o == nil || o.Imports == nil {
		var ret []ImportDeclaration
		return ret
	}
	return *o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetImportsOk() (*[]ImportDeclaration, bool) {
	if o == nil || o.Imports == nil {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *File) HasImports() bool {
	if o != nil && o.Imports != nil {
		return true
	}

	return false
}

// SetImports gets a reference to the given []ImportDeclaration and assigns it to the Imports field.
func (o *File) SetImports(v []ImportDeclaration) {
	o.Imports = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *File) GetBody() []Statement {
	if o == nil || o.Body == nil {
		var ret []Statement
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetBodyOk() (*[]Statement, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *File) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given []Statement and assigns it to the Body field.
func (o *File) SetBody(v []Statement) {
	o.Body = &v
}

func (o File) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableFile struct {
	value *File
	isSet bool
}

func (v NullableFile) Get() *File {
	return v.value
}

func (v *NullableFile) Set(val *File) {
	v.value = val
	v.isSet = true
}

func (v NullableFile) IsSet() bool {
	return v.isSet
}

func (v *NullableFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFile(val *File) *NullableFile {
	return &NullableFile{value: val, isSet: true}
}

func (v NullableFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}