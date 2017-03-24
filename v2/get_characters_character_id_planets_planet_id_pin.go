/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev14
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

package goesiv2

import (
	"time"
)

/* pin object */
type GetCharactersCharacterIdPlanetsPlanetIdPin struct {

	/* expiry_time string */
	ExpiryTime time.Time `json:"expiry_time,omitempty"`

	ExtractorDetails GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails `json:"extractor_details,omitempty"`

	FactoryDetails GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails `json:"factory_details,omitempty"`

	/* install_time string */
	InstallTime time.Time `json:"install_time,omitempty"`

	/* last_cycle_start string */
	LastCycleStart time.Time `json:"last_cycle_start,omitempty"`

	/* latitude number */
	Latitude float32 `json:"latitude,omitempty"`

	/* longitude number */
	Longitude float32 `json:"longitude,omitempty"`

	/* pin_id integer */
	PinId int64 `json:"pin_id,omitempty"`

	/* schematic_id integer */
	SchematicId int32 `json:"schematic_id,omitempty"`

	/* type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
