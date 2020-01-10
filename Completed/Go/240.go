// dfs O(log(row * col))
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	r, c := len(matrix)-1, 0
	for {
		if !isInbounds(matrix, r, c) {
			return false
		}

		if val := matrix[r][c]; val == target {
			return true
		} else if val < target {
			c += 1
		} else if val > target {
			r -= 1
		} else {
			return false
		}
	}
}

func isInbounds(matrix [][]int, r int, c int) bool {
	return r >= 0 && c < len(matrix[r])
}

// // simple O(row (log col))
// func searchMatrix(matrix [][]int, target int) bool {
//   if len(matrix) == 0 || len(matrix[0]) == 0 {
//     return false
//   }

//   for _, arr := range matrix {
//     if contains(arr , target) {
//       return true
//     }
//   }

//   return false
// }

// func contains(arr []int, t int) bool {
//   lo, hi := 0, len(arr)-1
//   for lo <= hi {
//     mid := (lo + hi) / 2
//     if arr[mid] == t {
//       return true
//     } else if t < arr[mid] {
//       hi = mid -1
//     } else {
//       lo = mid +1
//     }
//   }
//   return false
// }