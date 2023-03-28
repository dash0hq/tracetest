/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PollingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollingInfo{}

// PollingInfo struct for PollingInfo
type PollingInfo struct {
	Type                *string              `json:"type,omitempty"`
	ReasonNextIteration *string              `json:"reasonNextIteration,omitempty"`
	IsComplete          *bool                `json:"isComplete,omitempty"`
	Periodic            *PollingInfoPeriodic `json:"periodic,omitempty"`
}

// NewPollingInfo instantiates a new PollingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollingInfo() *PollingInfo {
	this := PollingInfo{}
	return &this
}

// NewPollingInfoWithDefaults instantiates a new PollingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollingInfoWithDefaults() *PollingInfo {
	this := PollingInfo{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PollingInfo) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollingInfo) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PollingInfo) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PollingInfo) SetType(v string) {
	o.Type = &v
}

// GetReasonNextIteration returns the ReasonNextIteration field value if set, zero value otherwise.
func (o *PollingInfo) GetReasonNextIteration() string {
	if o == nil || isNil(o.ReasonNextIteration) {
		var ret string
		return ret
	}
	return *o.ReasonNextIteration
}

// GetReasonNextIterationOk returns a tuple with the ReasonNextIteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollingInfo) GetReasonNextIterationOk() (*string, bool) {
	if o == nil || isNil(o.ReasonNextIteration) {
		return nil, false
	}
	return o.ReasonNextIteration, true
}

// HasReasonNextIteration returns a boolean if a field has been set.
func (o *PollingInfo) HasReasonNextIteration() bool {
	if o != nil && !isNil(o.ReasonNextIteration) {
		return true
	}

	return false
}

// SetReasonNextIteration gets a reference to the given string and assigns it to the ReasonNextIteration field.
func (o *PollingInfo) SetReasonNextIteration(v string) {
	o.ReasonNextIteration = &v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *PollingInfo) GetIsComplete() bool {
	if o == nil || isNil(o.IsComplete) {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollingInfo) GetIsCompleteOk() (*bool, bool) {
	if o == nil || isNil(o.IsComplete) {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *PollingInfo) HasIsComplete() bool {
	if o != nil && !isNil(o.IsComplete) {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *PollingInfo) SetIsComplete(v bool) {
	o.IsComplete = &v
}

// GetPeriodic returns the Periodic field value if set, zero value otherwise.
func (o *PollingInfo) GetPeriodic() PollingInfoPeriodic {
	if o == nil || isNil(o.Periodic) {
		var ret PollingInfoPeriodic
		return ret
	}
	return *o.Periodic
}

// GetPeriodicOk returns a tuple with the Periodic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollingInfo) GetPeriodicOk() (*PollingInfoPeriodic, bool) {
	if o == nil || isNil(o.Periodic) {
		return nil, false
	}
	return o.Periodic, true
}

// HasPeriodic returns a boolean if a field has been set.
func (o *PollingInfo) HasPeriodic() bool {
	if o != nil && !isNil(o.Periodic) {
		return true
	}

	return false
}

// SetPeriodic gets a reference to the given PollingInfoPeriodic and assigns it to the Periodic field.
func (o *PollingInfo) SetPeriodic(v PollingInfoPeriodic) {
	o.Periodic = &v
}

func (o PollingInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ReasonNextIteration) {
		toSerialize["reasonNextIteration"] = o.ReasonNextIteration
	}
	if !isNil(o.IsComplete) {
		toSerialize["isComplete"] = o.IsComplete
	}
	if !isNil(o.Periodic) {
		toSerialize["periodic"] = o.Periodic
	}
	return toSerialize, nil
}

type NullablePollingInfo struct {
	value *PollingInfo
	isSet bool
}

func (v NullablePollingInfo) Get() *PollingInfo {
	return v.value
}

func (v *NullablePollingInfo) Set(val *PollingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePollingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePollingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollingInfo(val *PollingInfo) *NullablePollingInfo {
	return &NullablePollingInfo{value: val, isSet: true}
}

func (v NullablePollingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
