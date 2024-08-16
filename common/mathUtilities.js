function greatestCommonDivisor(a, b) {
    return b === 0 ? a : greatestCommonDivisor(b, a % b);
}

function leastCommonMultiple(arr) {
    let lcm = arr[0];
    for (let i = 1; i < arr.length; i++) {
        lcm = (lcm * arr[i]) / greatestCommonDivisor(lcm, arr[i]);
    }
    return lcm;
}

module.exports = {
    greatestCommonDivisor,
    leastCommonMultiple,
    gcd: greatestCommonDivisor,
    lcm: leastCommonMultiple,
};
