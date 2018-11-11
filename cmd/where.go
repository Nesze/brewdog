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
	"strings"

	"github.com/Nesze/brewdog/brewery"
	"github.com/spf13/cobra"
)

var beer string

// whereCmd represents the where command
var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "List of bars where the input beer is available",
	Long:  `List of bars where the input beer is available`,
	Run: func(cmd *cobra.Command, args []string) {
		b := brewery.Checkout()
		bars := b.WhereIsOnTap(strings.ToLower(beer))
		if len(bars) == 0 {
			fmt.Printf("%s: not found\n", beer)
			return
		}

		fmt.Printf("%s:\n", beer)
		for _, bar := range bars {
			fmt.Printf("\t%s:\n", bar.Name)
			for _, beer := range bar.OnTap {
				fmt.Printf("\t\t%s | %s (%s %s)\n", beer.Name, beer.Style, beer.Brewery, beer.ABV)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(whereCmd)

	whereCmd.Flags().StringVarP(&beer, "beer", "b", "IPA", "Partial beer name to filter on")
}
