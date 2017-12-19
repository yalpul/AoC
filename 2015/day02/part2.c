#include <stdio.h>

int calc(int x, int y, int z)
{
	int ribbon = x*y*z;
	int perimeter = x > y ? (x > z ? y+z : y+x)
	                      : (y > z ? x+z : x+y);
	return 2*perimeter + ribbon;
}

int main()
{
	int x,y,z;
	int sum = 0;

	while (scanf("%dx%dx%d", &x, &y, &z) > 0)
		sum += calc(x,y,z);
	
	printf("%d\n", sum);
	return 0;
}
