#include<iostream>
using namespace std;
class onelist 
{
private:
	int data;
	onelist* next;
public:
	onelist()
	{
		data = NULL;
		next = NULL;
	}
	onelist(int x)
	{
		data = x;
		next = NULL;
	}
	int Returnsizeoflist(onelist* head) 
	{
		head = head->next;
	int count = 0;
		while (head != NULL) 
		{
			count++;
			head = head->next;
		}
	return count;
	}
	void printlist(onelist* head)
	{
		head = head->next;
		while (head != NULL)
		{
			cout << head->data << " ";
			head = head->next;
		}
		cout << endl;
	}
	bool findlist(onelist* head, int x)
	{
		head = head->next;
		while (head != NULL)
		{
			if (head->data == x)
			{
				return true;
			}
			head = head->next;
		}
		return false;
	}
	bool add(onelist* head, int x)
	{
		if (findlist(head, x))
		{
			return false;
		}
		else 
		{
			while (head->next != NULL)
			{
				head = head->next;
			}
			head->next = new onelist(x);
			return true;
		}
	}
	bool remove(onelist* head, int x)
	{
		if (!findlist(head, x))
		{
			return false;
		}
		else
		{
			onelist* temp = head;
			while (temp->next->data != x)
			{
				temp = temp->next;
			}
			onelist *temp2 = temp->next;
			temp->next = temp->next->next;
			delete temp2;
			return true;
		}
	}
};
int main()
{
	onelist* head = new onelist;
	head->add(head, 1);
	head->add(head, 2);
	head->add(head, 3);
	head->add(head, 4);
	head->add(head, 5);
	cout<<"The list is: ";
	head->printlist(head);
	cout<<"The size of the list is: ";
	cout << head->Returnsizeoflist(head) << endl;
	cout << "if this list contain:";
	int x;
	cin >> x;
	if (head->findlist(head,x)) {
		cout << "yes" << endl;
	}
	else 
	{
		cout << "no" << endl;
	}
	cout << "if this list contain:";
	cin >> x;
	if (head->findlist(head,x)) {
		cout << "yes" << endl;
	}
	else {
		cout << "no" << endl;
	}
	cout << "remove if it's contain";
	cin >> x;
	if (head->remove(head, x)) 
	{
		cout<<"remove success"<<endl;
	}
	else 
	{
		cout<<"remove fail"<<endl;
	}
	cout<<"The list is: ";
	head->printlist(head);
	cout << "remove if it's contain";
	cin >> x;
	if (head->remove(head, x))
	{
		cout << "remove success" << endl;
	}
	else
	{
		cout << "remove fail" << endl;
	}
	cout << "The list is: ";
	head->printlist(head);
	cout << "add if it's not contain";
	cin >> x;
	if (head->add(head, x))
	{
		cout << "add success" << endl;
	}
	else
	{
		cout << "add fail" << endl;
	}
	cout << "The list is: ";
	head->printlist(head);
	cout << "add if it's not contain";
	cin>>x;
	if (head->add(head, x))
	{
		cout << "add success" << endl;
	}
	else
	{
		cout << "add fail" << endl;
	}
	cout << "The list is: ";
	head->printlist(head);
	return 0;
}