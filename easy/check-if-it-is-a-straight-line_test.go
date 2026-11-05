package easy

import "testing"

func Test_checkStraightLine(t *testing.T) {
	cases := []struct {
		name string
		in   [][]int
		want bool
	}{
		{"官方示例1-直线", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, true},
		{"官方示例2-非直线", [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, false},
		{"只有两个点", [][]int{{0, 0}, {1, 1}}, true},
		{"竖直直线", [][]int{{0, 0}, {0, 5}, {0, -3}}, true},
		{"水平直线", [][]int{{-3, 2}, {0, 2}, {7, 2}}, true},
		{"非竖直不同x", [][]int{{0, 0}, {0, 5}, {1, 2}}, false},
		{"斜率2/3浮点误差用例", [][]int{{-5, -5}, {-2, -3}, {1, -1}}, true},
		{"斜率2/3反例", [][]int{{-5, -5}, {-2, -3}, {1, -2}}, false},
		{"大坐标", [][]int{{-10000, -10000}, {0, 0}, {10000, 10000}}, true},
		{"大坐标非直线", [][]int{{-10000, -10000}, {0, 0}, {10000, 9999}}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := checkStraightLine(c.in); got != c.want {
				t.Fatalf("checkStraightLine(%v) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}
