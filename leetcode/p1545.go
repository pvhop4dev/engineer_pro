package leetcode

//TODO

// 0(n) time complexity, 0(n) space complexity
func FindKthBit(n int, k int) byte {
	return FindS(n)[k-1]
}

func FindS(n int) []byte {
	if n == 1 {
		return []byte{0}
	}
	prev := FindS(n - 1)
	result := make([]byte, 2*len(prev)+1)
	copy(result, prev)
	result[len(prev)] = 1
	for i := 0; i < len(prev); i++ {
		result[len(prev)+1+i] = prev[len(prev)-1-i] ^ 1
	}
	return result
}

// O(log(n)) time complexity, O(1) space complexity
func FindKthBit2(n int, k int) byte {
	if n == 1 {
		return 0
	}
	mid := 1 << (n - 1)
	if k == mid {
		return 1
	}
	if k < mid {
		return FindKthBit2(n-1, k)
	}
	return FindKthBit2(n-1, 2*mid-k) ^ 1
}

func FindKthBit3(n int, k int) byte {
	return compute(n, k, false)
}

func compute(n int, k int, inverse bool) byte {
	if n == 1 {
		if inverse {
			return 1
		} else {
			return 0
		}
	}
	len := (1 << n) - 1
	mid := len/2 + 1
	if k == mid {
		if inverse {
			return 1
		} else {
			return 0
		}
	}
	if k < mid {
		return compute(n-1, k, inverse)
	}
	k -= mid
	k = mid - k
	return compute(n-1, k, !inverse)
}
