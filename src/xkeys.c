
#include <stdio.h>
#include <stdlib.h>
#include <xdo.h>

#include "xkeys.h"

static xdo_t *xdo;

int xkeys_init(int argc, char **argv) {
    char *xserver;

    if (argc > 1)
        xserver = argv[1];
    else 
        xserver = getenv("EDPAD_DISPLAY");

    if (! xserver) {
        fprintf(stderr, "Either specify remote X server as an argument or set env var 'EDPAD_DISPLAY' accordingly\n");
        return -1;
    }

    xdo = xdo_new(xserver);
    if (! xdo) {
        fprintf(stderr, "Failed to connect to remote display\n");
        return -1;
    }

    return 0;
}


void xkeys_pressed(const char *key) {
    xdo_send_keysequence_window_down(xdo, CURRENTWINDOW, key, 1);
}

void xkeys_released(const char *key) {
    xdo_send_keysequence_window_up(xdo, CURRENTWINDOW, key, 1);
}

