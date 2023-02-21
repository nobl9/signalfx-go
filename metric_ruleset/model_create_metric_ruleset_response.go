/*
Metric Ruleset API

Metric ruleset API 

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metric_ruleset

import (
	"encoding/json"
)

// CreateMetricRulesetResponse Create metric ruleset response body 
type CreateMetricRulesetResponse struct {
	// Aggregation rules in the ruleset 
	AggregationRules []AggregationRule `json:"aggregationRules,omitempty"`
	// User ID of the user who created this metric ruleset 
	Creator *string `json:"creator,omitempty"`
	// Name of the user who created this metric ruleset 
	CreatorName *string `json:"creatorName,omitempty"`
	// Date and time when this ruleset was created  
	Created *int64 `json:"created,omitempty"`
	// Ruleset ID 
	Id *string `json:"id,omitempty"`
	// ID of the user who last updated the ruleset 
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// Name of the user who last updated the ruleset 
	LastUpdatedByName *string `json:"lastUpdatedByName,omitempty"`
	// Time at which this ruleset was last updated 
	LastUpdated *int64 `json:"lastUpdated,omitempty"`
	// Name of the metric 
	MetricName *string `json:"metricName,omitempty"`
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
	// Version of the ruleset 
	Version *int64 `json:"version,omitempty"`
}

// NewCreateMetricRulesetResponse instantiates a new CreateMetricRulesetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMetricRulesetResponse() *CreateMetricRulesetResponse {
	this := CreateMetricRulesetResponse{}
	return &this
}

// NewCreateMetricRulesetResponseWithDefaults instantiates a new CreateMetricRulesetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMetricRulesetResponseWithDefaults() *CreateMetricRulesetResponse {
	this := CreateMetricRulesetResponse{}
	return &this
}

// GetAggregationRules returns the AggregationRules field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetAggregationRules() []AggregationRule {
	if o == nil || isNil(o.AggregationRules) {
		var ret []AggregationRule
		return ret
	}
	return o.AggregationRules
}

// GetAggregationRulesOk returns a tuple with the AggregationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetAggregationRulesOk() ([]AggregationRule, bool) {
	if o == nil || isNil(o.AggregationRules) {
    return nil, false
	}
	return o.AggregationRules, true
}

// HasAggregationRules returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasAggregationRules() bool {
	if o != nil && !isNil(o.AggregationRules) {
		return true
	}

	return false
}

// SetAggregationRules gets a reference to the given []AggregationRule and assigns it to the AggregationRules field.
func (o *CreateMetricRulesetResponse) SetAggregationRules(v []AggregationRule) {
	o.AggregationRules = v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetCreator() string {
	if o == nil || isNil(o.Creator) {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetCreatorOk() (*string, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *CreateMetricRulesetResponse) SetCreator(v string) {
	o.Creator = &v
}

// GetCreatorName returns the CreatorName field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetCreatorName() string {
	if o == nil || isNil(o.CreatorName) {
		var ret string
		return ret
	}
	return *o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetCreatorNameOk() (*string, bool) {
	if o == nil || isNil(o.CreatorName) {
    return nil, false
	}
	return o.CreatorName, true
}

// HasCreatorName returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasCreatorName() bool {
	if o != nil && !isNil(o.CreatorName) {
		return true
	}

	return false
}

// SetCreatorName gets a reference to the given string and assigns it to the CreatorName field.
func (o *CreateMetricRulesetResponse) SetCreatorName(v string) {
	o.CreatorName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetCreated() int64 {
	if o == nil || isNil(o.Created) {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetCreatedOk() (*int64, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *CreateMetricRulesetResponse) SetCreated(v int64) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateMetricRulesetResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetLastUpdatedBy() string {
	if o == nil || isNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdatedBy) {
    return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasLastUpdatedBy() bool {
	if o != nil && !isNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *CreateMetricRulesetResponse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetLastUpdatedByName returns the LastUpdatedByName field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetLastUpdatedByName() string {
	if o == nil || isNil(o.LastUpdatedByName) {
		var ret string
		return ret
	}
	return *o.LastUpdatedByName
}

// GetLastUpdatedByNameOk returns a tuple with the LastUpdatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetLastUpdatedByNameOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdatedByName) {
    return nil, false
	}
	return o.LastUpdatedByName, true
}

// HasLastUpdatedByName returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasLastUpdatedByName() bool {
	if o != nil && !isNil(o.LastUpdatedByName) {
		return true
	}

	return false
}

// SetLastUpdatedByName gets a reference to the given string and assigns it to the LastUpdatedByName field.
func (o *CreateMetricRulesetResponse) SetLastUpdatedByName(v string) {
	o.LastUpdatedByName = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetLastUpdated() int64 {
	if o == nil || isNil(o.LastUpdated) {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || isNil(o.LastUpdated) {
    return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasLastUpdated() bool {
	if o != nil && !isNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *CreateMetricRulesetResponse) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetMetricName() string {
	if o == nil || isNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetMetricNameOk() (*string, bool) {
	if o == nil || isNil(o.MetricName) {
    return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasMetricName() bool {
	if o != nil && !isNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *CreateMetricRulesetResponse) SetMetricName(v string) {
	o.MetricName = &v
}

// GetRoutingRule returns the RoutingRule field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetRoutingRule() RoutingRule {
	if o == nil || isNil(o.RoutingRule) {
		var ret RoutingRule
		return ret
	}
	return *o.RoutingRule
}

// GetRoutingRuleOk returns a tuple with the RoutingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetRoutingRuleOk() (*RoutingRule, bool) {
	if o == nil || isNil(o.RoutingRule) {
    return nil, false
	}
	return o.RoutingRule, true
}

// HasRoutingRule returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasRoutingRule() bool {
	if o != nil && !isNil(o.RoutingRule) {
		return true
	}

	return false
}

// SetRoutingRule gets a reference to the given RoutingRule and assigns it to the RoutingRule field.
func (o *CreateMetricRulesetResponse) SetRoutingRule(v RoutingRule) {
	o.RoutingRule = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateMetricRulesetResponse) GetVersion() int64 {
	if o == nil || isNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMetricRulesetResponse) GetVersionOk() (*int64, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateMetricRulesetResponse) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CreateMetricRulesetResponse) SetVersion(v int64) {
	o.Version = &v
}

func (o CreateMetricRulesetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AggregationRules) {
		toSerialize["aggregationRules"] = o.AggregationRules
	}
	if !isNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !isNil(o.CreatorName) {
		toSerialize["creatorName"] = o.CreatorName
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if !isNil(o.LastUpdatedByName) {
		toSerialize["lastUpdatedByName"] = o.LastUpdatedByName
	}
	if !isNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !isNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !isNil(o.RoutingRule) {
		toSerialize["routingRule"] = o.RoutingRule
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMetricRulesetResponse struct {
	value *CreateMetricRulesetResponse
	isSet bool
}

func (v NullableCreateMetricRulesetResponse) Get() *CreateMetricRulesetResponse {
	return v.value
}

func (v *NullableCreateMetricRulesetResponse) Set(val *CreateMetricRulesetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMetricRulesetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMetricRulesetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMetricRulesetResponse(val *CreateMetricRulesetResponse) *NullableCreateMetricRulesetResponse {
	return &NullableCreateMetricRulesetResponse{value: val, isSet: true}
}

func (v NullableCreateMetricRulesetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMetricRulesetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

