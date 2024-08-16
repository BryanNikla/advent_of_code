const path = require("path");
const fs = require("fs");

console.clear();

// Set global variables
global.ROOT_PATH = require("path").resolve(__dirname);
global.UTILITIES_PATH = require("path").resolve(__dirname, "javascript", "common", "utilities.js");
global.ARRAY_FILTERS_PATH = require("path").resolve(__dirname, "javascript", "common", "arrayFilters.js");
global.MATRIX_UTILITIES_PATH = require("path").resolve(__dirname, "javascript", "common", "matrixUtilities.js");
global.GENERAL_UTILITIES_PATH = require("path").resolve(__dirname, "javascript", "common", "generalUtilities.js");

const {zeroPad, isBetween} = require(global.UTILITIES_PATH);
const {colorText} = require(global.GENERAL_UTILITIES_PATH);

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

    if (isBetween(dayArg, 1, 25) && process.argv.includes("-generate")) {
        generateDay(year, day);
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

        const solution_path = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "solution.js");
        const input_path = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "input.txt");
        const test1_path = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "test1.txt");
        const test2_path = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "test2.txt");

        const rsp = {};

        try {
            let {one, two, solutions} = require(solution_path);

            if (!process.argv.includes("-test")) {
                try {
                    Object.assign(rsp, {"Solution 1": one(_getFile(input_path), {isTest: false})});
                } catch (err) {
                    return _handleError(input_path)(err);
                }

                try {
                    Object.assign(rsp, {"Solution 2": two(_getFile(input_path), {isTest: false})});
                } catch (err) {
                    return _handleError(input_path)(err);
                }
            }

            try {
                Object.assign(rsp, {"Test 1": one(_getFile(test1_path), {isTest: true}) === solutions[0] ? "✅" : "❌"});
            } catch (err) {
                return _handleError(test1_path)(err);
            }

            try {
                Object.assign(rsp, {"Test 2": two(_getFile(test2_path), {isTest: true}) === solutions[1] ? "✅" : "❌"});
            } catch (err) {
                return _handleError(test2_path)(err);
            }
        } catch (err) {
            return _handleError(solution_path)(err);
        }

        return rsp;
    }
}

// Generate the day workspace
function generateDay(year, day) {
    console.log(colorText("yellow", `Generating day ${day} for year ${year}`));

    const yearPath = path.resolve(__dirname, "javascript", String(year));
    const dayPath = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2));

    if (!fs.existsSync(yearPath)) {
        fs.mkdirSync(yearPath);
    }

    if (!fs.existsSync(dayPath)) {
        fs.mkdirSync(dayPath);
    }

    const solutionPath = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "solution.js");
    const inputPath = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "input.txt");
    const test1Path = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "test1.txt");
    const test2Path = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "test2.txt");
    const readmePath = path.resolve(__dirname, "javascript", String(year), zeroPad(day, 2), "README.md");

    if (!fs.existsSync(solutionPath)) {
        fs.writeFileSync(
            solutionPath,
            `// ${year}/${zeroPad(day, 2)}\n` +
                `module.exports = {one, two, solutions: [0, 0]};\n` +
                `\n` +
                `function one(input) {\n` +
                `    return 0;\n` +
                `}\n` +
                `\n` +
                `function two(input) {\n` +
                `    return 0;\n` +
                `}\n`
        );
    }

    if (!fs.existsSync(inputPath)) {
        fs.writeFileSync(inputPath, "");
    }

    if (!fs.existsSync(test1Path)) {
        fs.writeFileSync(test1Path, "");
    }

    if (!fs.existsSync(test2Path)) {
        fs.writeFileSync(test2Path, "");
    }

    if (!fs.existsSync(readmePath)) {
        fs.writeFileSync(readmePath, "");
    }
}
