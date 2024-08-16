module.exports = {one, two, solutions: [288, 71503]};

const {arrOfLength} = require(global.UTILITIES_PATH);

function parseInput(input, partTwo = false) {
    const [t, d] = input.split("\n").map((line) => {
        line = line
            .split(":")[1]
            .trim()
            .replace(/\s{2,}/g, " ");
        if (partTwo) {
            line = line.replaceAll(" ", "");
        }
        return line.split(" ").map(Number);
    });
    return t.map((x, i) => ({time: x, distance: d[i]}));
}

// TODO: This is just brute forcing this solution. There's probably a better way to do it.
function calculateWinningMethods({time, distance}) {
    return arrOfLength(time).filter((t) => {
        return t * (time - t) > distance;
    }).length;
}

function one(input) {
    return parseInput(input, false)
        .map(calculateWinningMethods)
        .reduce((acc, val) => acc * val, 1);
}

function two(input) {
    return parseInput(input, true)
        .map(calculateWinningMethods)
        .reduce((acc, val) => acc * val, 1);
}
