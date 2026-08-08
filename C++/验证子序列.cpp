#include<iostream>
#include<vector>
#include<string>
using namespace std;
class Solution {
public:
    bool isSubsequence(string s, string t) {
        int left=0;
        int right=0;
        while (left < s.size() && right < t.size()) {
            if (s[left] == t[right]) {
                left++;
			}
            right++;
        }
        if(left==s.size()) return true;
		else return false;
    }
};