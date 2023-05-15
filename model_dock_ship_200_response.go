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

// checks if the DockShip200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DockShip200Response{}

// DockShip200Response 
type DockShip200Response struct {
	Data OrbitShip200ResponseData `json:"data"`
}

// NewDockShip200Response instantiates a new DockShip200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockShip200Response(data OrbitShip200ResponseData) *DockShip200Response {
	this := DockShip200Response{}
	this.Data = data
	return &this
}

// NewDockShip200ResponseWithDefaults instantiates a new DockShip200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockShip200ResponseWithDefaults() *DockShip200Response {
	this := DockShip200Response{}
	return &this
}

// GetData returns the Data field value
func (o *DockShip200Response) GetData() OrbitShip200ResponseData {
	if o == nil {
		var ret OrbitShip200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DockShip200Response) GetDataOk() (*OrbitShip200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DockShip200Response) SetData(v OrbitShip200ResponseData) {
	o.Data = v
}

func (o DockShip200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockShip200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableDockShip200Response struct {
	value *DockShip200Response
	isSet bool
}

func (v NullableDockShip200Response) Get() *DockShip200Response {
	return v.value
}

func (v *NullableDockShip200Response) Set(val *DockShip200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDockShip200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDockShip200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockShip200Response(val *DockShip200Response) *NullableDockShip200Response {
	return &NullableDockShip200Response{value: val, isSet: true}
}

func (v NullableDockShip200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockShip200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


