#include <stdio.h>

#include "problem.h"

int main(int argc, char *argv[]) {
    Problem problem = parseProblem(argc - 1, argv + 1);

    FILE *input = readInput(&problem);
    if (problem.day == 1 && problem.part == 1) {
        problem11(input);
    } else if (problem.day == 1 && problem.part == 2) {
        problem12(input);
    } else {
        printf("Invalid day\n");
        return 1;
    }
}
