public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        List<int> solution = new List<int>();
        
        for(int n = 0; n < nums.Length; n++){
            for(int m = 0; m < nums.Length; m++){
                if(n != m && nums[n] + nums[m] == target) {
                    solution.Add(n);
                    solution.Add(m);
                    return solution.ToArray();
                } 
            }
        }
        
        return solution.ToArray();
        
    }
}