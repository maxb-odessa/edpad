
#include "edpad.h"
#include "xkeys.h"

#include <errno.h>




// main func
int main(int argc, char *argv[]) {
    GtkBuilder *builder;
    GObject *window;
    GError *error = NULL;

    gtk_init(&argc, &argv);


    /* Construct a GtkBuilder instance and load our UI description */
    builder = gtk_builder_new();
    if (gtk_builder_add_from_file(builder, "edpad.glade", &error) == 0) {
        g_printerr("Error loading file: %s\n", error->message);
        g_clear_error(&error);
        return 1;
    }

    /*
       {
    // start file reading thread

    // wait for data from pipe

    GObject *text_view1;
    GtkTextBuffer *buffer1;
    GtkTextIter iter1;
    text_view1 = gtk_builder_get_object(builder, "view1");
    buffer1 = gtk_text_view_get_buffer(GTK_TEXT_VIEW(GTK_WIDGET(text_view1)));
    gtk_text_buffer_get_iter_at_offset(buffer1, &iter1, 0);
    gtk_text_buffer_insert(buffer1, &iter1, "Plain text\n", -1);

    // ditto for buffer2/view2

    GObject *text_view3;
    GtkTextBuffer *buffer3;
    GtkTextIter iter3, iter3s, iter3e;
    text_view3 = gtk_builder_get_object(builder, "view3");
    buffer3 = gtk_text_view_get_buffer(GTK_TEXT_VIEW(GTK_WIDGET(text_view3)));
    gtk_text_buffer_get_start_iter(buffer3, &iter3s);
    gtk_text_buffer_get_end_iter(buffer3, &iter3e);
    gtk_text_buffer_delete(buffer3, &iter3s, &iter3e);
    gtk_text_buffer_get_start_iter(buffer3, &iter3);
    gtk_text_buffer_insert(buffer3, &iter3, "new", -1);
    }
    */

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
