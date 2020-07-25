#include <stdio.h>

#define START   1
#define END     0


int main() {

    int c, n1, nw, nc, state;

    state = END;

    n1 = nw = nc = 0;

    while ((c = getchar()) != EOF) {

        ++nc;

        if (c == '\n') {
            ++n1;
        }

        if (c == ' ' || c == '\n' || c == '\t') {
            state = END;
        } else if (state == END) {
            state = START;
            ++nw;
        }
    }

    printf("%d %d %d\n", n1, nw, nc);

}
