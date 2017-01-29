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

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Run: func(cmd *cobra.Command, args []string) {
		logger := NewClientLogger()
		compCli, err := NewComputingClientForClient(logger)

		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.Background()
		proj, err := compCli.CreateProject(ctx, ProjectName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Create project#%v successed!\n", proj.ID)
	},
}

func init() {
	projectCmd.AddCommand(createProjectCmd)

	createProjectCmd.Flags().StringVarP(&ProjectName, "name", "n", "", "Project Name")
}
