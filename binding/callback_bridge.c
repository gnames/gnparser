#include "callback_bridge.h"

void callback_bridge(void *callback, char *parsed) {
  ((Callback*) callback)(parsed);
};
