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

// UpstreamsConfig Upstreams configuration
type UpstreamsConfig struct {
	// Bootstrap servers, port is optional after colon.  Empty value will reset it to default values. 
	BootstrapDns []string `json:"bootstrap_dns"`
	// Upstream servers, port is optional after colon.  Empty value will reset it to default values. 
	UpstreamDns []string `json:"upstream_dns"`
	// Local PTR resolvers, port is optional after colon.  Empty value will reset it to default values. 
	PrivateUpstream []string `json:"private_upstream,omitempty"`
}

// NewUpstreamsConfig instantiates a new UpstreamsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpstreamsConfig(bootstrapDns []string, upstreamDns []string) *UpstreamsConfig {
	this := UpstreamsConfig{}
	this.BootstrapDns = bootstrapDns
	this.UpstreamDns = upstreamDns
	return &this
}

// NewUpstreamsConfigWithDefaults instantiates a new UpstreamsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpstreamsConfigWithDefaults() *UpstreamsConfig {
	this := UpstreamsConfig{}
	return &this
}

// GetBootstrapDns returns the BootstrapDns field value
func (o *UpstreamsConfig) GetBootstrapDns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BootstrapDns
}

// GetBootstrapDnsOk returns a tuple with the BootstrapDns field value
// and a boolean to check if the value has been set.
func (o *UpstreamsConfig) GetBootstrapDnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BootstrapDns, true
}

// SetBootstrapDns sets field value
func (o *UpstreamsConfig) SetBootstrapDns(v []string) {
	o.BootstrapDns = v
}

// GetUpstreamDns returns the UpstreamDns field value
func (o *UpstreamsConfig) GetUpstreamDns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UpstreamDns
}

// GetUpstreamDnsOk returns a tuple with the UpstreamDns field value
// and a boolean to check if the value has been set.
func (o *UpstreamsConfig) GetUpstreamDnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamDns, true
}

// SetUpstreamDns sets field value
func (o *UpstreamsConfig) SetUpstreamDns(v []string) {
	o.UpstreamDns = v
}

// GetPrivateUpstream returns the PrivateUpstream field value if set, zero value otherwise.
func (o *UpstreamsConfig) GetPrivateUpstream() []string {
	if o == nil || o.PrivateUpstream == nil {
		var ret []string
		return ret
	}
	return o.PrivateUpstream
}

// GetPrivateUpstreamOk returns a tuple with the PrivateUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamsConfig) GetPrivateUpstreamOk() ([]string, bool) {
	if o == nil || o.PrivateUpstream == nil {
		return nil, false
	}
	return o.PrivateUpstream, true
}

// HasPrivateUpstream returns a boolean if a field has been set.
func (o *UpstreamsConfig) HasPrivateUpstream() bool {
	if o != nil && o.PrivateUpstream != nil {
		return true
	}

	return false
}

// SetPrivateUpstream gets a reference to the given []string and assigns it to the PrivateUpstream field.
func (o *UpstreamsConfig) SetPrivateUpstream(v []string) {
	o.PrivateUpstream = v
}

func (o UpstreamsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bootstrap_dns"] = o.BootstrapDns
	}
	if true {
		toSerialize["upstream_dns"] = o.UpstreamDns
	}
	if o.PrivateUpstream != nil {
		toSerialize["private_upstream"] = o.PrivateUpstream
	}
	return json.Marshal(toSerialize)
}

type NullableUpstreamsConfig struct {
	value *UpstreamsConfig
	isSet bool
}

func (v NullableUpstreamsConfig) Get() *UpstreamsConfig {
	return v.value
}

func (v *NullableUpstreamsConfig) Set(val *UpstreamsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpstreamsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpstreamsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpstreamsConfig(val *UpstreamsConfig) *NullableUpstreamsConfig {
	return &NullableUpstreamsConfig{value: val, isSet: true}
}

func (v NullableUpstreamsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpstreamsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


