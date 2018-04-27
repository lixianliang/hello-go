
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int ai_geekeye(int scope, char *img, char **val) {

    fprintf(stderr, "scope:%d img: %s\n", scope, img);
    
    *val = calloc(1, 10);
    memcpy(*val, "123456789", 9);

    return 0;
}
