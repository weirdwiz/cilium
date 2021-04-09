// Copyright 2021 Authors of Cilium
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
	"strconv"

	"github.com/spf13/cobra"
)

// recorderDeleteCmd represents the recorder_delete command
var recorderDeleteCmd = &cobra.Command{
	Use:    "delete <recorder id>",
	Short:  "Delete individual pcap recorder",
	PreRun: requireRecorderID,
	Run: func(cmd *cobra.Command, args []string) {
		recIDstr := args[0]
		id, err := strconv.ParseInt(recIDstr, 0, 64)
		if err != nil {
			Fatalf("Unable to parse recorder ID: %s", recIDstr)
		}

		err = client.DeleteRecorderID(id)
		if err != nil {
			Fatalf("Cannot delete recorder '%v': %s\n", id, err)
		}
	},
}

func init() {
	recorderCmd.AddCommand(recorderDeleteCmd)
}
