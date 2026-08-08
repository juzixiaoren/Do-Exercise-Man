#include<iostream>
#include<vector>
#include<string>
using namespace std;
class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        int left = 0;
        int right = numbers.size() - 1;
        while (left < right) {
            if (numbers[right] + numbers[left] > target) {
                right--;
            }
            else if (numbers[right] + numbers[left] < target) {
                left++;
            }
            else {
                return vector<int>{left+1, right+1};
            }
        }
        return vector<int>{-1, -1};
    }
};
int main() {
    Solution s;
    vector<int> numbers = {2,7,11,15};
    int target = 9;
    vector<int> result = s.twoSum(numbers, target);
    for (int i : result) {
		cout << i << " ";
	}
}