#include "seaqt/QtWidgets/gen_qabstractbutton.h"
#include "seaqt/QtWidgets/gen_qapplication.h"
#include "seaqt/QtWidgets/gen_qpushbutton.h"
#include "seaqt/QtWidgets/gen_qwidget.h"
#include "seaqt/libseaqt-runtime.h"

#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct seaqt_string seaqt_literal(const char *s) {
  return (struct seaqt_string){strlen(s), (char *)s};
}

struct HelloWorld {
  VirtualQPushButton *button;
  int counter;
};

static QApplication_VTable a_vtbl = {};
static QPushButton_VTable pb_vtbl = {};

void onPressed(intptr_t data) {
  struct HelloWorld *env = (struct HelloWorld *)data;
  env->counter += 1;

  const char *fmt = "You have clicked the button %d time(s)";

  int bytes = snprintf(NULL, 0, fmt, env->counter);
  char *str = (char *)malloc(bytes);
  sprintf(str, fmt, env->counter);

  QAbstractButton_setText((QAbstractButton *)env->button, seaqt_literal(str));

  free(str);
}

int main(int argc, char **argv) {
  QApplication_new_int_char(&a_vtbl, 0, &argc, argv);

  struct HelloWorld env = {.button = QPushButton_new_text(
                               &pb_vtbl, 0, seaqt_literal("Hello world!"))};

  QWidget_setFixedWidth((QWidget *)env.button, 320);

  QAbstractButton_connect_pressed((QAbstractButton *)env.button, (intptr_t)&env,
                                  onPressed, NULL);

  QWidget_show((QWidget *)env.button);

  return QApplication_exec();
}
