#include "problem.h"
#include "common.h"

#include <regex.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

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

void problem11(FILE *file) {
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
    printf("%d\n", sum);
    free(line);
}

void problem12(FILE *file) {
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
    printf("%d\n", sum);
    free(line);
}

typedef struct
{
    int green;
    int red;
    int blue;
} GameSet;

void problem21(FILE *file) {
    char *line = (char *)malloc(MAX_BUFFER_SIZE * sizeof(char));

    int sum = 0;
    while (fgets(line, MAX_BUFFER_SIZE, file)) {
        int len = strlen(line);
        char *it = line;
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0';
            len -= 1;
        }
        it += 5;
        char *colon = strchr(it, ':');
        int gameId = natoi(it, colon - it);
        it = colon + 2;

#ifdef DEBUG_EXEC
        printf("line  : %s\ngameId: %d\n", line, gameId);
#endif
        int setSize = countnchar(it, strlen(it), ';') + 1;
        GameSet *sets = (GameSet *)malloc(setSize * sizeof(GameSet));

        int hasInvalidSet = FALSE;
        for (int i = 0; i < setSize; i++) {
            sets[i].green = sets[i].red = sets[i].blue = 0;
            char *semicolon = strchr(it, ';');
            if (semicolon == NULL) {
                semicolon = it + strlen(it);
            }
            while (*it != ';' && *it != '\0') {
                char *sep = strchr(it, ' ');
                int count = natoi(it, sep - it);
                it = sep + 1;
                switch (*it) {
                case 'r':
                    sets[i].red = count;
                    it += 3;
                    break;
                case 'g':
                    sets[i].green = count;
                    it += 5;
                    break;
                case 'b':
                    sets[i].blue = count;
                    it += 4;
                    break;
                }
            }

            if (sets[i].red > 12 || sets[i].green > 13 || sets[i].blue > 14) {
                hasInvalidSet = TRUE;
            }
#ifdef DEBUG_EXEC
            printf("\tset[%d]: %d,%d,%d\n", i, sets[i].red, sets[i].green, sets[i].blue);
#endif
            it++;
        }

        if (!hasInvalidSet) {
            sum += gameId;
        }

        free(sets);
    }
    printf("%d\n", sum);
    free(line);
}

void problem22(FILE *file) {
    char *line = (char *)malloc(MAX_BUFFER_SIZE * sizeof(char));

    int sum = 0;
    while (fgets(line, MAX_BUFFER_SIZE, file)) {
        int len = strlen(line);
        char *it = line;
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0';
            len -= 1;
        }
        it += 5;
        char *colon = strchr(it, ':');
        int gameId = natoi(it, colon - it);
        it = colon + 2;

#ifdef DEBUG_EXEC
        printf("line  : %s\ngameId: %d\n", line, gameId);
#endif
        int setSize = countnchar(it, strlen(it), ';') + 1;
        GameSet *sets = (GameSet *)malloc(setSize * sizeof(GameSet));

        int maxRed = 0, maxGreen = 0, maxBlue = 0;
        for (int i = 0; i < setSize; i++) {
            sets[i].green = sets[i].red = sets[i].blue = 0;
            char *semicolon = strchr(it, ';');
            if (semicolon == NULL) {
                semicolon = it + strlen(it);
            }
            while (*it != ';' && *it != '\0') {
                char *sep = strchr(it, ' ');
                int count = natoi(it, sep - it);
                it = sep + 1;
                switch (*it) {
                case 'r':
                    if (count > maxRed) {
                        maxRed = count;
                    }
                    sets[i].red = count;
                    it += 3;
                    break;
                case 'g':
                    if (count > maxGreen) {
                        maxGreen = count;
                    }
                    sets[i].green = count;
                    it += 5;
                    break;
                case 'b':
                    if (count > maxBlue) {
                        maxBlue = count;
                    }
                    sets[i].blue = count;
                    it += 4;
                    break;
                }
            }

#ifdef DEBUG_EXEC
            printf("\tset[%d]: %d,%d,%d\n", i, sets[i].red, sets[i].green, sets[i].blue);
#endif
            it++;
        }

        int power = maxRed * maxGreen * maxBlue;
        sum += power;

        free(sets);
    }
    printf("%d\n", sum);
    free(line);
}