#include <stdio.h>

#define LOWER 0   /* lower limit of temp */
#define UPPER 300 /* upper limit of temp */
#define STEP 20   /* step size */

float fhrToCelsius(float f);

int main() {
  for (float fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP) {
    printf("%3.0f %6.1f\n", fahr, fhrToCelsius(fahr));
  }
}

float fhrToCelsius(float f) { return 5.0 / 9.0 * (f - 32.0); }
