const path = require("path");
const fs = require("fs");

console.clear();

// Set global variables
global.ROOT_PATH = require("path").resolve(__dirname);
global.UTILITIES_PATH = require("path").resolve(__dirname, "common", "utilities.js");

const {zeroPad, colorText} = require(global.UTILITIES_PATH);

printGreeting();

const yearIndex = process.argv.indexOf("-year");
const dayIndex = process.argv.indexOf("-day");

if (yearIndex === -1 || dayIndex === -1) {
    const readline = require("readline").createInterface({input: process.stdin, output: process.stdout});

    readline.question(colorText("red", "What Year? "), (year) => {
        readline.question(colorText("green", "What Day? "), (day) => {
            readline.close();
            year = parseInt(year) || new Date().getFullYear();
            day = day === "all" || day === "" ? -1 : parseInt(day);

            if (!year || isNaN(year) || !day || isNaN(day)) {
                console.error("year & day arguements must be valid");
                process.exit();
            }
            execute(year, day);
        });
    });
} else {
    const year = parseInt(process.argv[yearIndex + 1]);
    const dayArg = process.argv[dayIndex + 1];
    const day = dayArg === "all" ? -1 : parseInt(process.argv[dayIndex + 1]);
    if (!year || isNaN(year) || !day || isNaN(day)) {
        console.error("year & day arguements must be valid");
        process.exit();
    }
    execute(year, day);
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

function printGreeting() {
    console.log("\n");
    console.log(colorText("green", "🎄  Advent of Code  🎅"));
    console.log("-".repeat(process.stdout.columns));
    console.log(colorText("cyan", "https://adventofcode.com/"));
    console.log(colorText("cyan", "Code By: Bryan Nika"));
    console.log("\n");
}

function execute(year, day) {
    //////////////////////////////////////////////////////////////////////////////////////
    const output = {};
    if (day === -1) {
        for (let i = 1; i < 26; i++) {
            output[parseFloat(year + "." + zeroPad(i, 2)).toFixed(2)] = _solve(year, i);
        }
    } else {
        output[parseFloat(year + "." + zeroPad(day, 2)).toFixed(2)] = _solve(year, day);
    }

    // Final output
    console.log("\n");
    console.table(output);
    //////////////////////////////////////////////////////////////////////////////////////

    function _solve(year, day) {
        function _getFile(path) {
            return fs.readFileSync(path, "utf8");
        }

        function _handleError(path) {
            return (err) => {
                if (String(err).includes(path)) {
                    return rsp;
                }
                console.error(err);
                process.exit();
            };
        }

        const solution_path = path.resolve(__dirname, String(year), zeroPad(day, 2), "solution.js");
        const input_path = path.resolve(__dirname, String(year), zeroPad(day, 2), "input.txt");
        const test_path = path.resolve(__dirname, String(year), zeroPad(day, 2), "test.txt");

        const rsp = {};

        try {
            let {one, two, solutions} = require(solution_path);

            try {
                Object.assign(rsp, {"Solution 1": one(_getFile(input_path))});
            } catch (err) {
                return _handleError(input_path)(err);
            }

            try {
                Object.assign(rsp, {"Solution 2": two(_getFile(input_path))});
            } catch (err) {
                return _handleError(input_path)(err);
            }

            try {
                Object.assign(rsp, {"Test 1": one(_getFile(test_path)) === solutions[0] ? "✅" : "❌"});
            } catch (err) {
                return _handleError(test_path)(err);
            }

            try {
                Object.assign(rsp, {"Test 2": two(_getFile(test_path)) === solutions[1] ? "✅" : "❌"});
            } catch (err) {
                return _handleError(test_path)(err);
            }
        } catch (err) {
            return _handleError(solution_path)(err);
        }

        return rsp;
    }
}
