
#include "edpad.h"
#include "xkeys.h"

gboolean on_button_down(GtkButton *button, void *data) {
    char *val = NULL;
    g_object_get(button, "name", &val, NULL);
    if (!val)
        return FALSE;
    xkeys_pressed(val);
    return FALSE;
}

gboolean on_button_up(GtkButton *button, void *data) {
    char *val = NULL;
    g_object_get(button, "name", &val, NULL);
    if (!val)
        return FALSE;
    xkeys_released(val);
    return FALSE;
}

