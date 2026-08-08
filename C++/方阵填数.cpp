#include<iostream>
int N;
using namespace std;
int lx(int a[51][51], int x, int y,int zt) {
	if (zt == 1 && (x != N)&&a[x+1][y]==0) {
		return 1;
	}
	else if (zt == 2 && (y != 1)&& a[x][y-1]==0) {
		return 2;
	}
	else if ((zt == 3) && (x != 1)&&a[x-1][y]==0) {
		return 3;
	}
	else if ((zt == 4) && (y != N) && a[x][y + 1] == 0) {
		return 4;
	}
	else
		return ++zt;
}
int main() {
	int a[51][51] = { {},{} };
	int b[3000] = {};
	int zt=1;
	cin >> N;
	int x = 1; 
	int y = N;
	int i = 1;
	a[x][y] = 1;
	for (int i = 1; i <= N*N; i++)
		b[i] = i;
	for (int i = 2; i <= N * N; i++) {
		zt = lx(a, x, y, zt);
		if (zt == 5)
			zt = 1;
		if (zt == 1) {
			a[x + 1][y] = b[i];
			x = x + 1;
		}
		else if (zt == 2) {
			a[x][y - 1] = b[i];
			y = y - 1;
		}
		else if (zt == 3) {
			a[x - 1][y] = b[i];
			x = x - 1;
		}
		else if (zt == 4) {
			a[x][y + 1] = b[i];
			y = y + 1;
		}
	}
		for (int i = 1; i <= N; i++) {
			for (int n = 1; n <= N; n++) {
				cout << a[i][n]<<" ";
			}
			cout << endl;
    }
}