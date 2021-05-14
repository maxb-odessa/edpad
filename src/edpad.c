

#include "edpad.h"


gboolean mode_analyzer_map_cb(GtkWidget *widget, GdkEvent  *event, gpointer   user_data) {
    gtk_widget_show(widget);
    g_message("%s", gtk_widget_get_name(widget));
    g_message("switch1!");
    return FALSE;
}

gboolean mode_reserved_map_cb(GtkWidget *widget, GdkEvent  *event, gpointer   user_data) {
    gtk_widget_show(widget);
    g_message("%s", gtk_widget_get_name(widget));
    g_message("switch2!");
    return FALSE;
}

gboolean mode_rover_map_cb(GtkWidget *widget, GdkEvent  *event, gpointer   user_data) {
    gtk_widget_show(widget);
    g_message("%s", gtk_widget_get_name(widget));
    g_message("switch3!");
    return FALSE;
}

gboolean mode_combat_map_cb(GtkWidget *widget, GdkEvent  *event, gpointer   user_data) {
    gtk_widget_show(widget);
    g_message("%s", gtk_widget_get_name(widget));
    g_message("switch4!");
    return FALSE;
}

gboolean mode_landing_map_cb(GtkWidget *widget, GdkEvent  *event, gpointer   user_data) {
    gtk_widget_show(widget);
    g_message("%s", gtk_widget_get_name(widget));
    g_message("switch5!");
    return FALSE;
}

gboolean quit_button_clicked_cb(GtkWidget *widget, GdkEvent  *event, gpointer   user_data) {
    gtk_widget_show(widget);
    g_message("button");
    return TRUE;
}

int main(int argc, char *argv[]) {
    GtkBuilder *builder;
    GObject *window;
    GError *error = NULL;

    gtk_init(&argc, &argv);

    /* Construct a GtkBuilder instance and load our UI description */
    builder = gtk_builder_new();
    if (gtk_builder_add_from_file(builder, "edpad.ui", &error) == 0) {
        g_printerr("Error loading file: %s\n", error->message);
        g_clear_error(&error);
        return 1;
    }

    /* Connect signal handlers to the constructed widgets. */
    gtk_builder_connect_signals(builder, NULL);
    window = gtk_builder_get_object(builder, "window");
    g_signal_connect(window, "destroy", G_CALLBACK(gtk_main_quit), NULL);
    //GtkWindow *w = GTK_WINDOW(window);
    //gtk_window_fullscreen(w);
    //gtk_window_maximize(w);
    gtk_widget_show_all(GTK_WIDGET(window));
    gtk_main();

    return 0;
}
