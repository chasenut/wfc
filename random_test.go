package wfc

import "testing"

func TestRandomIndex(t *testing.T) {
	cases := []struct{
		arr			[]float64
		r			float64
		expected	int
	} {
		{
			arr: 		[]float64{1.0, 2.0, 3.0},
			r: 			0.1,
			expected: 	0,
		},
		{
			arr: 		[]float64{1.0, 3.0},
			r: 			0.5,
			expected: 	1,
		}, 
		{
			arr: 		[]float64{0.0, 0.0, 3.0, 1.0},
			r: 			0.5,
			expected: 	2,
		}, 
		{
			arr: 		[]float64{},
			r: 			0.5,
			expected: 	0,
		}, 
		{
			arr:		[]float64{10.0, 10.0, 20.0, 50.0, 10.0},
			r: 			0.9,
			expected: 	3,
		},
	}

	for _, c := range cases {
		actual := randomIndex(c.arr, c.r)
		if actual != c.expected {
			t.Errorf("randomIndex() = %v, want %v", actual, c.expected)
		}
	}
}
