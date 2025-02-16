#include <stdio.h>
#include <"main.h">

int Add(int a, int b) {
  return a + b;
}

int main(int argc, char *argv[])
{
  int result = Add(1, 2);
  printf("What is good in the neighborhood!\n");
  return 0;
}

