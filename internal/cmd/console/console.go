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

package console

import (
	"github.com/mia-platform/miactl/internal/cmd/console/deploy"
	"github.com/spf13/cobra"
)

func NewConsoleCmd() *cobra.Command {
	// Note: console should act as a resource that receives commands to be executed
	cmd := &cobra.Command{
		Use:   "console",
		Short: "select console resource",
	}

	cmd.AddCommand(deploy.NewDeployCmd())

	return cmd
}