#include "problem.h"
#include "common.h"
#include "vec.h"

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
        it = colon + 2;

#ifdef DEBUG_EXEC
        int gameId = natoi(it, colon - it);
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

void problem31(FILE *file) {
    char **lines = (char **)malloc(200 * sizeof(char *));
    int linesSize = 0;
    int sum = 0;
    while (!feof(file)) {
        char *line = (char *)malloc(MAX_BUFFER_SIZE * sizeof(char));
        fgets(line, MAX_BUFFER_SIZE, file);
        int len = strlen(line);
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0';
        }
        lines[linesSize++] = line;
    }

#ifdef DEBUG_EXEC
    for (int i = 0; i < linesSize; i++) {
        printf("% 4d: %s\n", strlen(lines[i]), lines[i]);
    }
#endif

    int number = -1;
    int hasCloseSymbol = 0;

    for (int i = 0; i < linesSize; i++) {
        char *line = lines[i];
        int len = strlen(line);
        for (int j = 0; j < len; j++) {
            int jIsNumber = line[j] >= '0' && line[j] <= '9';
            if (jIsNumber) {
                if (number == -1) {
                    number = line[j] - '0';
                } else {
                    number = number * 10 + line[j] - '0';
                }
            }

            int directions[8][2] = {
                {-1, 0},  // Up
                {1, 0},   // Down
                {0, -1},  // Left
                {0, 1},   // Right
                {-1, -1}, // Up Left
                {-1, 1},  // Up Right
                {1, -1},  // Down Left
                {1, 1}    // Down Right
            };

            if (number > 0 && hasCloseSymbol <= 0) {
                for (size_t k = 0; k < 8; k++) {
                    int x = i + directions[k][0];
                    int y = j + directions[k][1];
                    if (x >= 0 && x < linesSize && y >= 0 && y < len) {
                        char c = lines[x][y];

                        if (c != '.' && (c < '0' || c > '9')) {
                            hasCloseSymbol = 1;
                        }
                    }
                }
            }

            if (number != -1 && (j == len - 1 || !(line[j + 1] >= '0' && line[j + 1] <= '9'))) {
#ifdef DEBUG_EXEC
                printf("number: %d, %d\n", number, hasCloseSymbol);
#endif
                if (hasCloseSymbol) {
                    sum += number;
                }
                number = -1;
                hasCloseSymbol = 0;
            }
        }
    }
    printf("%d\n", sum);
    for (int i = 0; i < linesSize; i++) {
        free(lines[i]);
    }
    free(lines);
}

void problem32(FILE *file) {
    const int ARRAY_SIZE = 200;
    char **lines = (char **)malloc(ARRAY_SIZE * sizeof(char *));
    int gearCount[ARRAY_SIZE][ARRAY_SIZE];
    int gearRatio[ARRAY_SIZE][ARRAY_SIZE];
    int linesSize = 0;
    int sum = 0;

    for (size_t i = 0; i < ARRAY_SIZE; i++) {
        for (size_t j = 0; j < ARRAY_SIZE; j++) {
            gearCount[i][j] = gearRatio[i][j] = 0;
        }
    }

    while (!feof(file)) {
        char *line = (char *)malloc(MAX_BUFFER_SIZE * sizeof(char));
        fgets(line, MAX_BUFFER_SIZE, file);
        int len = strlen(line);
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0';
        }
        lines[linesSize] = line;
        linesSize++;
    }

#ifdef DEBUG_EXEC
    for (int i = 0; i < linesSize; i++) {
        printf("% 4d: %s\n", strlen(lines[i]), lines[i]);
    }
#endif

    int number = -1, numberStart = -1, numberEnd = -1;

    for (int i = 0; i < linesSize; i++) {
        char *line = lines[i];
        int len = strlen(line);
        for (int j = 0; j < len; j++) {
            int jIsNumber = isNumber(line[j]);
            if (jIsNumber) {
                if (number == -1) {
                    number = line[j] - '0';
                    numberStart = j;
                } else {
                    number = number * 10 + line[j] - '0';
                }
            }

            if (number > -1 && (j == len - 1 || !isNumber(line[j + 1]))) {
                numberEnd = j;

                for (int y = i - 1; y <= i + 1; y++) {
                    for (int x = numberStart - 1; x <= numberEnd + 1; x++) {
                        if (y >= 0 && y < linesSize && x >= 0 && x < len && lines[y][x] == '*') {

                            gearCount[y][x]++;
                            if (gearRatio[y][x] == 0) {
                                gearRatio[y][x] = 1;
                            }
                            gearRatio[y][x] *= number;
                        }
                    }
                }

                number = numberStart = numberEnd = -1;
            }
        }
    }

    for (int i = 0; i < linesSize; i++) {
        for (int j = 0; j < linesSize; j++) {
            if (gearCount[i][j] == 2) {
                sum += gearRatio[i][j];
            }
        }
    }
    printf("%d\n", sum);
    for (int i = 0; i < linesSize; i++) {
        free(lines[i]);
    }
    free(lines);
}

