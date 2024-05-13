package unit_tests

import (
	sportTests "sports-competition/app/modules/sport/unit_tests"
	"testing"
)

func TestModules(t *testing.T) {
	sportTests.TestSports(t)
}
