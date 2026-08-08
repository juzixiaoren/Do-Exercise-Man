#include<vector>
#include<iostream>
#include<queue>
#include<cmath>
using namespace std;

struct pos
{
	int x;
	int y;
};
struct box
{
	int x;
	int y;
	bool find;
};
vector<vector<int>> vdist(vector<vector<char>> tmap, vector<vector<int>> dist, int px, int py) {
	for (int i = 0; i < dist.size();i++) {
		for (int j = 0; j < dist[0].size();j++) {
			dist[i][j] = -1;
		}
	}
	dist[py][px] = 0;
	queue<pos> q;
	q.push({ px,py });
	while (!q.empty()) {
		pos nowpos;
		nowpos = q.front();
		vector<int>yp = { -1,1,0,0 };
		vector<int>xp = { 0,0,-1,1 };
		int nx, ny;
		for (int l = 0; l < 4; l++) {
			ny = nowpos.y + yp[l];
			nx = nowpos.x + xp[l];
			if (ny < 0 || ny == tmap.size())
				continue;
			if (nx < 0 || nx == tmap[0].size())
				continue;
			if (tmap[ny][nx] != '#' && dist[ny][nx] == -1)
			{
				dist[ny][nx] = dist[nowpos.y][nowpos.x] + 1;
				q.push({ nx,ny });
			}
		}
		q.pop();
	}
	return dist;
}
int main() {
	int T;
	cin >> T;
	vector<int>step(T,0);
	for (int t = 0; t < T; t++) {
		int m, n;
		cin >> m;
		cin >> n;
		vector<vector<char>> tmap(m, vector<char>(n));
		vector<vector<int>>dist(m, vector<int>(n, -1));
		for (int i = 0; i < m; i++)
		{
			string s;
			cin >> s;
			for (int j = 0; j < n; j++) {
				tmap[i][j] = s[j];
			}
		}
		int px;
		int py;
		for (int i = 0; i < m; i++) {
			for (int j = 0; j < n; j++) {
				if (tmap[i][j] == '*')
				{
					px = j;
					py = i;
					goto finished;
				}
			}
		}
	finished:
		vector <box> boxp(10);
		int count = 0;
		for (int i = 0; i < m; i++) {
			for (int j = 0; j < n; j++) {
				if (tmap[i][j] <= '9' && tmap[i][j] >= '0') {
					count++;
					boxp[tmap[i][j] - '0'] = { j,i,false };
				}
			}
		}
		while (true) {
			int best = -1;
			int worse = 99999;
			if (step[t] > 10000) {
				step[t] = -1;
				break;
			}
			for (int i = 0; i < count; i++) {
				if ((abs(boxp[i].x - px) + abs(boxp[i].y - py)) < worse && boxp[i].find == false)
				{
					worse = abs(boxp[i].x - px) + abs(boxp[i].y - py);
					best = i;
				}
			}
			if (best == -1) {
				break;
			}
			dist = vdist(tmap, dist, boxp[best].x,boxp[best].y);
			if (dist[py][px] == -1) {
				step[t] = -1;
				break;
			}
			if ((py - 1 >= 0) && dist[py - 1][px] == dist[py][px] - 1) {
				py = py - 1;
				step[t] += 1;
				if (px == boxp[best].x && py == boxp[best].y) {
					boxp[best].find = true;
					continue; // 回到选下一个宝箱
				}
			}
			else if ((py + 1 < m) && dist[py + 1][px] == dist[py][px] - 1) {
				py = py + 1;
				step[t] += 1;
				if (px == boxp[best].x && py == boxp[best].y) {
					boxp[best].find = true;
					continue; // 回到选下一个宝箱
				}
			}
			else {
				if ((px - 1 >= 0) && dist[py][px - 1] == dist[py][px] - 1) {
					px = px - 1;
					step[t] += 1;
					if (px == boxp[best].x && py == boxp[best].y) {
						boxp[best].find = true;
						continue; // 回到选下一个宝箱
					}
				}
				else if ((px + 1 < n) && dist[py][px + 1] == dist[py][px] - 1) {
					px = px + 1;
					step[t] += 1;
					if (px == boxp[best].x && py == boxp[best].y) {
						boxp[best].find = true;
						continue; // 回到选下一个宝箱
					}
				}
				else {
					if (boxp[best].x == px && boxp[best].y == py)
						boxp[best].find = true;
					else
						step[t] = -1;
					break;
				}
			}
		}
	}
	for (int i = 0; i < T; i++) {
		cout << step[i] << endl;
	}
}