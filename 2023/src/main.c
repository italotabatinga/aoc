#include <stdio.h>

#include "problem.h"

int main(int argc, char *argv[]) {
    Problem problem = parseProblem(argc - 1, argv + 1);

    FILE *input = readInput(&problem);
    if (problem.day == 1 && problem.part == 1) {
        problem11(input);
    } else if (problem.day == 1 && problem.part == 2) {
        problem12(input);
    } else if (problem.day == 2 && problem.part == 1) {
        problem21(input);
    } else if (problem.day == 2 && problem.part == 2) {
        problem22(input);
    } else if (problem.day == 3 && problem.part == 1) {
        problem31(input);
    } else if (problem.day == 3 && problem.part == 2) {
        problem32(input);
    } else if (problem.day == 4 && problem.part == 1) {
        problem41(input);
    } else if (problem.day == 4 && problem.part == 2) {
        problem42(input);
    } else {
        printf("Invalid day\n");
        return 1;
    }
}
