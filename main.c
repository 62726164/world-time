#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>
#include <stdbool.h>

#define NUM_TZ 26
#define TZ_LEN 50

const char * tz[NUM_TZ] = {
    "UTC",
    "Local",
    "Atlantic/Reykjavik",
    "Atlantic/Cape_Verde",
    "America/Noronha",
    "America/Buenos_Aires",
    "America/New_York",
    "America/Chicago",
    "America/Denver",
    "America/Los_Angeles",
    "America/Anchorage",
    "America/Adak",
    "Pacific/Honolulu",
    "Pacific/Midway",
    "Pacific/Wake",
    "Pacific/Guadalcanal",
    "Pacific/Guam",
    "Asia/Tokyo",
    "Asia/Shanghai",
    "Asia/Saigon",
    "Antarctica/Vostok",
    "Asia/Ashkhabad",
    "Asia/Dubai",
    "Europe/Moscow",
    "Africa/Johannesburg",
    "Europe/London"
};

static bool print_date_for_TZ(char *);

// Entry point for the program
int main(int argc, char ** argv) {
    for (int i = 0; i < NUM_TZ; i++) {
        print_date_for_TZ((char *) tz[i]);
    }

    printf("\n");

    // Loop forever until SIGINT
    while (true) {
        char *line = NULL;
        size_t len = 0;

        printf("Please enter a timezone. Press CTRL-C to exit: ");
        getline(&line, &len, stdin);
        printf("Getting results for: ");

        /* Bug #3: Print format vulnerability. This can be used to
         * print locations in the stack to jump to! */
        printf(line);

        bool success = print_date_for_TZ(line);
        if (!success) {
            printf("ERROR: Timezone invalid. Please try again.\n");
        }
    }

    return 0;
}

static bool print_date_for_TZ(char * my_tz) {
    // Check if my_tz is a valid timezone

    bool correct = false;
    for (int i = 0; i < NUM_TZ; i++) {
        /* Bug #1: We are not comparing the full string. What
         * happens if my_tz is longer than tz[i]? */
        if(strncmp(my_tz, tz[i], strlen(tz[i])) == 0)
            correct = true;
    }

    if (!correct)
        return false;

    /* Timezones shouldn't be above 50 characters, right? */
    /* Bug #2: Buffer overflow */
    char tz_str[TZ_LEN];
    strcpy(tz_str, my_tz);

    /* Set the env var */
    setenv("TZ", my_tz, 1);

    /* Get local time */
    struct tm * timeinfo;
    time_t raw_time;
    time(&raw_time);
    timeinfo = localtime(&raw_time);

    /* Print to the user */
    printf("%04d-%02d-%02d %02d:%02d:%02d %s\n", 1900 + timeinfo->tm_year,
            1 + timeinfo->tm_mon, timeinfo->tm_mday, timeinfo->tm_hour, 
            timeinfo->tm_min, timeinfo->tm_sec, my_tz);
    return true;
}
