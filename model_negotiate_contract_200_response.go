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

// checks if the NegotiateContract200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegotiateContract200Response{}

// NegotiateContract200Response 
type NegotiateContract200Response struct {
	Data NegotiateContract200ResponseData `json:"data"`
}

// NewNegotiateContract200Response instantiates a new NegotiateContract200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegotiateContract200Response(data NegotiateContract200ResponseData) *NegotiateContract200Response {
	this := NegotiateContract200Response{}
	this.Data = data
	return &this
}

// NewNegotiateContract200ResponseWithDefaults instantiates a new NegotiateContract200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegotiateContract200ResponseWithDefaults() *NegotiateContract200Response {
	this := NegotiateContract200Response{}
	return &this
}

// GetData returns the Data field value
func (o *NegotiateContract200Response) GetData() NegotiateContract200ResponseData {
	if o == nil {
		var ret NegotiateContract200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NegotiateContract200Response) GetDataOk() (*NegotiateContract200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NegotiateContract200Response) SetData(v NegotiateContract200ResponseData) {
	o.Data = v
}

func (o NegotiateContract200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegotiateContract200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableNegotiateContract200Response struct {
	value *NegotiateContract200Response
	isSet bool
}

func (v NullableNegotiateContract200Response) Get() *NegotiateContract200Response {
	return v.value
}

func (v *NullableNegotiateContract200Response) Set(val *NegotiateContract200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableNegotiateContract200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableNegotiateContract200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegotiateContract200Response(val *NegotiateContract200Response) *NullableNegotiateContract200Response {
	return &NullableNegotiateContract200Response{value: val, isSet: true}
}

func (v NullableNegotiateContract200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegotiateContract200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


