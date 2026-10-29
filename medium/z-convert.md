# 思路和伪代码

a  g
b fh
ce i
d  j

第 0 列，第 3 列，第 6 列，全
for i:=0; i < len(str) ;i++ {
	for col:=0;;col++ {
		if i == len(str) {
			return arr
		}
		if col % (totalRows - 1) == 0 {
			for row 0 -> totalRows - 1 {
				arr[row][col]=str[i]
				i++
			}
		} else {
			offset := col % (totalRows - 1)
			row_with_val=totalRows - 1-offset
			for row 0 -> totalRows - 1 {
				if col == row_with_val {
					arr[row][col]=str[i]
					i++
				} else {
					arr[row][col]=' '
				}
			}
		}
	}
}
