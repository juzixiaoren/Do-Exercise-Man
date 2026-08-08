#include<iostream>
#include<array>
using namespace std;
int main() {
	int i=1;
	array<int, 10000>a;
	while (cin >> a[i]) {
		if (a[i] == 0)break;
		i++;
	}
	i--;
	for (int n = i; n >= 1; n--,i--) {
		cout << a[i]<<" ";
	}

}