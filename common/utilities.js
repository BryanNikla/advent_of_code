module.exports.zeroPad = (num, places) => String(num).padStart(places, '0')

module.exports.arraySum = function(arr = []) {
    return arr.reduce((partialSum, a) => partialSum + Number(a), 0);
}