// Time:     O(n)
// Space:    O(n)
func twoSum(nums []int, target int) []int {
    //Solving by brute force. 
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[j] == target - nums[i] {
                    return []int{j, i}
            }
        }
    }
    return nil
}