void problem41(FILE *file) {
    const int ARRAY_SIZE = 205;
    char **lines = (char **)malloc(ARRAY_SIZE * sizeof(char *));
    int linesSize = 0;
    int sum = 0;

    VEC_NEW(winning);
    VEC_NEW(picked);

    while (!feof(file)) {
        char *line = (char *)malloc(MAX_BUFFER_SIZE * sizeof(char));
        fgets(line, MAX_BUFFER_SIZE, file);
        int len = strlen(line);
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0';
        }
        lines[linesSize] = line;
        linesSize++;
    }

    for (int i = 0; i < linesSize; i++) {
        VEC_CLR(winning);
        VEC_CLR(picked);

        char *line = lines[i];
        char *colon = strchr(line, ':');
        char *verticalBar = strchr(colon, '|');
        char *it = colon + 2;

        while (it < verticalBar) {
            int nextNumber = strcspn(it, "1234567890");
            char *nextSpace = strchr(it + nextNumber, ' ');
            int number = natoi(it + nextNumber, nextSpace - it - nextNumber);
            VEC_ADD(winning, number);
            it = nextSpace + 1;
        }
        it = verticalBar + 2;
        while (it < line + strlen(line)) {
            int nextNumber = strcspn(it, "1234567890");
            char *nextSpace = strchr(it + nextNumber, ' ');
            if (nextSpace == NULL) {
                nextSpace = it + strlen(it);
            }
            int number = natoi(it + nextNumber, nextSpace - it - nextNumber);
            VEC_ADD(picked, number);
            it = nextSpace + 1;
        }
#ifdef DEBUG_EXEC
        printf("%s\n", line);
        printf("\twinning %s -- ", VEC_STR(winning));
        for (int j = 0; j < VEC_LEN(winning); j++) {
            printf("%d ", VEC_GET(winning, int, j));
        }
        printf("\n\tpicked %s -- \t", VEC_STR(picked));
        for (int j = 0; j < VEC_LEN(picked); j++) {
            printf("%d ", VEC_GET(picked, int, j));
        }
        printf("\n");
#endif

        int score = 0;
        for (int j = 0; j < VEC_LEN(winning); j++) {
            int number = VEC_GET(winning, int, j);
            int hasNumber = FALSE;
            for (int k = 0; k < VEC_LEN(picked); k++) {
                if (number == VEC_GET(picked, int, k)) {
                    if (score == 0) {
                        score = 1;
                    } else {
                        score *= 2;
                    }
                    break;
                }
            }
        }

        sum += score;
    }

    printf("%d\n", sum);
}

void problem42(FILE *file) {
    const int ARRAY_SIZE = 205;
    char **lines = (char **)malloc(ARRAY_SIZE * sizeof(char *));
    int countCards[ARRAY_SIZE];
    for (size_t i = 0; i < ARRAY_SIZE; i++) {
        countCards[i] = 0;
    }
    int linesSize = 0;
    int sum = 0;

    VEC_NEW(winning);
    VEC_NEW(picked);

    while (!feof(file)) {
        char *line = (char *)malloc(MAX_BUFFER_SIZE * sizeof(char));
        fgets(line, MAX_BUFFER_SIZE, file);
        int len = strlen(line);
        if (line[len - 1] == '\n') {
            line[len - 1] = '\0';
        }
        lines[linesSize] = line;
        linesSize++;
    }

    for (int i = 0; i < linesSize; i++) {
        VEC_CLR(winning);
        VEC_CLR(picked);

        char *line = lines[i];
        char *firstNumber = line + strcspn(line, "1234567890");
        char *colon = strchr(line, ':');
        int lineNumber = natoi(firstNumber, colon - firstNumber);
        char *verticalBar = strchr(colon, '|');
        char *it = colon + 2;

        while (it < verticalBar) {
            int nextNumber = strcspn(it, "1234567890");
            char *nextSpace = strchr(it + nextNumber, ' ');
            int number = natoi(it + nextNumber, nextSpace - it - nextNumber);
            VEC_ADD(winning, number);
            it = nextSpace + 1;
        }
        it = verticalBar + 2;
        while (it < line + strlen(line)) {
            int nextNumber = strcspn(it, "1234567890");
            char *nextSpace = strchr(it + nextNumber, ' ');
            if (nextSpace == NULL) {
                nextSpace = it + strlen(it);
            }
            int number = natoi(it + nextNumber, nextSpace - it - nextNumber);
            VEC_ADD(picked, number);
            it = nextSpace + 1;
        }

#ifdef DEBUG_EXEC
        printf("%d: %s\n", lineNumber, line);
        printf("\twinning %s -- ", VEC_STR(winning));
        for (int j = 0; j < VEC_LEN(winning); j++) {
            printf("%d ", VEC_GET(winning, int, j));
        }
        printf("\n\tpicked %s -- \t", VEC_STR(picked));
        for (int j = 0; j < VEC_LEN(picked); j++) {
            printf("%d ", VEC_GET(picked, int, j));
        }
        printf("\n");
#endif

        int countMatches = 0;
        for (int j = 0; j < VEC_LEN(winning); j++) {
            int number = VEC_GET(winning, int, j);
            int hasNumber = FALSE;
            for (int k = 0; k < VEC_LEN(picked); k++) {
                if (number == VEC_GET(picked, int, k)) {
                    countMatches++;
                    break;
                }
            }
        }
        countCards[lineNumber]++;
        sum += countCards[lineNumber];
        for (int k = 1; k <= countMatches; k++) {
            countCards[lineNumber + k] += countCards[lineNumber];
        }
    }

#ifdef DEBUG_EXEC
    printf("\ncard counters\n", sum);
    for (size_t i = 1; i <= linesSize; i++) {
        printf("%d ", countCards[i]);
    }
    printf("\n");
#endif

    printf("%d\n", sum);
}
