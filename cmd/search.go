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

	"strings"

	"github.com/spf13/cobra"
)

var (
	artist  string
	sType   string
	limit   string
	station bool
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for music",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		var url string
		if station {
			url = cmd.Flag("addr").Value.String() + "/get_new_station_by_search"
		} else {
			url = cmd.Flag("addr").Value.String() + "/get_by_search"
		}

		result, err := get(url,
			kv{"artist", artist},
			kv{"title", strings.Join(args, " ")},
			kv{"exact", "false"},
			kv{"type", sType},
			kv{"num_tracks", limit},
		)
		if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}

		fmt.Print(result)
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&artist, "artist", "", "a string to search in the name of the artist in any kind of search")
	searchCmd.Flags().StringVar(&sType, "type", "matches", "search for artist, album or song")
	searchCmd.Flags().StringVar(&limit, "limit", "20", "the number of songs to return")
	searchCmd.Flags().BoolVar(&station, "station", false, "create a radio station from the search result")
}
