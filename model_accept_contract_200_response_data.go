/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spacetraders_api

import (
	"encoding/json"
)

// checks if the AcceptContract200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceptContract200ResponseData{}

// AcceptContract200ResponseData struct for AcceptContract200ResponseData
type AcceptContract200ResponseData struct {
	Agent Agent `json:"agent"`
	Contract Contract `json:"contract"`
}

// NewAcceptContract200ResponseData instantiates a new AcceptContract200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptContract200ResponseData(agent Agent, contract Contract) *AcceptContract200ResponseData {
	this := AcceptContract200ResponseData{}
	this.Agent = agent
	this.Contract = contract
	return &this
}

// NewAcceptContract200ResponseDataWithDefaults instantiates a new AcceptContract200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptContract200ResponseDataWithDefaults() *AcceptContract200ResponseData {
	this := AcceptContract200ResponseData{}
	return &this
}

// GetAgent returns the Agent field value
func (o *AcceptContract200ResponseData) GetAgent() Agent {
	if o == nil {
		var ret Agent
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *AcceptContract200ResponseData) GetAgentOk() (*Agent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *AcceptContract200ResponseData) SetAgent(v Agent) {
	o.Agent = v
}

// GetContract returns the Contract field value
func (o *AcceptContract200ResponseData) GetContract() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *AcceptContract200ResponseData) GetContractOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *AcceptContract200ResponseData) SetContract(v Contract) {
	o.Contract = v
}

func (o AcceptContract200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptContract200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent"] = o.Agent
	toSerialize["contract"] = o.Contract
	return toSerialize, nil
}

type NullableAcceptContract200ResponseData struct {
	value *AcceptContract200ResponseData
	isSet bool
}

func (v NullableAcceptContract200ResponseData) Get() *AcceptContract200ResponseData {
	return v.value
}

func (v *NullableAcceptContract200ResponseData) Set(val *AcceptContract200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptContract200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptContract200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptContract200ResponseData(val *AcceptContract200ResponseData) *NullableAcceptContract200ResponseData {
	return &NullableAcceptContract200ResponseData{value: val, isSet: true}
}

func (v NullableAcceptContract200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptContract200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


