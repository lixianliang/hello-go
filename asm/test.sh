#! /bin/sh

#nasm -f elf test.asm -o test.o ; gcc -m32 test.o -o test
nasm -f elf test.asm -o test.o ; gcc -m32 -g test.o -o test
