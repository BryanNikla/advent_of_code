const {eachMatrix, eachSurrounding, isLastColOfMatrix, arraySum} = require(global.UTILITIES_PATH);

const isDigit = (char) => /[0-9]/.test(char);
const isSymbol = (char) => char !== "." && !isDigit(char);

module.exports = {
    solutions: [4361, null],
    one: (input) => {
        const input_lines = input.split("\n");
        const matrix = input_lines.map((line) => line.split(""));

        let currentNumberIsPart = false;
        let currentNumber = "";
        const parts = [];

        eachMatrix(matrix, (char, [col, row]) => {
            if (isDigit(char)) {
                currentNumber = currentNumber + char;

                eachSurrounding(matrix, [col, row], (char) => {
                    if (isSymbol(char)) {
                        currentNumberIsPart = true;
                    }
                });
            }

            // If we are at the end of the matrix or the next char is not a digit
            if (isLastColOfMatrix(matrix, [col, row]) || !isDigit(char)) {
                if (currentNumberIsPart) {
                    parts.push(Number(currentNumber));
                }
                currentNumberIsPart = false;
                currentNumber = "";
            }
        });

        return arraySum(parts);
    },
    two: (input) => {
        return;
    },
};
