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
	"fmt"
)

// ShipRole The registered role of the ship
type ShipRole string

// List of ShipRole
const (
	SHIPROLE_FABRICATOR ShipRole = "FABRICATOR"
	SHIPROLE_HARVESTER ShipRole = "HARVESTER"
	SHIPROLE_HAULER ShipRole = "HAULER"
	SHIPROLE_INTERCEPTOR ShipRole = "INTERCEPTOR"
	SHIPROLE_EXCAVATOR ShipRole = "EXCAVATOR"
	SHIPROLE_TRANSPORT ShipRole = "TRANSPORT"
	SHIPROLE_REPAIR ShipRole = "REPAIR"
	SHIPROLE_SURVEYOR ShipRole = "SURVEYOR"
	SHIPROLE_COMMAND ShipRole = "COMMAND"
	SHIPROLE_CARRIER ShipRole = "CARRIER"
	SHIPROLE_PATROL ShipRole = "PATROL"
	SHIPROLE_SATELLITE ShipRole = "SATELLITE"
	SHIPROLE_EXPLORER ShipRole = "EXPLORER"
	SHIPROLE_REFINERY ShipRole = "REFINERY"
)

// All allowed values of ShipRole enum
var AllowedShipRoleEnumValues = []ShipRole{
	"FABRICATOR",
	"HARVESTER",
	"HAULER",
	"INTERCEPTOR",
	"EXCAVATOR",
	"TRANSPORT",
	"REPAIR",
	"SURVEYOR",
	"COMMAND",
	"CARRIER",
	"PATROL",
	"SATELLITE",
	"EXPLORER",
	"REFINERY",
}

func (v *ShipRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipRole(value)
	for _, existing := range AllowedShipRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipRole", value)
}

// NewShipRoleFromValue returns a pointer to a valid ShipRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipRoleFromValue(v string) (*ShipRole, error) {
	ev := ShipRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipRole: valid values are %v", v, AllowedShipRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipRole) IsValid() bool {
	for _, existing := range AllowedShipRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipRole value
func (v ShipRole) Ptr() *ShipRole {
	return &v
}

type NullableShipRole struct {
	value *ShipRole
	isSet bool
}

func (v NullableShipRole) Get() *ShipRole {
	return v.value
}

func (v *NullableShipRole) Set(val *ShipRole) {
	v.value = val
	v.isSet = true
}

func (v NullableShipRole) IsSet() bool {
	return v.isSet
}

func (v *NullableShipRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipRole(val *ShipRole) *NullableShipRole {
	return &NullableShipRole{value: val, isSet: true}
}

func (v NullableShipRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

