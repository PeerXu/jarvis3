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

// listExecutorsCmd represents the listExecutors command
var listExecutorsCmd = &cobra.Command{
	Use:   "list",
	Short: "List executors",
	Run: func(cmd *cobra.Command, args []string) {
		logger := NewClientLogger()
		compCli, err := NewComputingClientForClient(logger)

		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.Background()
		execs, err := compCli.ListExecutors(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}

		// TODO(Peer): print formated items
		for _, e := range execs {
			fmt.Printf("%v\t%v\t%v\n", e.ID, e.Name, e.Pack)
		}
	},
}

func init() {
	executorCmd.AddCommand(listExecutorsCmd)
}
