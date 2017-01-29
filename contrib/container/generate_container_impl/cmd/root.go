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
	"io/ioutil"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var inputDir string
var output string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "generate_container_impl",
	Short: "generate container_impl.go",
	Run: func(cmd *cobra.Command, args []string) {
		container_impl_file := inputDir + "/container_impl.go.template"
		container_main_file := inputDir + "/main.go.template"

		impl_buf, err := ioutil.ReadFile(container_impl_file)
		if err != nil {
			cmd.Printf("failed to open %v: %v\n", container_impl_file, err)
			return
		}

		main_buf, err := ioutil.ReadFile(container_main_file)
		if err != nil {
			cmd.Printf("failed to open %v: %v\n", container_main_file, err)
			return
		}

		output_file, err := os.Create(output)
		if err != nil {
			cmd.Printf("failed to write file %v: %v\n", output, err)
			return
		}
		// output_writer := bufio.NewWriter(output_file)

		impl_tmpl := template.Must(template.New("container_impl").Parse(string(impl_buf)))
		impl_tmpl.Execute(output_file, struct{ Code string }{string(main_buf)})

		cmd.Printf("Generated...\n")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().StringVarP(&inputDir, "dir", "d", "container", "Container template directory")
	RootCmd.Flags().StringVarP(&output, "output", "o", "container/container_impl.go", "Ouput container_impl.go path")
}
