# include <stdio.h>

int main() {
    int fahr, celsius;
    int lower, upper, step;


    fahr = 0;
    upper = 300;
    step = 20;

    while (fahr <= upper) {
        celsius = 5 * (fahr - 32) / 9;
        printf("%d\t%d\n", fahr, celsius);
        fahr = fahr + step;
    }

}