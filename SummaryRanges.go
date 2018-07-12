// Given a sorted integer array without duplicates, return the summary of its ranges.

// Example 1:
// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
// Explanation: 0,1,2 form a continuous range; 4,5 form a continuous range.

func summaryRanges(nums []int) []string {
	arr := []string{}
    //if array is empty
	if len(nums) == 0 {
		return arr
	}
	s := nums[0]
	str := ""
    
	for i := 0; i < len(nums); i++ {
        //if i is the last index OR the next element is not consecutive
		if i == len(nums)-1 || nums[i] != nums[i+1]-1 {
            //for single elements
			if nums[i] == s {
				str = fmt.Sprintf("%d", s)
			} else {
                //for consecutive sets
				str = fmt.Sprintf("%d->%d", s, nums[i])
			}
			if i+1 < len(nums) {
				s = nums[i+1]
			}
			arr = append(arr, str)
		}
	}
	return arr
}
