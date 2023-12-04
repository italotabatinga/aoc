#include "problem.h"
#include "common.h"

#include <regex.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_BUFFER_SIZE 256

Problem parseProblem(int argc, char *argv[]) {
    Problem problem;
    problem.test = FALSE;
    regex_t regex;
    regmatch_t groups[3];

    if (regcomp(&regex, "^([1-3]?[0-9])(\\.[12])?$", REG_EXTENDED)) {
        printf("Invalid regex\n");
        exit(1);
    }

    for (int i = 0; i < argc; i++) {
        if (strcmp(argv[i], "--test") == 0) {
            problem.test = TRUE;
        } else if (regexec(&regex, argv[i], 3, groups, 0) == 0) {
            char day[3];
            strncpy(day, argv[i] + groups[1].rm_so, groups[1].rm_eo - groups[1].rm_so);
            day[groups[1].rm_eo - groups[1].rm_so] = '\0';
            problem.day = atoi(day);

            if (groups[2].rm_so != -1) {
                char part[2];
                strncpy(part, argv[i] + groups[2].rm_so + 1, groups[2].rm_eo - groups[2].rm_so - 1);
                part[groups[2].rm_eo - groups[2].rm_so - 1] = '\0';
                problem.part = atoi(part);
            }
        }
    }

    regfree(&regex);
    return problem;
}

FILE *readInput(Problem *problem) {
    char fileString[25];
    sprintf(fileString, "files/%d.%d", problem->day, problem->part);
    if (problem->test) {
        strcat(fileString, "_test");
    }
    strcat(fileString, ".txt");
    FILE *file = fopen(fileString, "r");
    if (file == NULL) {
        printf("Could not open file %s\n", fileString);
        exit(1);
    }

    return file;
}

char *problem11(FILE *file) {
    char *line = (char *)malloc(MAX_BUFFER_SIZE * sizeof(char));

    int sum = 0;
    while (fgets(line, MAX_BUFFER_SIZE, file)) {
        int len = strlen(line);
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0';
            len -= 1;
        }

        int firstDigit = -1;
        int lastDigit = -1;
        for (int i = 0; i < len; i++) {
            char c = line[i];
            if (c >= '0' && c <= '9') {
                if (firstDigit == -1) {
                    firstDigit = lastDigit = c - '0';
                } else {
                    lastDigit = c - '0';
                }
            }
        }
#ifdef DEBUG_EXEC
        printf("line: %50s | fd: %d, sd: %d, num: %d", line, firstDigit, lastDigit, firstDigit * 10 + lastDigit);
        printf("\n", line);
#endif

        sum += firstDigit * 10 + lastDigit;
    }
    printf("sum: %d\n", sum);

    return NULL;
}

char *problem12(FILE *file) {
    char *line = (char *)malloc(MAX_BUFFER_SIZE * sizeof(char));

    int sum = 0;
    while (fgets(line, MAX_BUFFER_SIZE, file)) {
        int len = strlen(line);
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0';
            len -= 1;
        }

        int firstDigit = -1;
        int lastDigit = -1;
        for (int i = 0; i < len; i++) {
            char c = line[i];
            int digit = -1;
            if (c >= '0' && c <= '9') {
                digit = c - '0';
            } else if (strncmp(line + i, "one", 3) == 0) {
                digit = 1;
            } else if (strncmp(line + i, "two", 3) == 0) {
                digit = 2;
            } else if (strncmp(line + i, "three", 5) == 0) {
                digit = 3;
            } else if (strncmp(line + i, "four", 4) == 0) {
                digit = 4;
            } else if (strncmp(line + i, "five", 4) == 0) {
                digit = 5;
            } else if (strncmp(line + i, "six", 3) == 0) {
                digit = 6;
            } else if (strncmp(line + i, "seven", 5) == 0) {
                digit = 7;
            } else if (strncmp(line + i, "eight", 5) == 0) {
                digit = 8;
            } else if (strncmp(line + i, "nine", 4) == 0) {
                digit = 9;
            }

            if (digit == -1) {
                continue;
            }

            if (firstDigit == -1) {
                firstDigit = digit;
            }
            lastDigit = digit;
        }
#ifdef DEBUG_EXEC
        printf("line: %50s | fd: %d, sd: %d, num: %d", line, firstDigit, lastDigit, firstDigit * 10 + lastDigit);
        printf("\n", line);
#endif

        sum += firstDigit * 10 + lastDigit;
    }
    printf("sum: %d\n", sum);

    return NULL;
}