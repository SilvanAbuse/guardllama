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

// ProfileInfo Information about the current user
type ProfileInfo struct {
	Name string `json:"name"`
	Language string `json:"language"`
	// Interface theme
	Theme string `json:"theme"`
}

// NewProfileInfo instantiates a new ProfileInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileInfo(name string, language string, theme string) *ProfileInfo {
	this := ProfileInfo{}
	this.Name = name
	this.Language = language
	this.Theme = theme
	return &this
}

// NewProfileInfoWithDefaults instantiates a new ProfileInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileInfoWithDefaults() *ProfileInfo {
	this := ProfileInfo{}
	return &this
}

// GetName returns the Name field value
func (o *ProfileInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProfileInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProfileInfo) SetName(v string) {
	o.Name = v
}

// GetLanguage returns the Language field value
func (o *ProfileInfo) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *ProfileInfo) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *ProfileInfo) SetLanguage(v string) {
	o.Language = v
}

// GetTheme returns the Theme field value
func (o *ProfileInfo) GetTheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *ProfileInfo) GetThemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *ProfileInfo) SetTheme(v string) {
	o.Theme = v
}

func (o ProfileInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["theme"] = o.Theme
	}
	return json.Marshal(toSerialize)
}

type NullableProfileInfo struct {
	value *ProfileInfo
	isSet bool
}

func (v NullableProfileInfo) Get() *ProfileInfo {
	return v.value
}

func (v *NullableProfileInfo) Set(val *ProfileInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileInfo(val *ProfileInfo) *NullableProfileInfo {
	return &NullableProfileInfo{value: val, isSet: true}
}

func (v NullableProfileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


