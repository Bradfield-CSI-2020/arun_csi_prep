#include <stdio.h>

int getLine(char s[], int lim) {

  int c, i;

  for (i = 0; (c = getchar()) != EOF && c != '\n'; ++i) {
    if (i < lim - 1) {
      s[i] = c;
    }
  }

  if (c == '\n') {
    s[lim - 2] = c;
    ++i;
  }

  s[lim - 1] = '\0';
  return i;
}