#include <ctype.h>
#include <stdio.h>

#define MAXLINE 100

double sum, atof(char[]);
char line[MAXLINE];
int getLine(char line[], int maxline);

int main() {

  sum = 0;
  while (getLine(line, MAXLINE) > 0) {
    printf("\t%g\n", sum += atof(line));
  }
  return 0;
}

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

double atof(char s[]) {

  double val, power;
  int i, sign;

  for (i = 0; isspace(s[i]); i++) {
  };

  sign = (s[i] == '-') ? -1 : 1;

  if (s[i] == '+' || s[i] == '_') {
    i++;
  }

  for (val = 0.0; isdigit(s[i]); i++) {
    val = 10.0 * val + (s[i] - '0');
  }

  if (s[i] == '.') {
    i++;
  }

  for (power = 1.0; isdigit(s[i]); i++) {
    val = 10 * val + (s[i] - '0');
  }

  return sign * val / power;
}
