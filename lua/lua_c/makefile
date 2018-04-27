
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./

all:
	gcc test.c -Wall -shared -fPIC -o libtest.so
	gcc hello.c -Wall -shared -fPIC -o hello.so  -L./ -ltest


# gcc -c test.c -o test0
# ar crv libtest.a test.o     
# gcc -static hello.c -Wall -shared -fPIC -o hello.so -L./ -ltest
# gcc -static hello.c -Wall -shared -fPIC -o hello.so libtest.a
