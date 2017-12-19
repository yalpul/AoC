#include <stdio.h>

int calc(int x, int y, int z)
{
	int surface1 = x*y;
	int surface2 = y*z;
	int surface3 = x*z;

	int min = surface1 < surface2 ? (surface1 < surface3 ? surface1 : surface3)
                                : (surface2 < surface3 ? surface2 : surface3);
	return 2*(surface1+surface2+surface3) + min;
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
