#include<iostream>
#include<list>
using namespace std;
template <typename T>
list <T> Intersection_list (const list<T> &L1,const list<T> &L2) {
		list<T> L;
	typename list<T>::const_iterator p1 = L1.begin();
	typename list<T>::const_iterator p2 = L2.begin();
	while (p1 != L1.end() && p2 != L2.end()) {
		if (*p1 < *p2) {
			p1++;
		}
		else if (*p1 > *p2) {
			p2++;
		}
		else {
			L.push_back(*p1);
			p1++;
			p2++;
		}
	}
	return L;
}

int main() {
	list<int> L1, L2;
	L1.push_back(1);
	L1.push_back(2);
	L1.push_back(3);
	L1.push_back(4);
	L1.push_back(5);
	L2.push_back(3);
	L2.push_back(4);
	L2.push_back(5);
	L2.push_back(6);
	L2.push_back(7);
	list<int> L = Intersection_list(L1, L2);
	for (list<int>::iterator p = L.begin(); p != L.end(); p++) {
		cout << *p << " ";
	}
	cout << endl;
	return 0;
}
