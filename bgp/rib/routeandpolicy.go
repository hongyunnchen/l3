//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

// route.go
package rib

import (
	"l3/bgp/packet"
)

type RIBInOutRoute struct {
	NLRI   packet.NLRI
	Path   *Path
	PathId uint32
}

type RouteAndPolicy struct {
	RIBInOutRoute
	PolicyList       []string
	PolicyHitCounter int
}

func NewRouteAndPolicy(nlri packet.NLRI, path *Path, pathId uint32) *RouteAndPolicy {
	return &RouteAndPolicy{
		RIBInOutRoute: RIBInOutRoute{
			NLRI:   nlri,
			Path:   path,
			PathId: pathId,
		},
		PolicyList:       make([]string, 0),
		PolicyHitCounter: 0,
	}
}
