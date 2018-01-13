/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.5
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

package esi

import (
	"time"
)

/* A list of GetCorporationsCorporationIdOk. */
//easyjson:json
type GetCorporationsCorporationIdOkList []GetCorporationsCorporationIdOk

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdOk struct {
	Name          string    `json:"name,omitempty"`            /* the full name of the corporation */
	Ticker        string    `json:"ticker,omitempty"`          /* the short name of the corporation */
	MemberCount   int32     `json:"member_count,omitempty"`    /* member_count integer */
	CeoId         int32     `json:"ceo_id,omitempty"`          /* ceo_id integer */
	AllianceId    int32     `json:"alliance_id,omitempty"`     /* ID of the alliance that corporation is a member of, if any */
	Description   string    `json:"description,omitempty"`     /* description string */
	TaxRate       float32   `json:"tax_rate,omitempty"`        /* tax_rate number */
	DateFounded   time.Time `json:"date_founded,omitempty"`    /* date_founded string */
	CreatorId     int32     `json:"creator_id,omitempty"`      /* creator_id integer */
	Url           string    `json:"url,omitempty"`             /* url string */
	FactionId     int32     `json:"faction_id,omitempty"`      /* faction_id integer */
	HomeStationId int32     `json:"home_station_id,omitempty"` /* home_station_id integer */
	Shares        int64     `json:"shares,omitempty"`          /* shares integer */
}
