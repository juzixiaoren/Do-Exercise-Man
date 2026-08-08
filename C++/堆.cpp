#include <iostream>
#include <vector>
#include <Windows.h>
#include <chrono>
#include <fstream>
#include <sstream>
#include <time.h>
using namespace std;
class heapTree
{
private:
	vector<int> heap;
public:
	heapTree()
	{
		heap.push_back(0);
	}
	void insect(int num) {
		int tempsize= heap.size();
		heap.resize(tempsize + 1);
		while (tempsize > 1)
		{
			if (num < heap[tempsize / 2])
			{
				heap[tempsize] = heap[tempsize / 2];
				tempsize /= 2;
			}
			else
			{
				heap[tempsize] = num;
				break;
			}
		}
		if (tempsize == 1)
			heap[1] = num;
	}
	void print() 
	{
		for (int i = 1; i < heap.size(); i++)
		{
			cout << heap[i] << " ";
		}
		cout << endl;
	}
	void deleteMin() 
	{
		heap[1] = NULL;
		int temp =heap[heap.size() - 1];
		heap.pop_back();
		PercDown(1, temp);
	}
	void PercDown(int i, int x)//i is the hole x is the value to insert
	{
		if (2 * i > heap.size() - 1)
			heap[i] = x;
		else if (2 * i == heap.size() - 1)
		{
			if (heap[2 * i] < x)
			{
				heap[i] = heap[2 * i];
				heap[2 * i] = x;
			}
			else
			{
				heap[i] = x;
			}
		}
		else 
		{
			int j;
			if (heap[2*i + 1] > heap[2*i]) {
				j=2*i;
			}
			else {
				j=2*i+1;
			}
			if (heap[j] < x)
			{
				heap[i] = heap[j];
				PercDown(j, x);
			}
			else
			{
				heap[i] = x;
			}
		}
	}
	void buildHeapLinear(int size)
	{
		heap.resize(size + 1); // 直接分配足够的空间，heap[0]保持为占位符
		for (int i = 1; i <= size; i++) // 从索引1开始添加元素
		{
			int temp;
			cin >> temp;
			heap[i] = temp;
		}
		for (int i = size / 2; i > 0; i--)
		{
			PercDown(i, heap[i]);
		}
	}
	void buildHeap(int size) {
		for (int i = 1; i <= size; i++) {
			int temp;
			cin >> temp;
			insect(temp);
		}
	}
	void buildHeapLinearFromFile(const std::string& filename) {
		std::ifstream file(filename);
		if (!file.is_open()) {
			std::cerr << "无法打开文件: " << filename << std::endl;
			return;
		}

		std::string line;
		int number;
		heap.clear();
		heap.push_back(0); // 保留位置0为哨兵

		while (getline(file, line)) {
			std::istringstream iss(line);
			while (iss >> number) {
				heap.push_back(number);
			}
		}

		// 构建最小堆
		for (int i = heap.size() / 2; i > 0; --i) {
			PercDown(i,heap[i]);
		}
	}
	void buildHeapFromFile(const std::string& filename) {
		std::ifstream file(filename);
		if (!file.is_open()) {
			std::cerr << "无法打开文件: " << filename << std::endl;
			return;
		}
		std::string line;
		int number;
		heap.clear();
		heap.push_back(0); // 保留位置0为哨兵

		while (getline(file, line)) {
			std::istringstream iss(line);
			while (iss >> number) {
				heap.insert(heap.begin() + 1, number);
			}
		}
	}
};
int main()
{
	/*cout<<"请输入元素个数:"<<endl;
	int n;
	cin>>n;*/
	//heapTree b;
	//cout<<"请输入元素，中间用空格隔开:"<<endl;
	//b.buildHeapLinear(n);
	////cin>>1 3 2 12 6 4 8 15 14 9 7 5 11 13 10;
	//cout<<"初始堆:"<<endl;
	//b.print();
	//b.deleteMin();
	//cout<<"第一次删除最小值后的堆:"<<endl;
	//b.print();
	//b.deleteMin();
	//cout<<"第二次删除最小值后的堆:"<<endl;
	//b.print();
	//b.deleteMin();
	//cout<<"第三次删除最小值后的堆:"<<endl;
	//b.print();
	//heapTree a;
	//cout << "请输入元素个数:" << endl;
	//cin >> n;
	//cout << "请输入元素，中间用空格隔开:" << endl;
	//a.buildHeap(n)
	////cin>>1 3 2 12 6 4 8 15 14 9 7 5 11 13 10;
	//cout << "初始堆:" << endl;
	//a.print();
	//a.deleteMin();
	//cout<<"第一次删除最小值后的堆:"<<endl;
	//a.print();
	//a.deleteMin();
	//cout << "第二次删除最小值后的堆:" << endl;
	//a.print();
	//a.deleteMin();
	//cout << "第三次删除最小值后的堆:" << endl;
	//a.print();
	/*heapTree c;*/
	//cout << "请输入元素,忽略第一个数字:" << endl;
	//int temp;
	//cin >> temp;
	//auto start=chrono::high_resolution_clock::now();
	//c.buildHeap(n);
	//auto end = chrono::high_resolution_clock::now();
	//chrono::duration<double> elapsed = end - start;
	//cout << "建堆时间:" << elapsed.count() << "s" << endl;
	////c.print();
	//heapTree d;
	//cout << "请输入元素,忽略第一个数字:" << endl;
	//cin>> temp;
	//auto start1 = chrono::high_resolution_clock::now();
	//d.buildHeapLinear(n);
	////d.print();
	//auto end1= chrono::high_resolution_clock::now();
	//chrono::duration<double> elapsed1 = end1 - start1;
	//cout << "建堆时间:" << elapsed1.count()<< "s" << endl;
	heapTree e;
	clock_t start;
	clock_t end;
	start = clock();
	e.buildHeapFromFile("heap.txt");
	end = clock();
	double time;
	time=(double)(end - start) / CLOCKS_PER_SEC;
	cout << "建堆时间:" << time << "s" << endl;
	heapTree f;
	clock_t start1;
	clock_t end1;
	start1 = clock();
	f.buildHeapLinearFromFile("heap.txt");
	end1 = clock();
	double time1;
	time1 = (double)(end1 - start1) / CLOCKS_PER_SEC;
	cout << "建堆时间:" << time1 << "s" << endl;
	heapTree g;
	clock_t start3;
	clock_t end3;
	start3 = clock();
	e.buildHeapFromFile("heap2.txt");
	end3 = clock();
	double time3;
	time3 = (double)(end3 - start3) / CLOCKS_PER_SEC;
	cout << "建堆时间:" << time3 << "s" << endl;
	heapTree h;
	clock_t start4;
	clock_t end4;
	start4 = clock();
	f.buildHeapLinearFromFile("heap2.txt");
	end4 = clock();
	double time4;
	time4 = (double)(end4 - start4) / CLOCKS_PER_SEC;
	cout << "建堆时间:" << time4 << "s" << endl;
	heapTree i;
	clock_t start5;
	clock_t end5;
	start5 = clock();
	e.buildHeapFromFile("heap3.txt");
	end5 = clock();
	double time5;
	time5 = (double)(end5 - start5) / CLOCKS_PER_SEC;
	cout << "建堆时间:" << time5 << "s" << endl;
	heapTree j;
	clock_t start6;
	clock_t end6;
	start6 = clock();
	f.buildHeapLinearFromFile("heap3.txt");
	end6 = clock();
	double time6;
	time6 = (double)(end6 - start6) / CLOCKS_PER_SEC;
	cout << "建堆时间:" << time6 << "s" << endl;
	return 0;
}