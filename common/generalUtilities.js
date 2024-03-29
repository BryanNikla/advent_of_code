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
 * @description Sort an array of objects by a given property ASCENDING. To be used in the .sort() method
 * @param {string!} property
 * @returns {function(*, *): number}
 */
const sortASC = (property) => (a, b) => a[property] < b[property] ? -1 : a[property] > b[property] ? 1 : 0;

/**
 * @description Sort an array of objects by a given property DESCENDING. To be used in the .sort() method
 * @param {string!} property
 * @returns {function(*, *): number}
 */
const sortDESC = (property) => (a, b) => a[property] > b[property] ? -1 : a[property] < b[property] ? 1 : 0;

/**
 * Prep to color the console output text
 * @param {'black'|'red'|'green'|'yellow'|'blue'|'magenta'|'cyan'|'white'}} color
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

/**
 * @description  Create an array of a certain length whose values are the index of the array
 * @param {Number} length
 */
function arrOfLength(length = 0) {
    length = parseInt(length);
    return length > 0 ? Array.from({length}, (_, i) => i) : [];
}

/**
 * @description Multiply all numbers in an array
 * @param {Array<Number>} arr
 */
function multiplyArray(arr = []) {
    arr.reduce((acc, val) => acc * val, 1);
}

function arrayFrequencies(arr = []) {
    return arr.reduce((acc, cur, i) => {
        acc[cur] = (acc[cur] || 0) + 1;
        return acc;
    }, {});
}

/**
 * @description Run a function x times
 * @param {Number} x - How many times to run the function
 * @param {function} fn - Single parameter, the iteration number (1-based)
 * @param {Number} startAt - What number to start the iteration at (default 1)
 */
function forX(x = 0, fn, startAt = 1) {
    if (typeof x !== "number") {
        throw new Error("forX: x must be a number");
    }
    if (typeof fn !== "function") {
        throw new Error("forX: fn must be a function");
    }
    for (let i = Number(startAt); i < x + startAt; i++) {
        fn(i);
    }
}

function equalArrays(arr1, arr2) {
    return arr1.every((value, index) => {
        return value === arr2[index];
    });
}

function maxObjectKey(obj) {
    return Object.keys(obj).reduce((prevKey, currKey) => {
        return obj[prevKey] < obj[currKey] ? currKey : prevKey;
    });
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
    arrOfLength,
    forX,
    multiplyArray,
    sortASC,
    sortDESC,
    arrayFrequencies,
    equalArrays,
    maxObjectKey,
};
