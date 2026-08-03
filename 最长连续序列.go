package main

func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
	}
	max := 0
	for num := range set {
		_, ok := set[num-1]
		if ok {
			continue
		} else {
			head := num
			count := 1
			for true {
				_, ok := set[head+1]
				if ok {
					count++
					head = head + 1
				} else {
					break
				}
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}

/*
遍历一个 map 的方法是 for num := range set { ... }
*/
