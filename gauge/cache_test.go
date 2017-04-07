package gauge

import (
	"reflect"
	"testing"
	"time"
)

func TestConcatDeDuplicatesAndOrdersByTime(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")
	r3 := Reading{
		DateTime: timestamp,
		Value:    3.21,
	}
	r2 := Reading{
		DateTime: timestamp.Add(time.Second),
		Value:    2.13,
	}
	r1 := Reading{
		DateTime: timestamp.Add(time.Second * 10),
		Value:    1.23,
	}

	var result []Reading

	result = concat([]Reading{r1, r3}, []Reading{r2, r1})
	if !reflect.DeepEqual(result, []Reading{r1, r2, r3}) {
		t.Error("concat() unexpected result", result)
	}

	result = concat([]Reading{r1, r2, r3}, []Reading{r2, r3, r1})
	if !reflect.DeepEqual(result, []Reading{r1, r2, r3}) {
		t.Error("concat() unexpected result", result)
	}

	result = concat([]Reading{}, []Reading{r3, r1})
	if !reflect.DeepEqual(result, []Reading{r1, r3}) {
		t.Error("concat() unexpected result", result)
	}

	result = concat([]Reading{r3, r2, r1}, []Reading{})
	if !reflect.DeepEqual(result, []Reading{r1, r2, r3}) {
		t.Error("concat() unexpected result", result)
	}
}
