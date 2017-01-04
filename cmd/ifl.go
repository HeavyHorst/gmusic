// Copyright © 2017 Rene Kaufmann <kaufmann.r@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// iflCmd represents the ifl command
var iflCmd = &cobra.Command{
	Use:   "ifl",
	Short: "get the tracks associated to the automatic “I’m feeling lucky” station",
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flag("addr").Value.String() + "/get_ifl_station"

		ifl, err := get(url,
			kv{"num_tracks", cmd.Flag("limit").Value.String()},
		)
		if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}

		fmt.Print(ifl)
	},
}

func init() {
	iflCmd.Flags().String("limit", "20", "the number of songs to return")
	RootCmd.AddCommand(iflCmd)
}
