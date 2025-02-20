#go-brawl-stats

It is a CLI program that requests data from [Supercell API](https://developer.brawlstars.com),
save as JSON, parse and create statistics.

## Contents

1. [Why?](#why?)
2. [How to run](#how-to-run)
3. [How it works](#how-it-works)
4. [Settings](#settings)
5. [License](#license-mit)

## Why?
I'm a lawyer in transition to a dev career,
and a senior Information Systems Development student,
with no job experience in the area.
Listening/reading all the tips for "How to get experience" or
"How to learn code", I created this project.

Therefore, the main purpose of this project for study is (but not limited to):

    * API client
    * CLI application
    * Go standard lib
    * Data structures
      * Map
        * map[string]int
        * map[string]struct
        * map[string]map[string]string
      * Slices/Arrays
      * []Byte
      * Structs
        * Nested structs
        * []struct
        * Struct Methods
    * Pointers
    * JSON
      * Read a JSON response
      * Read a JSON file
      * Parse []Byte to JSON
      * Parse JSON to []Byte
      * Save JSON file
    * Data manipulation
    * Abstraction
      * to build Requests
      * to Open/Read file
      * to Save in a file


## How to run
Clone this repo.
Create a .env file with the your variable APIKEY (from Supercell API)
```sh
go run . PLAYERTAG
```

## How it works
It takes the PlayerTag, request to Supercell API the player info and all brawlers.
Save to a json file.
Parse the json file and show the statistics to max level.

## Settings
TODO


## License MIT
Project License can be found [here](LICENSE).
