#pragma once

#define TRUE 1
#define FALSE 0

#define DEBUG_EXEC
#undef DEBUG_EXEC

#define MAX_BUFFER_SIZE 1024

int natoi(char *str, int len);
int countnchar(char *str, int len, char c);

int isNumber(char c);