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

// DnsInfo200ResponseAllOf struct for DnsInfo200ResponseAllOf
type DnsInfo200ResponseAllOf struct {
	DefaultLocalPtrUpstreams []string `json:"default_local_ptr_upstreams,omitempty"`
}

// NewDnsInfo200ResponseAllOf instantiates a new DnsInfo200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsInfo200ResponseAllOf() *DnsInfo200ResponseAllOf {
	this := DnsInfo200ResponseAllOf{}
	return &this
}

// NewDnsInfo200ResponseAllOfWithDefaults instantiates a new DnsInfo200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsInfo200ResponseAllOfWithDefaults() *DnsInfo200ResponseAllOf {
	this := DnsInfo200ResponseAllOf{}
	return &this
}

// GetDefaultLocalPtrUpstreams returns the DefaultLocalPtrUpstreams field value if set, zero value otherwise.
func (o *DnsInfo200ResponseAllOf) GetDefaultLocalPtrUpstreams() []string {
	if o == nil || o.DefaultLocalPtrUpstreams == nil {
		var ret []string
		return ret
	}
	return o.DefaultLocalPtrUpstreams
}

// GetDefaultLocalPtrUpstreamsOk returns a tuple with the DefaultLocalPtrUpstreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsInfo200ResponseAllOf) GetDefaultLocalPtrUpstreamsOk() ([]string, bool) {
	if o == nil || o.DefaultLocalPtrUpstreams == nil {
		return nil, false
	}
	return o.DefaultLocalPtrUpstreams, true
}

// HasDefaultLocalPtrUpstreams returns a boolean if a field has been set.
func (o *DnsInfo200ResponseAllOf) HasDefaultLocalPtrUpstreams() bool {
	if o != nil && o.DefaultLocalPtrUpstreams != nil {
		return true
	}

	return false
}

// SetDefaultLocalPtrUpstreams gets a reference to the given []string and assigns it to the DefaultLocalPtrUpstreams field.
func (o *DnsInfo200ResponseAllOf) SetDefaultLocalPtrUpstreams(v []string) {
	o.DefaultLocalPtrUpstreams = v
}

func (o DnsInfo200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultLocalPtrUpstreams != nil {
		toSerialize["default_local_ptr_upstreams"] = o.DefaultLocalPtrUpstreams
	}
	return json.Marshal(toSerialize)
}

type NullableDnsInfo200ResponseAllOf struct {
	value *DnsInfo200ResponseAllOf
	isSet bool
}

func (v NullableDnsInfo200ResponseAllOf) Get() *DnsInfo200ResponseAllOf {
	return v.value
}

func (v *NullableDnsInfo200ResponseAllOf) Set(val *DnsInfo200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsInfo200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsInfo200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsInfo200ResponseAllOf(val *DnsInfo200ResponseAllOf) *NullableDnsInfo200ResponseAllOf {
	return &NullableDnsInfo200ResponseAllOf{value: val, isSet: true}
}

func (v NullableDnsInfo200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsInfo200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


