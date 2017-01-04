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

// discoCmd represents the disco command
var discoCmd = &cobra.Command{
	Use:   "disco",
	Short: "returns the list of available albums of the specified artist as an M3U playlist",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		url := cmd.Flag("addr").Value.String() + "/get_discography_artist"

		// first get the id of the artist
		id, err := getArtistID(strings.Join(args, " "), cmd.Flag("addr").Value.String())
		if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}

		// now get the discography for this id
		disco, err := get(url,
			kv{"id", id},
		)
		if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}

		fmt.Print(disco)
	},
}

func init() {
	RootCmd.AddCommand(discoCmd)
}
