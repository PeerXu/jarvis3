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

	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var Username string
var Password string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to jarvis3 server",
	Run: func(cmd *cobra.Command, args []string) {
		var logger log.Logger
		{
			logger = log.NewLogfmtLogger(os.Stderr)
			logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		}

		signCli, err := SimpleNewSigningClient(logger)
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.Background()
		token, err := signCli.Login(ctx, Username, Password)
		if err != nil {
			fmt.Println(err)
			return
		}

		ok, err := cmd.Flags().GetBool("env")
		if err == nil && ok {
			fmt.Printf(`
export JVS_USERNAME=%v
export JVS_ACCESS_TOKEN=%v
`, Username, token.Token)
		} else {
			fmt.Printf("access tokenk: %v\n", token.Token)
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&Username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&Password, "password", "w", "", "Password")
	loginCmd.Flags().BoolP("env", "e", false, "set up environment for jarvis3 client")
}
