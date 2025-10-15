#ifndef MYLIB_H
#define MYLIB_H

int add(int a, int b);
char* concat(const char* str1, const char* str2);
typedef struct {
    int x;
    int y;
} Point;
void movePoint(Point* p, int dx, int dy);
void printArray(int* arr, int length);
#endif