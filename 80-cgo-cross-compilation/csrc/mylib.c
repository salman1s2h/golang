#include <stdio.h>
#include <string.h>
#include <stdlib.h>

// Function to add two integers
int add(int a, int b) {
    return a + b;
}

// Function to concatenate two strings
char* concat(const char* str1, const char* str2) {
    char* result = malloc(strlen(str1) + strlen(str2) + 1); // Allocate memory for the result
    strcpy(result, str1);
    strcat(result, str2);
    return result;
}

// Function to modify a struct
typedef struct {
    int x;
    int y;
} Point;

void movePoint(Point* p, int dx, int dy) {
    p->x += dx;
    p->y += dy;
}

// Function to print an array of integers
void printArray(int* arr, int length) {
    for (int i = 0; i < length; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}