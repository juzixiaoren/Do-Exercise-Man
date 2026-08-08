#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int left = 0;
        int sum=0;
        int ans = nums.size();
        for(int right=0;right<nums.size();right++){
            sum += nums[right];
            if (sum < target) {
                continue;
            }
            if (sum >= target) {
                while (sum - nums[left] >= target) {
					sum -= nums[left];
					left++;
				}
                ans = min(right - left + 1, ans);
            }
        }
        if(sum<target)
        return 0;
        return ans;
    }
};