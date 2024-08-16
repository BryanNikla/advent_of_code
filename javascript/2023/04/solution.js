// 2023/04
module.exports = {one, two, solutions: [13, 30]};

const {arraySum, forX} = require(global.UTILITIES_PATH);

function parseInputIntoCardValues(inputLines) {
    const _getNumbers = (str) => {
        return str
            .trim()
            .replaceAll("  ", " ")
            .split(" ")
            .map((n) => Number(n));
    };
    return inputLines.map((line) => {
        const [winners, numbers] = line.split(":")[1].split("|");
        return [_getNumbers(winners), _getNumbers(numbers)];
    });
}

function wins(winners, numbers) {
    return winners.filter((value) => numbers.includes(value)).length;
}

function one(input) {
    return arraySum(
        parseInputIntoCardValues(input.split("\n")).map(([winners, numbers]) => {
            const count = wins(winners, numbers);
            return count <= 0 ? 0 : 1 * Math.pow(2, count - 1);
        })
    );
}

function two(input) {
    const cards = parseInputIntoCardValues(input.split("\n"));
    const count = cards.map(() => 1);
    cards.forEach(([w, n], i) => {
        forX(wins(w, n), (x) => {
            count[i + x] += count[i];
        });
    });
    return arraySum(count);
}
