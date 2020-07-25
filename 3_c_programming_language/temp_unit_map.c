# include <stdio.h>

int main() {
    float fahr, celsius;
    int lower, upper, step;

    fahr = 0;
    upper = 300;
    step = 20;

    // use while loop
    while (fahr <= upper) {
        celsius = 5.0 / 9.0 * (fahr - 32.0);
        printf("%3.0f %6.1f\n", fahr, celsius);
        fahr = fahr + step;
    }

    // use for loop
    for (fahr = 0; fahr <= 300; fahr = fahr + 20) {
        celsius = 5.0 / 9.0 * (fahr - 32.0);
        printf("%3.0f %6.1f\n", fahr, celsius);
    }
}
