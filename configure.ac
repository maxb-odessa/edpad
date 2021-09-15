#                                               -*- Autoconf -*-
# Process this file with autoconf to produce a configure script.

AC_PREREQ([2.69])
AC_INIT([edpad], [0.0.1], [maxb@tutamail.com])
AC_CONFIG_SRCDIR([src/edpad.c])
AC_CONFIG_HEADERS([src/config.h])
AC_CONFIG_FILES([Makefile src/Makefile])
AM_INIT_AUTOMAKE

# Checks for programs.
CFLAGS=" -std=gnu11 -ggdb -Wall"
AC_PROG_CC

# Checks for libraries.

# Checks for header files.
AC_CHECK_HEADERS([fcntl.h string.h unistd.h])

# Checks for typedefs, structures, and compiler characteristics.
AC_TYPE_SSIZE_T
AC_CHECK_HEADER_STDBOOL

# Checks for library functions.
AC_CHECK_FUNCS([memset])

PKG_CHECK_MODULES([GTK3], [gtk+-3.0])
PKG_CHECK_MODULES([GME2], [gmodule-export-2.0])


#AC_CHECK_HEADERS([xdo.h])
#AC_CHECK_LIB([xdo], [xdo_new])

AC_OUTPUT
