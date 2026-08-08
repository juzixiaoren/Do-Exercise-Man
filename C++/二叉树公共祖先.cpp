#include <iostream>
#include <queue>
#include <vector>
#include <algorithm>
using namespace std;
struct Nodes
{
	int id;
	int parent;
};
int main1()
{
	queue<Nodes> q;
	vector<int> tlenth;
	vector<Nodes> n = {{5, 4}, {2, 0}, {3, 1}, {1, 0}, {4, 1}, {0, -1}};
	for (int i = 0; i < n.size(); i++)
	{
		tlenth.push_back(-1);
		q.push(n[i]);
	}
	while (!q.empty())
	{
		if (q.front().id == 0)
		{
			tlenth[q.front().id] = 1;
			q.pop();
		}
		else if (tlenth[q.front().parent] == -1)
		{
			q.push(q.front());
			q.pop();
			continue;
		}
		else
		{
			tlenth[q.front().id] = tlenth[q.front().parent] + 1;
			q.pop();
		}
	}
	int ans = 1;
	for (int i = 0; i < n.size(); i++)
	{
		if (tlenth[i] > ans)
		{
			ans = tlenth[i];
		}
	}
	cout << ans;
	return 0;
}
// 题目：求二叉树的最大深度，给出若干个节点，每个节点包括自己的编号和父节点编号，节点顺序不确定，求整棵树的最大深度
// 做法：使用队列进行层次遍历，先将所有节点入队，然后依次出队，如果当前节点的父节点深度还未计算，则将当前节点重新入队，直到所有节点的深度都计算完毕。最后遍历深度数组，找出最大深度即可。
// 更好的做法，使用递归，先将所有节点存入一个map中，然后从根节点开始递归计算深度，返回最大深度即可。
int findDepth(vector<int> &parents, vector<int> &depth, int node)
{
	if (depth[node] == -1)
	{
		depth[node] = findDepth(parents, depth, parents[node]) + 1;
	}
	return depth[node];
}
int main()
{
	vector<Nodes> n = {{5, 4}, {2, 0}, {3, 1}, {1, 0}, {4, 1}, {0, -1}};
	vector<int> parents;
	vector<int> depth;
	parents.assign(n.size(), -1);
	int ans = 0;
	for (int i = 0; i < n.size(); i++)
	{
		parents[n[i].id] = n[i].parent;
		depth.push_back(-1);
	}
	depth[0] = 1;
	parents[0] = -1;
	for (int i = 0; i < n.size(); i++)
	{
		ans = max(ans, findDepth(parents, depth, n[i].id));
	}
	cout << ans << endl;
}