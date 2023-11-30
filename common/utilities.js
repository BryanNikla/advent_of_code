module.exports.arraySum = function(arr = []) {
    return arr.reduce((partialSum, a) => partialSum + Number(a), 0);
}