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

// checks if the ContractDeliverGood type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractDeliverGood{}

// ContractDeliverGood The details of a delivery contract. Includes the type of good, units needed, and the destination.
type ContractDeliverGood struct {
	// The symbol of the trade good to deliver.
	TradeSymbol string `json:"tradeSymbol"`
	// The destination where goods need to be delivered.
	DestinationSymbol string `json:"destinationSymbol"`
	// The number of units that need to be delivered on this contract.
	UnitsRequired int32 `json:"unitsRequired"`
	// The number of units fulfilled on this contract.
	UnitsFulfilled int32 `json:"unitsFulfilled"`
}

// NewContractDeliverGood instantiates a new ContractDeliverGood object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractDeliverGood(tradeSymbol string, destinationSymbol string, unitsRequired int32, unitsFulfilled int32) *ContractDeliverGood {
	this := ContractDeliverGood{}
	this.TradeSymbol = tradeSymbol
	this.DestinationSymbol = destinationSymbol
	this.UnitsRequired = unitsRequired
	this.UnitsFulfilled = unitsFulfilled
	return &this
}

// NewContractDeliverGoodWithDefaults instantiates a new ContractDeliverGood object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractDeliverGoodWithDefaults() *ContractDeliverGood {
	this := ContractDeliverGood{}
	return &this
}

// GetTradeSymbol returns the TradeSymbol field value
func (o *ContractDeliverGood) GetTradeSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeSymbol
}

// GetTradeSymbolOk returns a tuple with the TradeSymbol field value
// and a boolean to check if the value has been set.
func (o *ContractDeliverGood) GetTradeSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeSymbol, true
}

// SetTradeSymbol sets field value
func (o *ContractDeliverGood) SetTradeSymbol(v string) {
	o.TradeSymbol = v
}

// GetDestinationSymbol returns the DestinationSymbol field value
func (o *ContractDeliverGood) GetDestinationSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationSymbol
}

// GetDestinationSymbolOk returns a tuple with the DestinationSymbol field value
// and a boolean to check if the value has been set.
func (o *ContractDeliverGood) GetDestinationSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationSymbol, true
}

// SetDestinationSymbol sets field value
func (o *ContractDeliverGood) SetDestinationSymbol(v string) {
	o.DestinationSymbol = v
}

// GetUnitsRequired returns the UnitsRequired field value
func (o *ContractDeliverGood) GetUnitsRequired() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitsRequired
}

// GetUnitsRequiredOk returns a tuple with the UnitsRequired field value
// and a boolean to check if the value has been set.
func (o *ContractDeliverGood) GetUnitsRequiredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitsRequired, true
}

// SetUnitsRequired sets field value
func (o *ContractDeliverGood) SetUnitsRequired(v int32) {
	o.UnitsRequired = v
}

// GetUnitsFulfilled returns the UnitsFulfilled field value
func (o *ContractDeliverGood) GetUnitsFulfilled() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitsFulfilled
}

// GetUnitsFulfilledOk returns a tuple with the UnitsFulfilled field value
// and a boolean to check if the value has been set.
func (o *ContractDeliverGood) GetUnitsFulfilledOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitsFulfilled, true
}

// SetUnitsFulfilled sets field value
func (o *ContractDeliverGood) SetUnitsFulfilled(v int32) {
	o.UnitsFulfilled = v
}

func (o ContractDeliverGood) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractDeliverGood) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tradeSymbol"] = o.TradeSymbol
	toSerialize["destinationSymbol"] = o.DestinationSymbol
	toSerialize["unitsRequired"] = o.UnitsRequired
	toSerialize["unitsFulfilled"] = o.UnitsFulfilled
	return toSerialize, nil
}

type NullableContractDeliverGood struct {
	value *ContractDeliverGood
	isSet bool
}

func (v NullableContractDeliverGood) Get() *ContractDeliverGood {
	return v.value
}

func (v *NullableContractDeliverGood) Set(val *ContractDeliverGood) {
	v.value = val
	v.isSet = true
}

func (v NullableContractDeliverGood) IsSet() bool {
	return v.isSet
}

func (v *NullableContractDeliverGood) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractDeliverGood(val *ContractDeliverGood) *NullableContractDeliverGood {
	return &NullableContractDeliverGood{value: val, isSet: true}
}

func (v NullableContractDeliverGood) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractDeliverGood) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


