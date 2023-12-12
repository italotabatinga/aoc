#pragma once

#define PART_1 1
#define PART_2 2

#include <stdio.h>

struct Input
{
    char *value;
};

typedef struct
{
    int day;
    int part;
    int test;
} Problem;

Problem parseProblem(int argc, char *argv[]);
FILE *readInput(Problem *problem);

void problem11(FILE* file);
void problem12(FILE* file);
void problem21(FILE* file);
void problem22(FILE* file);
void problem31(FILE* file);
void problem32(FILE* file);