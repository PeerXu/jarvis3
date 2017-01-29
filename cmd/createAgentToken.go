// Copyright © 2017 Peer Xu <pppeerxu@gmail.com>
//
// This file is part of test.
//
// test is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// test is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with test. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/spf13/cobra"
)

// createAgentTokenCmd represents the createAgentToken command
var createAgentTokenCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a agent token for agent",
	Run: func(cmd *cobra.Command, args []string) {
		logger := NewClientLogger()
		signCli, err := NewSigningClientForClient(logger)

		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.Background()
		token, err := signCli.CreateAgentToken(ctx, AgentTokenName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("agent token: %v\n", token.Token)
	},
}

func init() {
	agentTokenCmd.AddCommand(createAgentTokenCmd)

	createAgentTokenCmd.Flags().StringVarP(&AgentTokenName, "name", "n", "", "Agent token name")
}
