#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>

#include <stdio.h>
#include <stdlib.h>
#include "ai_geekeye.h"

static int 
ai_lua_geekeye(lua_State *L) {
    char        *img, *val = NULL;
    int          scope;

    if (lua_gettop(L) != 2) {
        return luaL_error(L, "expecting two argument");
    }
    scope = lua_tointeger(L, -2);
    //img = (unsigned char*) luaL_checklstring(L, -1, &len);
    img = (char*) lua_tostring(L, -1);

    ai_geekeye(scope, img, &val);
    fprintf(stderr, "val: %s\n", val);
    lua_pushinteger(L, 10);
    lua_pushstring(L, val);
    //lua_pushnumber(L, 10);
    // lua_checkstack(L)
    free(val);
    val = NULL;

    return 2;
}

static const struct luaL_Reg ai_lib[] = {
    {"geekeye", ai_lua_geekeye},
    {NULL, NULL},
};

int luaopen_ai(lua_State *L) {
    luaL_register(L, "ai", ai_lib);
    return 1;
}
