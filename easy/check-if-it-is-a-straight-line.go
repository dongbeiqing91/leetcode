package easy

// https://leetcode.cn/problems/check-if-it-is-a-straight-line/

// checkStraightLine 判断坐标点是否全部在同一条直线上。
// 用整数叉积 (x-x0)*dy == (y-y0)*dx 替代斜率浮点比较，
// 避免除零和 float64 精度误差，同时统一处理竖直/水平/一般斜率。
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}
	p1 := coordinates[0]
	p2 := coordinates[1]
	dx := p2[0] - p1[0]
	dy := p2[1] - p1[1]
	for i := 2; i < len(coordinates); i++ {
		p := coordinates[i]
		deltaX := p[0] - p1[0]
		deltaY := p[1] - p1[1]
		if deltaX * dy != deltaY * dx {
			return false
		}
	}
	return true
}
