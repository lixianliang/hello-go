#! /bin/sh

nasm -f elf test02.asm -o test02.o
gcc -m32 -g test02.o -o test02
./test02 ; echo $?
