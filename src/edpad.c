
#include "edpad.h"
#include "xkeys.h"

#include <errno.h>

// main func
int main(int argc, char *argv[]) {
    GtkBuilder *builder;
    GObject *window;
    GError *error = NULL;

    gtk_init(&argc, &argv);

    if (xkeys_init(argc, argv) != 0) {
        return 1;
    }


    /* Construct a GtkBuilder instance and load our UI description */
    builder = gtk_builder_new();
    if (gtk_builder_add_from_file(builder, "edpad.ui", &error) == 0) {
        g_printerr("Error loading file: %s\n", error->message);
        g_clear_error(&error);
        return 1;
    }

    /* Connect signal handlers to the constructed widgets. */
    gtk_builder_connect_signals(builder, (void *)builder);
    window = gtk_builder_get_object(builder, "window");
    g_signal_connect(window, "destroy", G_CALLBACK(gtk_main_quit), NULL);
    //GtkWindow *w = GTK_WINDOW(window);
    //gtk_window_fullscreen(w);
    //gtk_window_maximize(w);
    gtk_widget_show_all(GTK_WIDGET(window));
    gtk_main();

    return 0;
}
