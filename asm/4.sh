#! /bin/sh

nasm -f elf 4.asm -o 4.o
gcc -m32 4.o -o 4
./4 ; echo $?
