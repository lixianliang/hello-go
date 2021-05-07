#! /bin/sh

nasm -f elf 3.asm -o 3.o
gcc -m32 3.o -o 3
./3 ; echo $?
