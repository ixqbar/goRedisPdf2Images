#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
#include "mypdf.h"

int main(int argc, const char *argv[])
{
    if (argc < 2) {
        printf("Usage like %s xxx.pdf", argv[0]);
        return 1;
    }

    int start = 0;
    int end = 0;

    if (argc >= 3) {
        long startL = strtol(argv[2], NULL, 10);
        if (startL > INT_MIN  && startL < INT_MAX) {
            start = startL;
        }
    }

    if (argc >= 4) {
        long endL = strtol(argv[3], NULL, 10);
        if (endL > INT_MIN  && endL < INT_MAX) {
            end = endL;
        }
    }

    mypdf_parse(argv[1],start, end);

    return 0;
}