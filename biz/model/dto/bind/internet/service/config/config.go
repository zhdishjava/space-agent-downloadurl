// Copyright (c) 2022 Institute of Software, Chinese Academy of Sciences (ISCAS)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import "agent/biz/model/dto/pair"

type ConfigReq struct {
	EnableInternetAccess bool   `json:"enableInternetAccess"`
	ClientUUID           string `json:"clientUUID"`
	PlatformApiBase      string `json:"platformApiBase"` // Platform api url setting, e.g. "https://ao.space"`
}

type ConfigRsp struct {
	EnableInternetAccess bool            `json:"enableInternetAccess"`
	EnableP2P            bool            `json:"enableP2P"`
	EnableLAN            bool            `json:"enableLAN"`
	UserDomain           string          `json:"userDomain"`
	ConnectedNetwork     []*pair.Network `json:"connectedNetwork"`
}
