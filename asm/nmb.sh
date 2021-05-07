#! /bin/sh

nasm -f elf nmb.asm -o nmb.o
gcc -m32 nmb.o -o nmb
./nmb ; echo $?
