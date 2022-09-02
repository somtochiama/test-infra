/*
Copyright 2022 The Flux authors

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

package tftestenv

import (
	"flag"
	"fmt"
)

// Options contains options for creating the terraform test environment
type Options struct {
	Provider string
	Retain   bool
	Existing bool
	Verbose  bool
}

var supportedProviders = []string{"aws", "azure", "gcp"}

func (o *Options) Bindflags(fs *flag.FlagSet) {
	fs.StringVar(&o.Provider, "provider", "", fmt.Sprintf("name of the provider %v", supportedProviders))

	// retain flag to prevent destroy and retaining the created infrastructure.
	fs.BoolVar(&o.Retain, "retain", false, "retain the infrastructure for debugging purposes")

	// existing flag to use existing infrastructure terraform state.
	fs.BoolVar(&o.Existing, "existing", false, "use existing infrastructure state for debugging purposes")

	// verbose flag to enable output of terraform execution.
	fs.BoolVar(&o.Verbose, "verbose", false, "verbose output of the environment setup")
}

func (o *Options) Validate() error {
	if o.Provider == "" {
		return fmt.Errorf("-provider flag must be set to one of %v", supportedProviders)
	}

	var supported bool

	for _, p := range supportedProviders {
		if p == o.Provider {
			supported = true
		}
	}
	if !supported {
		return fmt.Errorf("unsupported provider %q, must be one of %v", o.Provider, supportedProviders)
	}

	return nil
}
