
#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>
using namespace std;
class Solution
{
public:
    vector<int> twoSum(vector<int> &numbers, int target)
    {
        int left = 0;
        int right = numbers.size() - 1;
        while (left < right)
        {
            if (numbers[right] + numbers[left] > target)
            {
                right--;
            }
            else if (numbers[right] + numbers[left] < target)
            {
                left++;
            }
            else
            {
                return vector<int>{left + 1, right + 1};
            }
        }
        return vector<int>{-1, -1};
    }
};
// 这里的做法是，用双指针，一个指向头，一个指向尾，如果和大于目标值，尾指针向前移动，如果和小于目标值，头指针向后移动，如果等于目标值，返回头指针和尾指针的下标
// 这里假设数组已经排序好，实际不一定排序好，更优秀的做法是，一次遍历，用一个哈希表存储遇到的值，每次遇到时计算target-i，如果在哈希表中，则返回i和target-i的下标，如果不在哈希表中，则将i存储在哈希表中

class Solution_2
{
public:
    vector<int> twoSum(vector<int> &numbers, int target)
    {
        unordered_map<int, int> map;
        vector<vector<int>> ans = {};
        for (int i = 0; i < numbers.size(); i++)
        {
            int y = target - numbers[i];
            if (map.find(y) != map.end())
            {
                return {map[y] + 1, i + 1};
            }
            map[numbers[i]] = i;
        }
        return {-1, -1};
    }
};

int main()
{
    Solution s;
    vector<int> numbers = {2, 7, 11, 15};
    int target = 9;
    vector<int> result = s.twoSum(numbers, target);
    for (int i : result)
    {
        cout << i << " ";
    }
    Solution_2 s2;
    for (int i : s2.twoSum(numbers, target))
    {
        cout << i << " ";
    }
}