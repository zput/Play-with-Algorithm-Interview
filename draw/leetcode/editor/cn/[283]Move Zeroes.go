//Given an integer array nums, move all 0's to the end of it while maintaining t
//he relative order of the non-zero elements. 
//
//
//
// Note that you must do this in-place without making a copy of the array. 
//
//
// 
// Example 1: 
// Input: nums = [0,1,0,3,12]
//Output: [1,3,12,0,0]
// Example 2: 
// Input: nums = [0]
//Output: [0]
// 
// 
// 
// Constraints: 
//
//
// 
// 1 <= nums.length <= 104 
// -231 <= nums[i] <= 231 - 1 
// 
//
// 
//Follow up: Could you minimize the total number of operations done? Related Top
//ics 数组 双指针 


//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int)  {
	i, j := 0, 0	
	length := len(nums)
	for ;j<length; j++{
		if nums[j] == 0{
			continue
		}
		for ;i<j; i++{
				if nums[i] == 0{
					nums[i], nums[j]	= nums[j], nums[i]
					i++
					break
				}			
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
