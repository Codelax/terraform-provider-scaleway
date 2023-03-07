package scaleway

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func randomWithPrefix(name string) string {
	return acctest.RandomWithPrefix(fmt.Sprintf("tf-acc-%s", name))
}

// Generated name has format: "tf-{prefix}-{generated}
func extractGeneratedNamePrefix(name string) string {
	// {prefix}-{generated}
	name = strings.TrimPrefix(name, "tf-")

	// {prefix}-{generated}
	//         ^
	dashIndex := strings.Index(name, "-")

	if dashIndex == -1 {
		// some are only tf-{name}
		return name
	}

	// {prefix}
	return name[:dashIndex]
}

// compareJSONFieldsStrings compare two strings from request JSON bodies
// has special case when string are terraform generated names
func compareJSONFieldsStrings(expected, actual string) bool {
	if strings.HasPrefix(expected, "tf-") {
		expectedPrefix := extractGeneratedNamePrefix(expected)
		actualPrefix := extractGeneratedNamePrefix(actual)
		return expectedPrefix == actualPrefix
	}

	return expected == actual
}
