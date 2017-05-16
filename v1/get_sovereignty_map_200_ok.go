/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.6.dev10
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

/* 200 ok object */
type GetSovereigntyMap200Ok struct {
	AllianceId    int32 `json:"alliance_id,omitempty"`    /* alliance_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"` /* corporation_id integer */
	FactionId     int32 `json:"faction_id,omitempty"`     /* faction_id integer */
	SystemId      int32 `json:"system_id,omitempty"`      /* system_id integer */
}
