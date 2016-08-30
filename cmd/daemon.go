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
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/PeerXu/jarvis3/computing"
	"github.com/PeerXu/jarvis3/repository"
	"github.com/PeerXu/jarvis3/signing"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "start a Jarvis3 server",
	Run: func(cmd *cobra.Command, args []string) {
		var logger log.Logger
		{
			logger = log.NewLogfmtLogger(os.Stderr)
			logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		}

		var (
			ctx = context.Background()
		)

		var (
			users    = repository.NewUserRepository()
			projects = repository.NewProjectRepository()
		)

		var signingService signing.Service
		{
			signingService = signing.NewService(logger, users)
			signingService = signing.NewLoggingService(log.NewContext(logger).With("component", "signing"), signingService)
		}

		var computingService computing.Service
		{
			computingService = computing.NewService(logger, projects)
			computingService = computing.NewLoggingService(log.NewContext(logger).With("component", "computing"), computingService)
		}

		mux := http.NewServeMux()

		signCli, err := SimpleNewSigningClientWithEnvironment(logger)
		if err != nil {
			cmd.Printf("create singing client error: %v", err)
			return
		}

		mux.Handle("/signing/v1/", signing.MakeHandler(ctx, signingService))
		mux.Handle("/computing/v1/", computing.MakeHandler(ctx, computingService, signCli))

		http.Handle("/", accessControl(mux))

		errs := make(chan error, 2)

		go func() {
			errs <- http.ListenAndServe(":27182", nil)
		}()
		go func() {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT)
			errs <- fmt.Errorf("%s", <-c)
		}()

		<-errs
	},
}

func init() {
	RootCmd.AddCommand(daemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
