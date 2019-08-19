/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package update

import (
	"github.com/spf13/cobra"

	"github.com/emicklei/go-restful/log"
	"sigs.k8s.io/apiserver-builder-alpha/cmd/apiserver-boot/boot/init_repo"
)

var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Update the vendor packages managed by apiserver-builder.",
	Long:  `Update the vendor packages managed by apiserver-builder.`,
	Example: `# Replace the vendor packages managed by apiserver-builder with versions for the current install.
apiserver-boot update vendor
`,
	Run: RunUpdateVendor,
}

func AddUpdateVendorCmd(cmd *cobra.Command) {
	cmd.AddCommand(vendorCmd)
}

func RunUpdateVendor(cmd *cobra.Command, args []string) {
	init_repo.Update = true
	log.Printf("Replacing vendored libraries managed by apiserver-builder with the current version.")
	init_repo.RunVendorInstall(cmd, args)
}
