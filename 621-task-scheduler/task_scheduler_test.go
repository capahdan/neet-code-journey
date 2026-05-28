package taskscheduler

import "testing"

func TestTaskScheduler(t *testing.T) {
	tests := []struct {
		name  string
		tasks []byte
		n     int
		want  int
	}{
		{name: "test 1", tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 2, want: 8},
		{name: "test 2", tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, n: 2, want: 12},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := leastInterval(test.tasks, test.n)
			if got != test.want {
				t.Errorf("leastInterval(%v, %d) = %d, want %d", test.tasks, test.n, got, test.want)
			}
		})
	}
}
