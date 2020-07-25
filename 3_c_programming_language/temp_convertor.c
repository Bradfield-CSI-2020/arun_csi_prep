# include <stdio.h>

#define LOWER 0     /* lower limit of temp */
#define UPPER 300   /* upper limit of temp */
#define STEP 20     /* step size */

int main() {
    float fahr, celsius;

    // use for loop
    for (fahr = 0.0; fahr <= 300; fahr = fahr + 20) {
        celsius = 5.0 / 9.0 * (fahr - 32.0);
        printf("%3.0f %6.1f\n", fahr, celsius);
    }
}
