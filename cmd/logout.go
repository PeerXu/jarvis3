// Copyright © 2016 Peer Xu <pppeerxu@gmail.com>
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

package cmd

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
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

		signCli, err := SimpleNewSigningClientWithEnvironment(logger)
		if err != nil {
			fmt.Println(err)
			return
		}

		cfg := LoadConfigure()
		ctx := context.Background()
		err = signCli.Logout(ctx, cfg.Username)
		if err != nil {
			fmt.Println(err)
			return
		}

		ok, err := cmd.Flags().GetBool("env")
		if err == nil && ok {
			fmt.Printf(`
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
