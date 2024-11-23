package leetcode

func NumOfSubarrays(arr []int, k int, threshold int) int {
	result := 0
	begin := 0
	for begin <= len(arr)-k {
		end := begin + k
		sum := 0
		for i := begin; i < end; i++ {
			sum += arr[i]
		}
		if sum/k >= threshold {
			result++
		}
		begin++
	}
	return result
}

func NumOfSubarrays2(arr []int, k int, threshold int) int {
	result := 0
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if sum/k >= threshold {
		result++
	}
	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		if sum/k >= threshold {
			result++
		}
	}
	return result
}

func NumOfSubarrays3(arr []int, k int, threshold int) int {
	n := len(arr)
	begin := 0
	end := n-1
	sum := 0
	res := 0
	for begin <= n-k {
		for end < n-1 && end-begin+1 < k {
			end++
			sum += arr[end]
		}
		if sum >= threshold *k {
			res++
		}
		sum -= arr[begin]
		begin++
	}
	return res
}