#include <iostream>
#include <string>
#include <iomanip>
#include <stdlib.h>
#include <cstring>
#include <stdlib.h>
#include <math.h>
using namespace std;
class Myi
{
private:
	string data[10];

public:
	Myi(string in, string a)
	{
		int j = 0;
		string temp = "";
		for (int i = 0; i < in.size(); i++)
		{
			if (in[i] == a[0])
			{
				data[j++] = temp;
				temp = "";
			}
			else
				temp = temp + in[i];
		}
		data[j] = temp;
	}
	string Get(int index)
	{
		return data[index];
	}
};
class Point
{
private:
	double x, y;
public:
	void GetPoint(double a = 0, double b = 0)
	{
		x = a;
		y = b;
	}
	double Getx()
	{
		return x;
	}
	double Gety()
	{
		return y;
	}
	double Distance(Point p)
	{
		return sqrt((x - p.x) * (x - p.x) + (y - p.y) * (y - p.y));
	}
	double Getx_1(Point p)
	{
		return x - p.x;
	}
	double Gety_1(Point p)
	{
		return y - p.y;
	}
};
class Rectangle
{
private:
	Point p1, p2, p3, p4;
	double S;
public:
	Rectangle()
	{
	};
	void input(double x1, double x2, double x3, double x4, double y1, double y2, double y3, double y4)
	{
		p1.GetPoint(x1, y1);
		p2.GetPoint(x2, y2);
		p3.GetPoint(x3, y3);
		p4.GetPoint(x4, y4);
	}
	int Judge()
	{
		if (abs(p1.Getx_1(p2) * p3.Gety_1(p4) - p3.Getx_1(p4) * p1.Gety_1(p2)) <= 1e-8 && abs(p2.Getx_1(p3) * p4.Gety_1(p1) - p4.Getx_1(p1) * p2.Gety_1(p3)) <= 1e-8)
			return 1;
		else
			return 0;
	}
	int Judge_s()
	{
		if (abs(p1.Distance(p2) == 0 || p2.Distance(p3) == 0))
			return 0;
		else if (abs(p1.Distance(p2) - p3.Distance(p4)) < 1e-8 && abs(p2.Distance(p3) - p4.Distance(p1)) < 1e-8)
			return 1;
		else return 0;
	}
	int Judge_p() {
		if(abs(p1.Getx_1(p2) * p2.Getx_1(p3) + p1.Gety_1(p2) * p2.Gety_1(p3)) < 1e-8 && abs(p3.Getx_1(p4) * p4.Getx_1(p1) + p3.Gety_1(p4) * p4.Gety_1(p1)) < 1e-8)
		{
			return 1;
		}
		else
			return 0;
	}
	double Area()
	{
		S = abs(p1.Distance(p2) * p2.Distance(p3));
		return S;
	}
};
int main() 
{
	int j=0;
	string in;
	Rectangle* r = new Rectangle[100];
	//Book* b = new Book[n];
	while(getline(cin, in))
	{
		Myi inp(in, ",");
		double x1 = stod(inp.Get(0));
		double y1 = stod(inp.Get(1));
		double x2 = stod(inp.Get(2));
		double y2 = stod(inp.Get(3));
		double x3 = stod(inp.Get(4));
		double y3 = stod(inp.Get(5));
		double x4 = stod(inp.Get(6));
		double y4 = stod(inp.Get(7));
		r[j].input(x1, x2, x3, x4, y1, y2, y3, y4);
		j++;
	}
	for (int i = 0; i < j; i++) 
	{
		if (r[i].Judge() == 1) 
			if (r[i].Judge_s()==1)
				if(r[i].Judge_p()==1)
			cout<<fixed<<setprecision(4)<<r[i].Area()<<endl;
				else cout<<"0"<<endl;
			else cout<<"0"<<endl;
		else cout<<"0"<<endl;
	}
}