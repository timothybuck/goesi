/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.1
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
type GetDogmaAttributesAttributeIdOk struct {

	/* attribute_id integer */
	AttributeId int32 `json:"attribute_id,omitempty"`

	/* default_value number */
	DefaultValue float32 `json:"default_value,omitempty"`

	/* description string */
	Description string `json:"description,omitempty"`

	/* display_name string */
	DisplayName string `json:"display_name,omitempty"`

	/* high_is_good boolean */
	HighIsGood bool `json:"high_is_good,omitempty"`

	/* icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`

	/* name string */
	Name string `json:"name,omitempty"`

	/* published boolean */
	Published bool `json:"published,omitempty"`

	/* stackable boolean */
	Stackable bool `json:"stackable,omitempty"`

	/* unit_id integer */
	UnitId int32 `json:"unit_id,omitempty"`
}
