TARGET=pdf2png

cc=gcc
CFLAGS=-I. -g -Wall -std=c99
LIBS=/usr/local/lib/libimagequant.a /usr/local/lib/libmupdf.a /usr/local/lib/libmupdf-third.a -lm

all: mac

mac:
	$(CC) -o $(TARGET) pdf2png.c mypdf.c $(CFLAGS) $(LIBS)
	$(CC) -o pngcompress pngcompress.c mypdf.c $(CFLAGS) $(LIBS)

clean:
	rm -rf $(TARGET)