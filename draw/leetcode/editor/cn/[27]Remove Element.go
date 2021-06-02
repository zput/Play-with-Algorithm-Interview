//Given an array nums and a value val, remove all instances of that value in-pla
//ce and return the new length. 
//
// Do not allocate extra space for another array, you must do this by modifying 
//the input array in-place with O(1) extra memory. 
//
// The order of elements can be changed. It doesn't matter what you leave beyond
// the new length. 
//
// Clarification: 
//
// Confused why the returned value is an integer but your answer is an array? 
//
// Note that the input array is passed in by reference, which means a modificati
//on to the input array will be known to the caller as well. 
//
// Internally you can think of this: 
//
// Constraints: 
//
// 
// 0 <= nums.length <= 100 
// 0 <= nums[i] <= 50 
// 0 <= val <= 100 
// 
// Related Topics 数组 双指针 


//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	i, j := 0, 0	
	length := len(nums)
	ret := 0
	for ; i<length; i++{
		if nums[i] == val{
			continue
		}	
		ret++
		for ; j<i; j++{
			if nums[j] == val{
				nums[j], nums[i] = nums[i], nums[j]
				j++
				break
			}			
		}	
	}	
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
