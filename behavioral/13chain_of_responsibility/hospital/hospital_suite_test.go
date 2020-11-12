package hospital_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHospital(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hospital Suite")
}
