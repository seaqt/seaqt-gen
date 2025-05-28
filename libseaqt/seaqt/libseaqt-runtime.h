#pragma once
#ifndef MIQT_LIBMIQT_LIBMIQT_H
#define MIQT_LIBMIQT_LIBMIQT_H

#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

struct seaqt_string {
  size_t len;
  char *data;
};

struct seaqt_array {
  size_t len;
  void *data;
};

struct seaqt_map {
  size_t len;
  void *keys;
  void *values;
};

#ifdef __cplusplus
}
#endif

#endif
