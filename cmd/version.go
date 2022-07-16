// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version buildinfo item
	Version = "dev"
	// Commit buildinfo item
	Commit = "none"
	// Date buildinfo item
	Date = "unknown"
	// BuiltBy buildinfo item
	BuiltBy = "unknown"
	// HOME the home path
	HOME = ""
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the craft version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			fmt.Sprintf(
				`Current Craft Version %v Commit %v, Built @%v By %v.`,
				Version,
				Commit,
				Date,
				BuiltBy,
			),
		)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
