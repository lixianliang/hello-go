#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>

#include <stdio.h>
#include "test.h"

static int 
lua_hello(lua_State *l) {
    //printf("hello world.\n");
    test();
    lua_pushnumber(l, 10);
    return 1;
}

static const struct luaL_Reg hello_lib[] = {
    {"hello", lua_hello},
    {NULL, NULL},
};

/*int luaopen_luahello(lua_State *l) {
    //luaL_newlib(l, hello_lib);
    luaL_setfuncs(l, hello_lib, 0);
    return 1;
}*/

int luaopen_hello(lua_State *l) {
    luaL_register(l, "hello", hello_lib);
    return 1;
}
