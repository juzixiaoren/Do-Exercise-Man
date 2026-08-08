#include <iostream>
#include <cstring>
#include <string.h>
#include <string>
int zs(int a) {
	while (a % 10 == 0) {
		a = a / 10;
	}
	return a;
}
int ws(int a) {
	int n = 0;
	while (a != 0) {
		n++;
		a = a / 10;
	}
	return n;
}
int fz(int input) {
	int output = 0;
	int a = zs(input);
	int b = ws(a);
	for (int n = 0; n < b; n++) {
		output += a % 10 * pow(10, b-n-1);
		a = a / 10;
	}
	return output;
}
using namespace std;
int main() {
	char a[1000];
	char* b = NULL;
	cin >> a;
	if (strchr(a, '/') != NULL) {
		b = strchr(a, '/');
		int i = atoi(b + 1);
		int p = atoi(a);
		cout << zs(fz(p)) << "/" << zs(fz(i)) << endl;
	}
	else if (strchr(a, '.') != NULL) {
		b = strchr(a, '.');
		int i = atoi(b + 1);
		int p = atoi(a);
		cout << zs(fz(p)) << "." << zs(fz(i)) << endl;
	}
	else if (strchr(a, '%') != NULL) {
		int p = atoi(a);
		cout << zs(fz(p)) << "%" << endl;
	}
	else cout << zs(fz(atoi(a))) << endl;

}