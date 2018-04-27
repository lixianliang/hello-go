
all: libfoo.so

libfoo.so:
	gcc foo.c -Wall -shared -fPIC -o libfoo.so

clean:
	rm -f libfoo.so


export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./

lua:
	~/project/nginx/LuaJIT-2.1-20160517/bin/luajit ffi.lua  ffi.lua
