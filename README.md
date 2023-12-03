# Advent of Code

## To Run Code

Execute code using NodeJS:

`node index.js`

This will prompt you for the year & day.

| Question | Options                                               |
| -------- | ----------------------------------------------------- |
| Year?    | 2022, 2023, or no answer for current year             |
| Day?     | 1-25, 01-25, all, or no answer will also run all days |

---

### Optional arguments:

| Argument  | Options                                                       |
| --------- | ------------------------------------------------------------- |
| -year     | 2022, 2023, all                                               |
| -day      | 1-25, 01-25, all                                              |
| -test     | Will only run tests for the given year/day                    |
| -generate | Will generate directory & files needed for the given year/day |

#### Examples:

`node index.js -year 2023 -day 01`

`node index.js -year 2023 -day 3 -test`

`node index.js -year 2022 -day 1 -generate`
