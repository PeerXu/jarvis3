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
	"os"

	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/PeerXu/jarvis3/user"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from jarvis3 server",
	Run: func(cmd *cobra.Command, args []string) {
		var logger log.Logger
		{
			logger = log.NewLogfmtLogger(os.Stderr)
			logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		}

		signCli, err := NewSigningClientForClient(logger)
		if err != nil {
			fmt.Println(err)
			return
		}

		cfg := LoadConfigure()
		ctx := context.Background()
		err = signCli.Logout(ctx, user.UserID(cfg.UserID))
		if err != nil {
			fmt.Println(err)
			return
		}

		ok, err := cmd.Flags().GetBool("env")
		if err == nil && ok {
			fmt.Printf(`
unset JVS_USER_ID
unset JVS_USERNAME
unset JVS_ACCESS_TOKEN
`)
		} else {
			fmt.Printf("%v logout", cfg.Username)
		}
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)

	logoutCmd.Flags().BoolP("env", "e", false, "set up environment for jarvis3 client")
}
