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
 * Checks if a value (usually an object) has a property.
 * @param {*} val
 * @param {string} prop
 * @returns {boolean}
 */
function hasProp(val, prop) {
    if (val == null) {
        return false;
    }

    return hasOwnProperty.call(val, prop);
}

function isBetween(number, min, max) {
    return number >= min && number <= max;
}

function reverseString(str) {
    return str.split("").reverse().join("");
}

///////////////////////////////
///////////////////////////////
module.exports = {
    zeroPad,
    arraySum,
    colorText,
    hasProp,
    isBetween,
    reverseString,
};
