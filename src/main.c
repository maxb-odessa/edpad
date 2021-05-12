#include <gtk/gtk.h>
#include <xdo.h>

static void gal_map(GtkWidget * widget, gpointer data) {
    xdo_t *x = xdo_new("192.168.100.100:0.0");

    xdo_send_keysequence_window(x, CURRENTWINDOW, "m", 1);
    xdo_free(x);
}

static void sys_map(GtkWidget * widget, gpointer data) {
    xdo_t *x = xdo_new("192.168.100.100:0.0");

    xdo_send_keysequence_window(x, CURRENTWINDOW, "comma", 1);
    xdo_free(x);
}

static void destroy(GtkWidget * widget, gpointer data) {
    g_print("X pressed\n");
    gtk_main_quit();
}


static void activate(GtkApplication * app, gpointer user_data) {
    GtkWidget *window;
    GtkWidget *button;
    GtkWidget *button_box;

    window = gtk_application_window_new(app);
    gtk_window_set_title(GTK_WINDOW(window), "Window");
    gtk_window_fullscreen(GTK_WINDOW(window));
    g_signal_connect(window, "destroy", G_CALLBACK(destroy), NULL);

    button_box = gtk_button_box_new(GTK_ORIENTATION_HORIZONTAL);
    gtk_container_add(GTK_CONTAINER(window), button_box);

    // M
    button = gtk_button_new_with_label("Open Galaxy Map");
    g_signal_connect(button, "clicked", G_CALLBACK(gal_map), NULL);
    //    g_signal_connect_swapped (button, "clicked", G_CALLBACK (gtk_widget_destroy), window);
    gtk_container_add(GTK_CONTAINER(button_box), button);

    // L
    button = gtk_button_new_with_label("Open System Map");
    g_signal_connect(button, "clicked", G_CALLBACK(sys_map), NULL);
    //    g_signal_connect_swapped (button, "clicked", G_CALLBACK (gtk_widget_destroy), window);
    gtk_container_add(GTK_CONTAINER(button_box), button);


    gtk_widget_show_all(window);
}


int main(int argc, char **argv) {
    GtkApplication *app;
    int status;

    app = gtk_application_new("org.gtk.example", G_APPLICATION_FLAGS_NONE);
    g_signal_connect(app, "activate", G_CALLBACK(activate), NULL);
    status = g_application_run(G_APPLICATION(app), argc, argv);
    g_object_unref(app);

    return status;
}
