/*
AdGuard Home

AdGuard Home REST-ish API.  Our admin web interface is built on top of this REST-ish API. 

API version: 0.107
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package adguard

import (
	"encoding/json"
)

// CheckConfigResponseInfo struct for CheckConfigResponseInfo
type CheckConfigResponseInfo struct {
	Status string `json:"status"`
	CanAutofix bool `json:"can_autofix"`
}

// NewCheckConfigResponseInfo instantiates a new CheckConfigResponseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckConfigResponseInfo(status string, canAutofix bool) *CheckConfigResponseInfo {
	this := CheckConfigResponseInfo{}
	this.Status = status
	this.CanAutofix = canAutofix
	return &this
}

// NewCheckConfigResponseInfoWithDefaults instantiates a new CheckConfigResponseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckConfigResponseInfoWithDefaults() *CheckConfigResponseInfo {
	this := CheckConfigResponseInfo{}
	var status string = ""
	this.Status = status
	return &this
}

// GetStatus returns the Status field value
func (o *CheckConfigResponseInfo) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CheckConfigResponseInfo) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CheckConfigResponseInfo) SetStatus(v string) {
	o.Status = v
}

// GetCanAutofix returns the CanAutofix field value
func (o *CheckConfigResponseInfo) GetCanAutofix() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanAutofix
}

// GetCanAutofixOk returns a tuple with the CanAutofix field value
// and a boolean to check if the value has been set.
func (o *CheckConfigResponseInfo) GetCanAutofixOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanAutofix, true
}

// SetCanAutofix sets field value
func (o *CheckConfigResponseInfo) SetCanAutofix(v bool) {
	o.CanAutofix = v
}

func (o CheckConfigResponseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["can_autofix"] = o.CanAutofix
	}
	return json.Marshal(toSerialize)
}

type NullableCheckConfigResponseInfo struct {
	value *CheckConfigResponseInfo
	isSet bool
}

func (v NullableCheckConfigResponseInfo) Get() *CheckConfigResponseInfo {
	return v.value
}

func (v *NullableCheckConfigResponseInfo) Set(val *CheckConfigResponseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckConfigResponseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckConfigResponseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckConfigResponseInfo(val *CheckConfigResponseInfo) *NullableCheckConfigResponseInfo {
	return &NullableCheckConfigResponseInfo{value: val, isSet: true}
}

func (v NullableCheckConfigResponseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckConfigResponseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


