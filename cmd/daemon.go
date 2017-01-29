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
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/PeerXu/jarvis3/computing"
	computing_service "github.com/PeerXu/jarvis3/computing/service"
	"github.com/PeerXu/jarvis3/repository"
	"github.com/PeerXu/jarvis3/signing"
	signing_service "github.com/PeerXu/jarvis3/signing/service"
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

		var signingService signing_service.Service
		{
			signingService = signing.NewService(logger, users)
			signingService = signing.NewLoggingService(log.NewContext(logger).With("component", "signing"), signingService)
		}

		var computingService computing_service.Service
		{
			computingService = computing.NewService(logger, projects)
			computingService = computing.NewLoggingService(log.NewContext(logger).With("component", "computing"), computingService)
		}

		mux := http.NewServeMux()

		signCli, err := NewSigningClientForService(logger)
		if err != nil {
			cmd.Printf("create singing client error: %v", err)
			return
		}

		mux.Handle("/signing/v1/", signing.MakeHandler(ctx, signingService))
		mux.Handle("/computing/v1/", computing.MakeHandler(ctx, computingService, signCli))

		mux.Handle("/agent/v1/", http.StripPrefix("/agent/v1/", http.FileServer(http.Dir("dist/agent"))))
		mux.Handle("/example/v1/", http.StripPrefix("/example/v1/", http.FileServer(http.Dir("example/web"))))

		http.Handle("/", accessControl(mux))

		errs := make(chan error, 2)

		go func() {
			errs <- http.ListenAndServe(viper.GetString("development.server.host"), nil)
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
