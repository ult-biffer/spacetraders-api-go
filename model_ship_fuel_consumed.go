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
	"time"
)

// checks if the ShipFuelConsumed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipFuelConsumed{}

// ShipFuelConsumed struct for ShipFuelConsumed
type ShipFuelConsumed struct {
	// The amount of fuel consumed by the most recent transit or action.
	Amount int32 `json:"amount"`
	// The time at which the fuel was consumed.
	Timestamp time.Time `json:"timestamp"`
}

// NewShipFuelConsumed instantiates a new ShipFuelConsumed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipFuelConsumed(amount int32, timestamp time.Time) *ShipFuelConsumed {
	this := ShipFuelConsumed{}
	this.Amount = amount
	this.Timestamp = timestamp
	return &this
}

// NewShipFuelConsumedWithDefaults instantiates a new ShipFuelConsumed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipFuelConsumedWithDefaults() *ShipFuelConsumed {
	this := ShipFuelConsumed{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ShipFuelConsumed) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ShipFuelConsumed) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ShipFuelConsumed) SetAmount(v int32) {
	o.Amount = v
}

// GetTimestamp returns the Timestamp field value
func (o *ShipFuelConsumed) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ShipFuelConsumed) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ShipFuelConsumed) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o ShipFuelConsumed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipFuelConsumed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableShipFuelConsumed struct {
	value *ShipFuelConsumed
	isSet bool
}

func (v NullableShipFuelConsumed) Get() *ShipFuelConsumed {
	return v.value
}

func (v *NullableShipFuelConsumed) Set(val *ShipFuelConsumed) {
	v.value = val
	v.isSet = true
}

func (v NullableShipFuelConsumed) IsSet() bool {
	return v.isSet
}

func (v *NullableShipFuelConsumed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipFuelConsumed(val *ShipFuelConsumed) *NullableShipFuelConsumed {
	return &NullableShipFuelConsumed{value: val, isSet: true}
}

func (v NullableShipFuelConsumed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipFuelConsumed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


