package main

import "fmt"
import "testing"

type testcase struct {
	runner runner
	distance int
	expectedTime int
}

func TestRunners(t *testing.T) {
	for _, tc := range []testcase{
		testcase{runner: runbot9000{}, distance: 200, expectedTime: 200},
		testcase{runner: runbot9000{}, distance: 4, expectedTime: 4},
		testcase{runner: baby{}, distance: 4, expectedTime: 40},
		testcase{runner: baby{}, distance: 19, expectedTime: 190},
		testcase{runner: developer{clumsiness: 0.0}, distance: 100, expectedTime: 300},
		testcase{runner: developer{clumsiness: 0.5}, distance: 50, expectedTime: 75},
		testcase{runner: developer{clumsiness: 0.9}, distance: 200, expectedTime: 59},
	} {
		name := fmt.Sprintf("%s runs %d", tc.runner.name(), tc.distance)

		t.Run(name, func(t *testing.T) {
			if actualTime := tc.runner.run(tc.distance); actualTime != tc.expectedTime {
				t.Errorf("expected %d, but got %d", tc.expectedTime, actualTime)
			}
		})
	}
}

func BenchmarkDeveloper(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("Developer@%d", n), func (b *testing.B) {
			for i := 0; i < b.N; i++ {
				developer{clumsiness: 0.5}.run(10)
			}
		})
	}
}
