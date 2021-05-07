#! /bin/sh

nasm -f elf 1.asm -o 1.o
gcc -m32 1.o -o 1
./1 ; echo $?
