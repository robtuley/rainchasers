package gauge

import (
	"testing"
)

func TestIdGeneration(t *testing.T) {

	s1 := Station{
		DataURL: "http://environment.data.gov.uk/flood-monitoring/id/measures/1029TH-level-downstage-i-15_min-mASD",
	}
	s2 := Station{
		DataURL: "http://environment.data.gov.uk/flood-monitoring/id/measures/somethingdifferent",
	}

	if s1.UUID() == s2.UUID() {
		t.Error("UUID clash", s1.UUID(), s2.UUID())
	}
}
