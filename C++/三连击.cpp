#include <iostream>
using namespace std;
int hj(int i, int p)
{
	for (int k = 0; k < p; k++)
	{
		i = i / 10;
	}
		i = i % 10;
	return i;
}
int pd(int i, int p, int k) {
	int a[9] = { 0 };
	int d = 8;
	for (int c = 0; c < 3; c++)
		a[c] = hj(i, c);
	for (int c = 3; c < 6; c++)
		a[c] = hj(p, c-3);
	for (int c = 6; c < 9; c++)
		a[c] = hj(k, c - 6);
	for (int i = 0; i < 8; i++) {
		if (a[i] == 0)return 0;;
	}
	for (int c = 2; c < 9; c++) {
		for (int i = d - 1; i >= 0; i--) {
			if (a[d] == a[i])
				return 0;
		}
			d--;
	}
	return 1;
}
int main() 
{
	for (int i = 123; i <= 333; i++) 
	{
		for (int p = 248; p <= 666; p = p + 2) 
		{
			for (int k = 369; k <= 999; k = k + 3) 
			{
				if (k == i * 3 && p == i * 2) 
				{
					if (i % 10 == 0 || p % 10 == 0 || k % 10 == 0)
						break;
					if(pd(i,p,k)==1)
					cout<<i<<" "<<p<<" "<<k<<endl;
				}
			}
		}
	}
}