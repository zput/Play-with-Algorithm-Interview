//Given a sorted array nums, remove the duplicates in-place such that duplicates
// appeared at most twice and return the new length. 
//
// Do not allocate extra space for another array; you must do this by modifying 
//the input array in-place with O(1) extra memory. 
//
// Clarification: 
//
// Confused why the returned value is an integer, but your answer is an array? 
//
// Note that the input array is passed in by reference, which means a modificati
//on to the input array will be known to the caller. 
//
// Internally you can think of this: 
//
// 
//// nums is passed in by reference. (i.e., without making a copy)
//int len = removeDuplicates(nums);
//
//// any modification to nums in your function would be known by the caller.
//// using the length returned by your function, it prints the first len element
//s.
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//}
// 
//
// 
// Example 1: 
//
// 
//Input: nums = [1,1,1,2,2,3]
//Output: 5, nums = [1,1,2,2,3]
//Explanation: Your function should return length = 5, with the first five eleme
//nts of nums being 1, 1, 2, 2 and 3 respectively. It doesn't matter what you leav
//e beyond the returned length.
// 
//
// Example 2: 
//
// 
//Input: nums = [0,0,1,1,1,1,2,3,3]
//Output: 7, nums = [0,0,1,1,2,3,3]
//Explanation: Your function should return length = 7, with the first seven elem
//ents of nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively. It doesn't m
//atter what values are set beyond the returned length.
// 
//
// 
// Constraints: 
//
// 
// 1 <= nums.length <= 3 * 104 
// -104 <= nums[i] <= 104 
// nums is sorted in ascending order. 
// 
// Related Topics 数组 双指针 


//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	i, j := 0, 0
	length := len(nums)
	frontF, front := 105, 106
	for ; i<length; i++{
		if (nums[i] == front) && (nums[i] == frontF){
			continue
		}	
		frontF = front
		front = nums[i]
		nums[j] = nums[i]
		j++
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
