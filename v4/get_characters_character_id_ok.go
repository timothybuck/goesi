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

package goesiv4

import (
	"time"
)

/* 200 ok object */
type GetCharactersCharacterIdOk struct {
	AllianceId     int32     `json:"alliance_id,omitempty"`     /* The character's alliance ID */
	AncestryId     int32     `json:"ancestry_id,omitempty"`     /* ancestry_id integer */
	Birthday       time.Time `json:"birthday,omitempty"`        /* Creation date of the character */
	BloodlineId    int32     `json:"bloodline_id,omitempty"`    /* bloodline_id integer */
	CorporationId  int32     `json:"corporation_id,omitempty"`  /* The character's corporation ID */
	Description    string    `json:"description,omitempty"`     /* description string */
	Gender         string    `json:"gender,omitempty"`          /* gender string */
	Name           string    `json:"name,omitempty"`            /* name string */
	RaceId         int32     `json:"race_id,omitempty"`         /* race_id integer */
	SecurityStatus float32   `json:"security_status,omitempty"` /* security_status number */
}
