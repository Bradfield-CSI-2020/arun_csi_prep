#include <stdio.h>

#define MAXLINE 10
#define PADDING 2

int getLine(char line[], int maxline);
void copy(char to[], char from[]);

int main() {
  int len;
  int max = 0;

  int size = MAXLINE + PADDING;

  char line[size];
  char longest[size];

  while ((len = getLine(line, size)) > 0) {
    if (len > max) {
      max = len;
      copy(longest, line);
    }
  }

  if (max > 0) {
    printf("\nLength of longest line: \n");
    printf("%d\n", max);
    printf("Longest line: \n");
    printf("%s\n", longest);
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

void copy(char to[], char from[]) {
  int i;

  i = 0;

  while ((to[i] = from[i]) != '\0') {
    ++i;
  }
}
