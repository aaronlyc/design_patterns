package search_file

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSearchFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SearchFile Suite")
}
