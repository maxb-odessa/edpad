
#include "edpad.h"

// show auit dialog window
gboolean on_quit_button_clicked(GtkButton *button, GtkWidget *popup) {
    gtk_widget_show(popup);
    return FALSE;
}


// hide quit dialog window
gboolean on_quit_no_clicked(GtkButton *button, GtkWidget *popup) {
    gtk_widget_hide(popup);
    return FALSE;
}
