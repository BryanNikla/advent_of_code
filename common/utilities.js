/**
 * Pads a number with zeros to a set number of places.
 * @param {*} num
 * @param {Number} places
 */
function zeroPad(num, places) {
    return String(num).padStart(places, "0");
}

/**
 * Sum an array (will convert strings to numbers)
 * @param {Array} arr
 */
function arraySum(arr = []) {
    return arr.reduce((partialSum, a) => partialSum + Number(a), 0);
}

/**
 * Prep to color the console output text
 * @param {string} color
 * @param {string} text
 */
function colorText(color, text = "") {
    const reset = "\x1b[0m";
    const colors = {
        black: "\x1b[30m",
        red: "\x1b[31m",
        green: "\x1b[32m",
        yellow: "\x1b[33m",
        blue: "\x1b[34m",
        magenta: "\x1b[35m",
        cyan: "\x1b[36m",
        white: "\x1b[37m",
    };
    return String(colors[color] ? colors[color] + text + reset : text);
}

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

/**
 * Checks if a value (usually an object) has a property.
 * @param {*} val - the value to check for a property
 * @param {string} prop - the property name to check for
 * @returns {boolean} whether or the property is on the value
 */
function hasProp(val, prop) {
    if (val == null) {
        return false;
    }

    return hasOwnProperty.call(val, prop);
}

// Use as an array filter
// TODO: Make a new util file of just array filters
function distinctValues(value, index, array) {
    return array.indexOf(value) === index;
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////
module.exports = {
    zeroPad,
    arraySum,
    colorText,
    callAtCords,
    eachMatrix,
    eachSurrounding,
    hasProp,
    distinctValues,
};
///////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////
