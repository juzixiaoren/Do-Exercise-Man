#include <iostream>
using namespace std;
int SUM;
class tree 
{
private:
	tree *left;
	tree *right;
	int data;
public:
	tree(int d)
	{
		data = d;
		left = NULL;
		right = NULL;
	}
	void insert(int d)
	{
		if (d < data)
		{
			if (left == NULL)
			{
				left = new tree(d);
			}
			else
			{
				left->insert(d);
			}
		}
		else
		{
			if (right == NULL)
			{
				right = new tree(d);
			}
			else
			{
				right->insert(d);
			}
		}
	}
	void print()
	{
		if (left != NULL)
		{
			left->print();
		}
		if (right != NULL)
		{
			right->print();
		}
		cout << data << endl;
	}
	void find(int d)
	{
		SUM++;
		if (d == data)
		{
			cout << "find" << endl;
		}
		else if (d < data)
		{
			if (left == NULL)
			{
				cout << "not find" << endl;
			}
			else
			{
				left->find(d);
			}
		}
		else
		{
			if (right == NULL)
			{
				cout << "not find" << endl;
			}
			else
			{
				right->find(d);
			}
		}
	}
};
int main() 
{
	tree root(5);
	root.insert(3);
	root.insert(7);
	root.insert(2);
	root.insert(4);
	root.insert(6);
	root.insert(8);
	root.print();
	root.find(7);
	cout << SUM << endl;
	return 0;
}