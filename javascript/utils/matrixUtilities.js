const {hasProp} = require("./generalUtilities.js");

/**
 * Calls a function at a set of coordinates if they exist.
 * @param {Array.<Array.<string>>} matrix
 * @param {Array<[Number, Number]>} cords
 * @param {function name(cellValue, [col, row], matrix)} fn
 */
function callAtCords(matrix, cords, fn) {
    const [col, row] = cords;
    if (hasProp(matrix, row) && hasProp(matrix[row], col)) {
        fn(matrix[row][col], cords, matrix);
    }
}

/**
 * Iterates through a matrix and calls a function for each cell.
 * @param {Array.<Array.<string>>} matrix
 * @param {function name(cellValue, [col, name], matrix)} fn
 */
function eachMatrix(matrix, fn) {
    for (let row = 0; row < matrix.length; row++) {
        for (let col = 0; col < matrix[row].length; col++) {
            callAtCords(matrix, [col, row], fn);
        }
    }
}

function isLastColOfMatrix(matrix, [col, row]) {
    return col === matrix[row].length - 1;
}

/**
 * callAtCords for every cordinate surrounding a set of cords.
 * @param {Array.<Array.<string>>} matrix
 * @param {Array<[Number, Number]>} cords
 * @param {function name(cellValue, [col, row], matrix)} fn
 */
function eachSurrounding(matrix, [col, row], fn) {
    callAtCords(matrix, [col, row - 1], fn);
    callAtCords(matrix, [col + 1, row - 1], fn);
    callAtCords(matrix, [col + 1, row], fn);
    callAtCords(matrix, [col + 1, row + 1], fn);
    callAtCords(matrix, [col, row + 1], fn);
    callAtCords(matrix, [col - 1, row + 1], fn);
    callAtCords(matrix, [col - 1, row], fn);
    callAtCords(matrix, [col - 1, row - 1], fn);
}

///////////////////////////////
///////////////////////////////
module.exports = {
    callAtCords,
    eachMatrix,
    isLastColOfMatrix,
    eachSurrounding,
};
