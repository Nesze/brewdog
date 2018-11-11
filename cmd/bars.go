// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"sort"

	"github.com/Nesze/brewdog/brewery"
	"github.com/Nesze/brewdog/model"
	"github.com/spf13/cobra"
)

// barsCmd represents the bars command
var barsCmd = &cobra.Command{
	Use:   "bars",
	Short: "List bars",
	Long:  `List all brewdog bars.`,
	Run: func(cmd *cobra.Command, args []string) {
		b := brewery.Checkout()

		bars := b.Bars()
		sort.Sort(model.Bars(bars))
		for _, bar := range bars {
			fmt.Println(bar.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(barsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// barsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// barsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
