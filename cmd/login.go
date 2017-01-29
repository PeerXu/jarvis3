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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to jarvis3 server",
	Run: func(cmd *cobra.Command, args []string) {
		logger := NewClientLogger()
		signCli, err := NewSigningClientForClientWithoutEnvironment(logger)
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.Background()
		u, err := signCli.Login(ctx, Username, Password)
		if err != nil {
			fmt.Println(err)
			return
		}
		token := u.AccessTokens[0]

		ok, err := cmd.Flags().GetBool("env")
		if err == nil && ok {
			fmt.Printf(`
export JVS_USER_ID="%v"
export JVS_USERNAME="%v"
export JVS_ACCESS_TOKEN="%v"
`, u.ID.String(), Username, token.Token)
		} else {
			fmt.Printf(`login successed!
user id: %v
username: %v
access token: %v`, u.ID.String(), Username, token.Token)
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&Username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&Password, "password", "w", "", "Password")
	loginCmd.Flags().BoolP("env", "e", false, "set up environment for jarvis3 client")
}
