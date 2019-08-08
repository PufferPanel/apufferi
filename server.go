/*
 Copyright 2019 Padduck, LLC
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
  	http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package apufferi

type Server struct {
	Variables      map[string]Variable `json:"data,omitempty"`
	Display        string              `json:"display,omitempty"`
	Environment    TypeWithMetadata    `json:"environment,omitempty"`
	Installation   []TypeWithMetadata  `json:"install,omitempty"`
	Uninstallation []TypeWithMetadata  `json:"uninstall,omitempty"`
	Type           string              `json:"type,omitempty"`
	Identifier     string              `json:"id,omitempty"`
	Execution      Execution           `json:"run,omitempty"`
}

type Variable struct {
	Description  string      `json:"desc,omitempty"`
	Display      string      `json:"display,omitempty"`
	Internal     bool        `json:"internal,omitempty"`
	Required     bool        `json:"required,omitempty"`
	Value        interface{} `json:"value,omitempty"`
	UserEditable bool        `json:"userEdit,omitempty"`
	Type         string      `json:"type,omitempty"`
	Options      []string    `json:"options,omitempty"`
}

type Execution struct {
	Arguments               []string           `json:"arguments,omitempty"`
	ProgramName             string             `json:"program,omitempty"`
	StopCommand             string             `json:"stop,omitempty"`
	Enabled                 bool               `json:"enabled,omitempty"`
	AutoStart               bool               `json:"autostart,omitempty"`
	AutoRestartFromCrash    bool               `json:"autorecover,omitempty"`
	AutoRestartFromGraceful bool               `json:"autorestart,omitempty"`
	PreExecution            []TypeWithMetadata `json:"pre,omitempty"`
	PostExecution           []TypeWithMetadata `json:"post,omitempty"`
	StopCode                int                `json:"stopCode,omitempty"`
	EnvironmentVariables    map[string]string  `json:"environmentVars,omitempty"`
}

type TypeWithMetadata struct {
	Type     string                 `json:"type,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Template struct {
	Server
	SupportedEnvironments []TypeWithMetadata `json:"supportedEnvironments,omitEmpty"`
}