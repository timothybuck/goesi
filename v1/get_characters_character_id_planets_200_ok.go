/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.4
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goesiv1

import (
	"time"
)

/* 200 ok object */
type GetCharactersCharacterIdPlanets200Ok struct {
	LastUpdate    time.Time `json:"last_update,omitempty"`     /* last_update string */
	NumPins       int32     `json:"num_pins,omitempty"`        /* num_pins integer */
	OwnerId       int32     `json:"owner_id,omitempty"`        /* owner_id integer */
	PlanetId      int32     `json:"planet_id,omitempty"`       /* planet_id integer */
	PlanetType    string    `json:"planet_type,omitempty"`     /* planet_type string */
	SolarSystemId int32     `json:"solar_system_id,omitempty"` /* solar_system_id integer */
	UpgradeLevel  int32     `json:"upgrade_level,omitempty"`   /* upgrade_level integer */
}
