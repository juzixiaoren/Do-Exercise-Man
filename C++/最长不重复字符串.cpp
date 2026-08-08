#include <iostream>
#include <queue>
#include <string>
#include <map>
#include <vector>
#include <algorithm>
#include <unordered_set>
using namespace std;
int main1()
{
	string input;
	vector<int> v;
	for (int i = 0; i < 300; i++)
	{
		v.push_back(0);
	}
	int ans = 0;
	cin >> input;
	int temp = 0;
	for (int i = 0; i < input.length(); i++)
	{
		if (v[int(input[i])] == 0)
		{
			temp++;
			v[int(input[i])] = 1;
		}
		else if (v[int(input[i])] != 0)
		{
			if (temp > ans)
				ans = temp;
			temp = 0;
			for (int j = 0; j < 300; j++)
			{
				v[j] = 0;
			}
		}
	}
	if (temp > ans)
		ans = temp;
	cout << ans;
	return 0;
}

// 这段代码的功能是计算输入字符串中最长的不重复字符子串的长度。
// 使用了一个长度为300的向量v来记录每个字符是否已经出现过，当遇到重复字符时，更新最长子串长度并重置计数器和向量。
// 实际上这样做不对，应该使用滑动窗口

string input = "abcdaefghje";

int fixmain(string s)
{
	int left = 0;
	int right = 0;
	int ans = 0;
	int num = 0;
	unordered_set<char> st;
	for (; right < s.length(); right++)
	{
		if (st.empty())
		{
			st.insert(s[right]);
		}
		else
		{
			if (st.find(s[right]) != st.end())
			{
				num = right - left;
				ans = max(num, ans);
				for (; left < right; left++)
				{
					if (s[left] == s[right])
					{
						st.erase(s[left]);
						left++;
						break;
					}
				}
			}
			else
			{
				st.insert(s[right]);
			}
		}
	}
	num = right - left;
	ans = max(num, ans);
	return ans;
}

int main()
{
	cout << fixmain(input) << endl;
}