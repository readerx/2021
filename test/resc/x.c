#include <stdio.h>

void hanoi(int n, char *a, char *b, char *c) {
    if (n == 1) {
        printf("%s -> %s\n", a, c);
    } else {
        hanoi(n - 1, a, c, b);
        printf("%s -> %s\n", a, c);
        hanoi(n - 1, b, a, c);
    }
}

int main(int argc, char *argv[]) {
    hanoi(2, "a", "b", "c");
    return 0;
}
