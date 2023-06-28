package math

import "testing"

type AddData struct {
	x, y int
	res  int
}

func TestAdd(t *testing.T) {

	testData := []AddData{
		{1, 2, 3},
		{2, 5, 7},
	}

	for _, data := range testData {
		result := Add(data.x, data.y)

		if result != data.res {
			t.Errorf("Fail")
		}
	}
}
