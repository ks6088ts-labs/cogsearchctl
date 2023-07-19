/*
Copyright © 2023 ks6088ts

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package indexer

import (
	"fmt"
	"log"

	"github.com/ks6088ts-labs/cogsearchctl/internal"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an indexer for Azure Cognitive Search",
	Long:  `Run an indexer for Azure Cognitive Search.`,
	Run: func(cmd *cobra.Command, args []string) {
		searchServiceName, err := cmd.Flags().GetString("searchServiceName")
		if err != nil {
			log.Fatalf("failed to parse `searchServiceName`: %v", err)
		}
		indexerName, err := cmd.Flags().GetString("indexerName")
		if err != nil {
			log.Fatalf("failed to parse `indexerName`: %v", err)
		}
		searchApiKey, err := cmd.Flags().GetString("searchApiKey")
		if err != nil {
			log.Fatalf("failed to parse `searchApiKey`: %v", err)
		}
		bodyFilePath, err := cmd.Flags().GetString("bodyFilePath")
		if err != nil {
			log.Fatalf("failed to parse `bodyFilePath`: %v", err)
		}
		response, err := internal.HttpRequest(fmt.Sprintf("https://%s.search.windows.net/indexers/%s/run?api-version=2020-06-30", searchServiceName, indexerName), bodyFilePath, "POST", searchApiKey)
		if err != nil {
			log.Fatalf("failed to run indexer. err=%v, response=%v", err, response)
		}
		defer response.Body.Close()
		log.Printf("status code=%v", response.StatusCode)
	},
}

func init() {
	indexerCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("searchServiceName", "s", "searchServiceName", "search service name")
	runCmd.Flags().StringP("indexerName", "i", "indexerName", "indexer name")
	runCmd.Flags().StringP("searchApiKey", "k", "searchApiKey", "search api key")
	runCmd.Flags().StringP("bodyFilePath", "f", "body.json", "body file path")
}
