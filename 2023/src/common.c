#include "common.h"

int natoi(char *str, int len) {
    int num = 0;
    for (int i = 0; i < len; i++) {
        num *= 10;
        num += str[i] - '0';
    }
    return num;
}

int countnchar(char *str, int len, char c) {
    int count = 0;
    for (int i = 0; i < len; i++) {
        if (str[i] == c) {
            count++;
        }
    }
    return count;
}

int isNumber(char c) {
    return c >= '0' && c <= '9';
}