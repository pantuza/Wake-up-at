

package main


import (
    "fmt"
    "time"
    "flag"
    "strings"
)


const MORNING string = "AM"
const EVENING string = "PM"


const MIDDAY int = 12
const TIME_FORMAT = "3:04 PM"


// Parses the command line options
func parseOptions(hours, minutes *int, period *string) {

    flagHours := flag.Int("h", 8, "Hour to wake up")
    flagMinutes := flag.Int("m", 0, "Minute to wake up in the given hour")
    flagPeriod := flag.String("p", "am", "Period of the day: am/pm")

    flag.Parse()

    *hours = *flagHours
    *minutes = *flagMinutes
    *period = strings.ToUpper(*flagPeriod)

    if *period != MORNING && *period != EVENING {
        *period = MORNING
    }
}


// Returns true if the period of the day is Morning. Otherwise, false
func isMorning (period string) bool { return period == MORNING }


// Calculates the possibles times to go to sleep
func calcTimes(wake_time, first_time, second_time,
               third_time, fourth_time *time.Time) {

    *first_time = wake_time.Add(-540 * time.Minute)
    *second_time = wake_time.Add(-450 * time.Minute)
    *third_time = wake_time.Add(-360 * time.Minute)
    *fourth_time = wake_time.Add(-270 * time.Minute)
}


// Formats the output message and print
func formatAndPrint(wake_time, first_time, second_time,
                    third_time, fourth_time *time.Time) {

    fmt.Printf("To wake up at %v, ", wake_time.Format(TIME_FORMAT))
    fmt.Printf("you should sleep at: %v\n\n", first_time.Format(TIME_FORMAT))
    fmt.Printf("Also at: %v | %v | %v\n", second_time.Format(TIME_FORMAT),
               third_time.Format(TIME_FORMAT), fourth_time.Format(TIME_FORMAT))
}


func main () {

    now := time.Now()

    // Variables for wake time calculation
    var period string
    var hours, minutes int


    // Parses command line options
    parseOptions(&hours, &minutes, &period)


    wake_time := now

    if isMorning(period) {

        // Increment wake up day to tomorrow
        if now.Hour() > MIDDAY {
            wake_time = wake_time.AddDate(0, 0, 1)
        }

    // If wake up period is PM
    } else {

        // Normalize location that uses hours like 22 instead of 10 PM.
        if hours <= MIDDAY {
            // then we increment it by 12 hours
            hours += MIDDAY
        }
    }

    // Update the wake time with correct values of hours and period
    wake_time = time.Date(wake_time.Year(), wake_time.Month(),
        wake_time.Day(), hours, minutes, 0, 0, wake_time.Location())

    // Calculate possible times
    var first_time, second_time, third_time, fourth_time time.Time
    calcTimes(&wake_time, &first_time, &second_time, &third_time, &fourth_time)

    // Prints times to go sleep
    formatAndPrint(&wake_time, &first_time, &second_time,
                   &third_time, &fourth_time)
}
