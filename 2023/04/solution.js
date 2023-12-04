// 2023/04
module.exports = {one, two, solutions: [13, 30]};

const {arraySum} = require(global.UTILITIES_PATH);

function parseInputIntoCardValues(inputLines) {
    const _getNumbers = (str) => {
        return str
            .trim()
            .replaceAll("  ", " ")
            .split(" ")
            .map((n) => Number(n));
    };
    return inputLines.map((line) => {
        const [_, val] = line.split(":");
        const [winners, numbers] = val.split("|");
        return [_getNumbers(winners), _getNumbers(numbers)];
    });
}

function numberOfWinningNumbers(winners, numbers) {
    return winners.filter((value) => numbers.includes(value)).length;
}

function one(input) {
    return arraySum(
        parseInputIntoCardValues(input.split("\n")).map(([winners, numbers]) => {
            const count = numberOfWinningNumbers(winners, numbers);
            return count <= 0 ? 0 : 1 * Math.pow(2, count - 1);
        })
    );
}

function two(input) {
    const cards = parseInputIntoCardValues(input.split("\n"));
    const pilesByIndex = cards.map(([winners, numbers]) => [[winners, numbers]]);
    const cardsInPilesAtEnd = pilesByIndex.map((pile, i) => {
        const cardsInPile = pile.length;
        // TODO: Redo this solve later to be more performant.. this is TERRIBLY slow to process (but works)
        while (pile.length) {
            const [winners, numbers] = pile.pop();
            for (let w = numberOfWinningNumbers(winners, numbers); w > 0; w--) {
                if (pilesByIndex[i + w]) {
                    pilesByIndex[i + w].push(structuredClone(pilesByIndex[i + w][0]));
                }
            }
        }
        return cardsInPile;
    });
    return arraySum(cardsInPilesAtEnd);
}
