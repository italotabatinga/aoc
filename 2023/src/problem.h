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

char *problem11(FILE* file);
char *problem12(FILE* file);