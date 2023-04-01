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

// LanguageSettings Language settings object.
type LanguageSettings struct {
	// The current language or the language to set.
	Language string `json:"language"`
}

// NewLanguageSettings instantiates a new LanguageSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageSettings(language string) *LanguageSettings {
	this := LanguageSettings{}
	this.Language = language
	return &this
}

// NewLanguageSettingsWithDefaults instantiates a new LanguageSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageSettingsWithDefaults() *LanguageSettings {
	this := LanguageSettings{}
	return &this
}

// GetLanguage returns the Language field value
func (o *LanguageSettings) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *LanguageSettings) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *LanguageSettings) SetLanguage(v string) {
	o.Language = v
}

func (o LanguageSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["language"] = o.Language
	}
	return json.Marshal(toSerialize)
}

type NullableLanguageSettings struct {
	value *LanguageSettings
	isSet bool
}

func (v NullableLanguageSettings) Get() *LanguageSettings {
	return v.value
}

func (v *NullableLanguageSettings) Set(val *LanguageSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageSettings(val *LanguageSettings) *NullableLanguageSettings {
	return &NullableLanguageSettings{value: val, isSet: true}
}

func (v NullableLanguageSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


