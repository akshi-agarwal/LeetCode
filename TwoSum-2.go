// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        hash[nums[i]] = i
    }
    
    for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
        if _, ok := hash[complement]; ok {
            if ok == true && i != hash[complement] {				
                return []int{i, hash[complement]}
            }
        }
    }
    return nil
}

// Time:  O(n)
// Space: O(1) 
