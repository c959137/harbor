// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package security

// Context abstracts the operations related with authN and authZ
type Context interface {
	// IsAuthenticated returns whether the context has been authenticated or not
	IsAuthenticated() bool
	// IsSysAdmin returns whether the user is system admin
	IsSysAdmin() bool
	// HasReadPerm returns whether the user has read permission to the project
	// whose ID is projectID
	HasReadPerm(projectID int64) bool
	// HasWritePerm  returns whether the user has write permission to the project
	// whose ID is projectID
	HasWritePerm(projectID int64) bool
	// HasAllPerm  returns whether the user has all permissions to the project
	// whose ID is projectID
	HasAllPerm(projectID int64) bool
}
