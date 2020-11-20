package sports

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSports(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sports Suite")
}
