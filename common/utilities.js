module.exports.zeroPad = (num, places) => String(num).padStart(places, "0");

module.exports.arraySum = function (arr = []) {
    return arr.reduce((partialSum, a) => partialSum + Number(a), 0);
};

module.exports.colorText = function (color, text = "") {
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
    return colors[color] ? colors[color] + text + reset : text;
};
