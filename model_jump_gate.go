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

// checks if the JumpGate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JumpGate{}

// JumpGate 
type JumpGate struct {
	// The maximum jump range of the gate.
	JumpRange float32 `json:"jumpRange"`
	// The symbol of the faction that owns the gate.
	FactionSymbol *string `json:"factionSymbol,omitempty"`
	// The systems within range of the gate that have a corresponding gate.
	ConnectedSystems []ConnectedSystem `json:"connectedSystems"`
}

// NewJumpGate instantiates a new JumpGate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJumpGate(jumpRange float32, connectedSystems []ConnectedSystem) *JumpGate {
	this := JumpGate{}
	this.JumpRange = jumpRange
	this.ConnectedSystems = connectedSystems
	return &this
}

// NewJumpGateWithDefaults instantiates a new JumpGate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJumpGateWithDefaults() *JumpGate {
	this := JumpGate{}
	return &this
}

// GetJumpRange returns the JumpRange field value
func (o *JumpGate) GetJumpRange() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.JumpRange
}

// GetJumpRangeOk returns a tuple with the JumpRange field value
// and a boolean to check if the value has been set.
func (o *JumpGate) GetJumpRangeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JumpRange, true
}

// SetJumpRange sets field value
func (o *JumpGate) SetJumpRange(v float32) {
	o.JumpRange = v
}

// GetFactionSymbol returns the FactionSymbol field value if set, zero value otherwise.
func (o *JumpGate) GetFactionSymbol() string {
	if o == nil || IsNil(o.FactionSymbol) {
		var ret string
		return ret
	}
	return *o.FactionSymbol
}

// GetFactionSymbolOk returns a tuple with the FactionSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JumpGate) GetFactionSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.FactionSymbol) {
		return nil, false
	}
	return o.FactionSymbol, true
}

// HasFactionSymbol returns a boolean if a field has been set.
func (o *JumpGate) HasFactionSymbol() bool {
	if o != nil && !IsNil(o.FactionSymbol) {
		return true
	}

	return false
}

// SetFactionSymbol gets a reference to the given string and assigns it to the FactionSymbol field.
func (o *JumpGate) SetFactionSymbol(v string) {
	o.FactionSymbol = &v
}

// GetConnectedSystems returns the ConnectedSystems field value
func (o *JumpGate) GetConnectedSystems() []ConnectedSystem {
	if o == nil {
		var ret []ConnectedSystem
		return ret
	}

	return o.ConnectedSystems
}

// GetConnectedSystemsOk returns a tuple with the ConnectedSystems field value
// and a boolean to check if the value has been set.
func (o *JumpGate) GetConnectedSystemsOk() ([]ConnectedSystem, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedSystems, true
}

// SetConnectedSystems sets field value
func (o *JumpGate) SetConnectedSystems(v []ConnectedSystem) {
	o.ConnectedSystems = v
}

func (o JumpGate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JumpGate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jumpRange"] = o.JumpRange
	if !IsNil(o.FactionSymbol) {
		toSerialize["factionSymbol"] = o.FactionSymbol
	}
	toSerialize["connectedSystems"] = o.ConnectedSystems
	return toSerialize, nil
}

type NullableJumpGate struct {
	value *JumpGate
	isSet bool
}

func (v NullableJumpGate) Get() *JumpGate {
	return v.value
}

func (v *NullableJumpGate) Set(val *JumpGate) {
	v.value = val
	v.isSet = true
}

func (v NullableJumpGate) IsSet() bool {
	return v.isSet
}

func (v *NullableJumpGate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJumpGate(val *JumpGate) *NullableJumpGate {
	return &NullableJumpGate{value: val, isSet: true}
}

func (v NullableJumpGate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJumpGate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


