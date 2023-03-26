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

// GetVersionRequest /version.json request data
type GetVersionRequest struct {
	// If false, server will check for a new version data only once in several hours. 
	RecheckNow *bool `json:"recheck_now,omitempty"`
}

// NewGetVersionRequest instantiates a new GetVersionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVersionRequest() *GetVersionRequest {
	this := GetVersionRequest{}
	return &this
}

// NewGetVersionRequestWithDefaults instantiates a new GetVersionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVersionRequestWithDefaults() *GetVersionRequest {
	this := GetVersionRequest{}
	return &this
}

// GetRecheckNow returns the RecheckNow field value if set, zero value otherwise.
func (o *GetVersionRequest) GetRecheckNow() bool {
	if o == nil || o.RecheckNow == nil {
		var ret bool
		return ret
	}
	return *o.RecheckNow
}

// GetRecheckNowOk returns a tuple with the RecheckNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVersionRequest) GetRecheckNowOk() (*bool, bool) {
	if o == nil || o.RecheckNow == nil {
		return nil, false
	}
	return o.RecheckNow, true
}

// HasRecheckNow returns a boolean if a field has been set.
func (o *GetVersionRequest) HasRecheckNow() bool {
	if o != nil && o.RecheckNow != nil {
		return true
	}

	return false
}

// SetRecheckNow gets a reference to the given bool and assigns it to the RecheckNow field.
func (o *GetVersionRequest) SetRecheckNow(v bool) {
	o.RecheckNow = &v
}

func (o GetVersionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecheckNow != nil {
		toSerialize["recheck_now"] = o.RecheckNow
	}
	return json.Marshal(toSerialize)
}

type NullableGetVersionRequest struct {
	value *GetVersionRequest
	isSet bool
}

func (v NullableGetVersionRequest) Get() *GetVersionRequest {
	return v.value
}

func (v *NullableGetVersionRequest) Set(val *GetVersionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVersionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVersionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVersionRequest(val *GetVersionRequest) *NullableGetVersionRequest {
	return &NullableGetVersionRequest{value: val, isSet: true}
}

func (v NullableGetVersionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVersionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


