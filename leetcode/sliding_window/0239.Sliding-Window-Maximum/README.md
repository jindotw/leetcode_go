[239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)  
Hard

You are given an array of integers `nums`, there is a sliding window of size  `k`  which is moving from the very left of the array to the very right. You can only see the  `k`  numbers in the window. Each time the sliding window moves right by one position.

Return  _the max sliding window_.

**Example 1:**

**Input:** nums = [1,3,-1,-3,5,3,6,7], k = 3  
**Output:** [3,3,5,5,6,7]  
**Explanation:**  
<pre>Window position                Max  
---------------               -----  
[1  3  -1] -3  5  3  6  7      <b>3</b>  
1 [3  -1  -3] 5  3  6  7       <b>3</b>  
1  3 [-1  -3  5] 3  6  7       <b>5</b>  
1  3  -1 [-3  5  3] 6  7       <b>5</b>  
1  3  -1  -3 [5  3  6] 7       <b>6</b>  
1  3  -1  -3  5 [3  6  7]      <b>7</b></pre>  

**Example 2:**

**Input:** nums = [1], k = 1  
**Output:** [1]

**Constraints:**

-   `1 <= nums.length <= 105`
-   `-104 <= nums[i] <= 104`
-   `1 <= k <= nums.length`