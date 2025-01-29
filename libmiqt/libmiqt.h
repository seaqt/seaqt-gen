#pragma once
#ifndef MIQT_LIBMIQT_LIBMIQT_H
#define MIQT_LIBMIQT_LIBMIQT_H

#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus

namespace seeqt {
struct release_callback {
  void (*release)(intptr_t);

  release_callback(void (*release)(intptr_t)) : release(release) {}
  release_callback(release_callback &&rhs) : release(rhs.release) {
    rhs.release = 0;
  }
  release_callback &operator=(release_callback &&rhs) {
    release = rhs.release;
    rhs.release = 0;
    return *this;
  };

  void operator()(intptr_t slot) {
    if (release)
      release(slot);
  }
  release_callback(const release_callback &) = delete;
  release_callback &operator=(const release_callback &) = delete;
};
} // namespace seeqt

extern "C" {
#endif

struct miqt_string {
  size_t len;
  char *data;
};

struct miqt_array {
  size_t len;
  void *data;
};

struct miqt_map {
  size_t len;
  void *keys;
  void *values;
};

typedef const char const_char;

#ifdef __cplusplus
}
#endif

#endif
