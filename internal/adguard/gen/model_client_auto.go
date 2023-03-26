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

// ClientAuto Auto-Client information
type ClientAuto struct {
	// IP address
	Ip *string `json:"ip,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// The source of this information
	Source *string `json:"source,omitempty"`
	WhoisInfo *map[string]string `json:"whois_info,omitempty"`
}

// NewClientAuto instantiates a new ClientAuto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientAuto() *ClientAuto {
	this := ClientAuto{}
	return &this
}

// NewClientAutoWithDefaults instantiates a new ClientAuto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientAutoWithDefaults() *ClientAuto {
	this := ClientAuto{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ClientAuto) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuto) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ClientAuto) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *ClientAuto) SetIp(v string) {
	o.Ip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientAuto) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuto) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientAuto) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientAuto) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ClientAuto) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuto) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ClientAuto) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ClientAuto) SetSource(v string) {
	o.Source = &v
}

// GetWhoisInfo returns the WhoisInfo field value if set, zero value otherwise.
func (o *ClientAuto) GetWhoisInfo() map[string]string {
	if o == nil || o.WhoisInfo == nil {
		var ret map[string]string
		return ret
	}
	return *o.WhoisInfo
}

// GetWhoisInfoOk returns a tuple with the WhoisInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuto) GetWhoisInfoOk() (*map[string]string, bool) {
	if o == nil || o.WhoisInfo == nil {
		return nil, false
	}
	return o.WhoisInfo, true
}

// HasWhoisInfo returns a boolean if a field has been set.
func (o *ClientAuto) HasWhoisInfo() bool {
	if o != nil && o.WhoisInfo != nil {
		return true
	}

	return false
}

// SetWhoisInfo gets a reference to the given map[string]string and assigns it to the WhoisInfo field.
func (o *ClientAuto) SetWhoisInfo(v map[string]string) {
	o.WhoisInfo = &v
}

func (o ClientAuto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.WhoisInfo != nil {
		toSerialize["whois_info"] = o.WhoisInfo
	}
	return json.Marshal(toSerialize)
}

type NullableClientAuto struct {
	value *ClientAuto
	isSet bool
}

func (v NullableClientAuto) Get() *ClientAuto {
	return v.value
}

func (v *NullableClientAuto) Set(val *ClientAuto) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAuto) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAuto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAuto(val *ClientAuto) *NullableClientAuto {
	return &NullableClientAuto{value: val, isSet: true}
}

func (v NullableClientAuto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAuto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


