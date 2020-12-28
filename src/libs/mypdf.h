#ifndef _MYPDF_H
#define _MYPDF_H

int mypdf_size(const char * filename);
int mypdf_parse(const char * filename, int zoom, int start, int end, int compress);
int png_compress(const char *filename);

#endif