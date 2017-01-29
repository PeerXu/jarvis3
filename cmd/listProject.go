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

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// listProjectCmd represents the listProject command
var listProjectCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		logger := NewClientLogger()
		compCli, err := NewComputingClientForClient(logger)

		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.Background()
		projs, err := compCli.ListProjects(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, p := range projs {
			fmt.Printf("%v\t%v\n", p.ID, p.Name)
		}
	},
}

func init() {
	projectCmd.AddCommand(listProjectCmd)
}
