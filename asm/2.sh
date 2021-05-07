#! /bin/sh

nasm -f elf 2.asm -o 2.o
gcc -m32 2.o -o 2
./2 ; echo $?
