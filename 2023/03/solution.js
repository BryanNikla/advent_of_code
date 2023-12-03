const {eachMatrix, eachSurrounding, isLastColOfMatrix, arraySum} = require(global.UTILITIES_PATH);

const isDigit = (char) => /[0-9]/.test(char);
const isSymbol = (char) => char !== "." && !isDigit(char);

module.exports = {
    solutions: [4361, 467835],
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
        const input_lines = input.split("\n");
        const matrix = input_lines.map((line) => line.split(""));

        const isGearSymbol = (char) => char === "*";

        let currentNumberIsPart = false;
        let currentNumber = "";
        let currentNumberCords = [];
        const parts = [];

        eachMatrix(matrix, (char, [col, row]) => {
            if (isDigit(char)) {
                currentNumber = currentNumber + char;
                currentNumberCords.push([col, row]);
                currentNumberIsPart = true;
            }

            // If we are at the end of the matrix or the next char is not a digit
            if (isLastColOfMatrix(matrix, [col, row]) || !isDigit(char)) {
                if (currentNumberIsPart) {
                    parts.push({
                        number: Number(currentNumber),
                        cords: [...currentNumberCords],
                    });
                }
                currentNumberIsPart = false;
                currentNumber = "";
                currentNumberCords = [];
            }
        });

        const gearRatios = [];

        eachMatrix(matrix, (char, [col, row]) => {
            if (isGearSymbol(char)) {
                const nearbyParts = parts.filter(({cords: partCords}) => {
                    return partCords.some(([partCol, partRow]) => {
                        let nearGear = false;
                        eachSurrounding(matrix, [partCol, partRow], (char, [c, r]) => {
                            if (c === col && r === row) {
                                nearGear = true;
                            }
                        });
                        return nearGear;
                    });
                });
                if (nearbyParts.length === 2) {
                    gearRatios.push(nearbyParts[0].number * nearbyParts[1].number);
                }
            }
        });

        return arraySum(gearRatios);
    },
};
