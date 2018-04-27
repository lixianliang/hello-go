
all:
	gcc ai_geekeye.c -Wall -shared -fPIC -o libgeekeye.so
	gcc lua_ai.c -Wall -shared -fPIC -o ai.so  -L./ -lgeekeye


