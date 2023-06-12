// Copyright Mia srl
// SPDX-License-Identifier: Apache-2.0
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

package resources

import (
	"time"

	"golang.org/x/oauth2"
)

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type AuthProvider struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type UserToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"expiresAt"`
}

func (ut *UserToken) JWTToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  ut.AccessToken,
		RefreshToken: ut.RefreshToken,
		Expiry:       time.Unix(ut.ExpiresAt, 0),
	}
}

type Company struct {
	ID         string     `json:"_id"` //nolint:tagliatelle
	Name       string     `json:"name"`
	TenantID   string     `json:"tenantId"`
	Pipelines  Pipelines  `json:"pipelines"`
	Repository Repository `json:"repository"`
}

type Pipelines struct {
	Type string `json:"type"`
}

type Repository struct {
	Type string `json:"type"`
}

type Project struct {
	ID                   string        `json:"_id"` //nolint:tagliatelle
	Name                 string        `json:"name"`
	ConfigurationGitPath string        `json:"configurationGitPath"`
	Environments         []Environment `json:"environments"`
	ProjectID            string        `json:"projectId"`
	Pipelines            Pipelines     `json:"pipelines"`
	TenantID             string        `json:"tenantId"`
}

type Environment struct {
	DisplayName string  `json:"label"` //nolint:tagliatelle
	EnvID       string  `json:"value"` //nolint:tagliatelle
	Cluster     Cluster `json:"cluster"`
}

type Cluster struct {
	Hostname  string `json:"hostname"`
	Namespace string `json:"namespace"`
}

type DeployProject struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type PipelineStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}