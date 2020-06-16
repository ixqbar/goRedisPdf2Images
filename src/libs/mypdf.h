#ifndef _MYPDF_H
#define _MYPDF_H

int mypdf_size(const char * filename);
int mypdf_parse(const char * filename, int start, int end);
int png_compress(const char *filename);

#endif