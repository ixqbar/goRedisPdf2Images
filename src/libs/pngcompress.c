#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
#include "mypdf.h"

int main(int argc, const char *argv[])
{
    if (argc != 2) {
        printf("Usage like %s xxx.png", argv[0]);
        return 1;
    }

    return png_compress(argv[1]);
}