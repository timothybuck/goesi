/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.0
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

/* A list of PostUiOpenwindowNewmailNewMail. */
//easyjson:json
type PostUiOpenwindowNewmailNewMailList []PostUiOpenwindowNewmailNewMail

/* new_mail object */
//easyjson:json
type PostUiOpenwindowNewmailNewMail struct {
	Body               string  `json:"body,omitempty"`                   /* body string */
	Recipients         []int32 `json:"recipients,omitempty"`             /* recipients array */
	Subject            string  `json:"subject,omitempty"`                /* subject string */
	ToCorpOrAllianceId int32   `json:"to_corp_or_alliance_id,omitempty"` /* to_corp_or_alliance_id integer */
	ToMailingListId    int32   `json:"to_mailing_list_id,omitempty"`     /* Corporations, alliances and mailing lists are all types of mailing groups. You may only send to one mailing group, at a time, so you may fill out either this field or the to_corp_or_alliance_ids field */
}