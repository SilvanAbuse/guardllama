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

// DhcpSearchV6 struct for DhcpSearchV6
type DhcpSearchV6 struct {
	OtherServer *DhcpSearchResultOtherServer `json:"other_server,omitempty"`
}

// NewDhcpSearchV6 instantiates a new DhcpSearchV6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSearchV6() *DhcpSearchV6 {
	this := DhcpSearchV6{}
	return &this
}

// NewDhcpSearchV6WithDefaults instantiates a new DhcpSearchV6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSearchV6WithDefaults() *DhcpSearchV6 {
	this := DhcpSearchV6{}
	return &this
}

// GetOtherServer returns the OtherServer field value if set, zero value otherwise.
func (o *DhcpSearchV6) GetOtherServer() DhcpSearchResultOtherServer {
	if o == nil || o.OtherServer == nil {
		var ret DhcpSearchResultOtherServer
		return ret
	}
	return *o.OtherServer
}

// GetOtherServerOk returns a tuple with the OtherServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSearchV6) GetOtherServerOk() (*DhcpSearchResultOtherServer, bool) {
	if o == nil || o.OtherServer == nil {
		return nil, false
	}
	return o.OtherServer, true
}

// HasOtherServer returns a boolean if a field has been set.
func (o *DhcpSearchV6) HasOtherServer() bool {
	if o != nil && o.OtherServer != nil {
		return true
	}

	return false
}

// SetOtherServer gets a reference to the given DhcpSearchResultOtherServer and assigns it to the OtherServer field.
func (o *DhcpSearchV6) SetOtherServer(v DhcpSearchResultOtherServer) {
	o.OtherServer = &v
}

func (o DhcpSearchV6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OtherServer != nil {
		toSerialize["other_server"] = o.OtherServer
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpSearchV6 struct {
	value *DhcpSearchV6
	isSet bool
}

func (v NullableDhcpSearchV6) Get() *DhcpSearchV6 {
	return v.value
}

func (v *NullableDhcpSearchV6) Set(val *DhcpSearchV6) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSearchV6) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSearchV6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSearchV6(val *DhcpSearchV6) *NullableDhcpSearchV6 {
	return &NullableDhcpSearchV6{value: val, isSet: true}
}

func (v NullableDhcpSearchV6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSearchV6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


