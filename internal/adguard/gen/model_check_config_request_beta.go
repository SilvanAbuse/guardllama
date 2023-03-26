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

// CheckConfigRequestBeta Configuration to be checked
type CheckConfigRequestBeta struct {
	Dns *CheckConfigRequestInfoBeta `json:"dns,omitempty"`
	Web *CheckConfigRequestInfoBeta `json:"web,omitempty"`
	SetStaticIp *bool `json:"set_static_ip,omitempty"`
}

// NewCheckConfigRequestBeta instantiates a new CheckConfigRequestBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckConfigRequestBeta() *CheckConfigRequestBeta {
	this := CheckConfigRequestBeta{}
	return &this
}

// NewCheckConfigRequestBetaWithDefaults instantiates a new CheckConfigRequestBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckConfigRequestBetaWithDefaults() *CheckConfigRequestBeta {
	this := CheckConfigRequestBeta{}
	return &this
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *CheckConfigRequestBeta) GetDns() CheckConfigRequestInfoBeta {
	if o == nil || o.Dns == nil {
		var ret CheckConfigRequestInfoBeta
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConfigRequestBeta) GetDnsOk() (*CheckConfigRequestInfoBeta, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *CheckConfigRequestBeta) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given CheckConfigRequestInfoBeta and assigns it to the Dns field.
func (o *CheckConfigRequestBeta) SetDns(v CheckConfigRequestInfoBeta) {
	o.Dns = &v
}

// GetWeb returns the Web field value if set, zero value otherwise.
func (o *CheckConfigRequestBeta) GetWeb() CheckConfigRequestInfoBeta {
	if o == nil || o.Web == nil {
		var ret CheckConfigRequestInfoBeta
		return ret
	}
	return *o.Web
}

// GetWebOk returns a tuple with the Web field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConfigRequestBeta) GetWebOk() (*CheckConfigRequestInfoBeta, bool) {
	if o == nil || o.Web == nil {
		return nil, false
	}
	return o.Web, true
}

// HasWeb returns a boolean if a field has been set.
func (o *CheckConfigRequestBeta) HasWeb() bool {
	if o != nil && o.Web != nil {
		return true
	}

	return false
}

// SetWeb gets a reference to the given CheckConfigRequestInfoBeta and assigns it to the Web field.
func (o *CheckConfigRequestBeta) SetWeb(v CheckConfigRequestInfoBeta) {
	o.Web = &v
}

// GetSetStaticIp returns the SetStaticIp field value if set, zero value otherwise.
func (o *CheckConfigRequestBeta) GetSetStaticIp() bool {
	if o == nil || o.SetStaticIp == nil {
		var ret bool
		return ret
	}
	return *o.SetStaticIp
}

// GetSetStaticIpOk returns a tuple with the SetStaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConfigRequestBeta) GetSetStaticIpOk() (*bool, bool) {
	if o == nil || o.SetStaticIp == nil {
		return nil, false
	}
	return o.SetStaticIp, true
}

// HasSetStaticIp returns a boolean if a field has been set.
func (o *CheckConfigRequestBeta) HasSetStaticIp() bool {
	if o != nil && o.SetStaticIp != nil {
		return true
	}

	return false
}

// SetSetStaticIp gets a reference to the given bool and assigns it to the SetStaticIp field.
func (o *CheckConfigRequestBeta) SetSetStaticIp(v bool) {
	o.SetStaticIp = &v
}

func (o CheckConfigRequestBeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	if o.Web != nil {
		toSerialize["web"] = o.Web
	}
	if o.SetStaticIp != nil {
		toSerialize["set_static_ip"] = o.SetStaticIp
	}
	return json.Marshal(toSerialize)
}

type NullableCheckConfigRequestBeta struct {
	value *CheckConfigRequestBeta
	isSet bool
}

func (v NullableCheckConfigRequestBeta) Get() *CheckConfigRequestBeta {
	return v.value
}

func (v *NullableCheckConfigRequestBeta) Set(val *CheckConfigRequestBeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckConfigRequestBeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckConfigRequestBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckConfigRequestBeta(val *CheckConfigRequestBeta) *NullableCheckConfigRequestBeta {
	return &NullableCheckConfigRequestBeta{value: val, isSet: true}
}

func (v NullableCheckConfigRequestBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckConfigRequestBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


