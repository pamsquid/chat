// Copyright 2017 Vector Creations Ltd

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authtypes

// Profile represents the profile for a Matrix account.
type Profile struct {
	Localpart   string  `json:"local_part"`
	ServerName  string  `json:"server_name,omitempty"` // NOTSPEC: only set by Pinecone user provider
	DisplayName string  `json:"display_name"`
	AvatarURL   string  `json:"avatar_url"`
	TelNumbers  []int64 `json:"tel_numbers"`
	MortgageFee string  `json:"mortgage_fee"`
}

// FullyQualifiedProfile represents the profile for a Matrix account.
type FullyQualifiedProfile struct {
	UserID      string  `json:"user_id"`
	DisplayName string  `json:"display_name"`
	AvatarURL   string  `json:"avatar_url"`
	TelNumbers  []int64 `json:"tel_numbers"`
}
