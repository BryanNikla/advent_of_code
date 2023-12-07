module.exports = {one, two, solutions: [6440, 5905]};

const {distinctValues, sortASC, sortDESC} = require(global.UTILITIES_PATH);

function arrayFrequencies(arr = []) {
    return arr.reduce((acc, cur, i) => {
        acc[cur] = (acc[cur] || 0) + 1;
        return acc;
    }, {});
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

function getHandValue(hand, jokers = false) {
    const chars = hand.split("");
    let val = "";

    let freqencies = arrayFrequencies(chars);
    if (jokers) {
        if (!chars.every((c) => c == "J")) {
            const charsWithoutJokers = chars.filter((c) => c !== "J");
            const totalJokers = chars.length - charsWithoutJokers.length;
            freqencies = arrayFrequencies(charsWithoutJokers);
            const maxKey = maxObjectKey(freqencies);
            if (maxKey) {
                freqencies[maxKey] += totalJokers;
            }
        }
    }

    const sortedFrequencies = Object.values(freqencies).sort();

    switch (true) {
        // high hand
        case equalArrays(sortedFrequencies, [1, 1, 1, 1, 1]):
            val = "1";
            break;
        // pair
        case equalArrays(sortedFrequencies, [1, 1, 1, 2]):
            val = "2";
            break;
        // two pair
        case equalArrays(sortedFrequencies, [1, 2, 2]):
            val = "3";
            break;
        // three of a kind
        case equalArrays(sortedFrequencies, [1, 1, 3]):
            val = "4";
            break;
        // full house
        case equalArrays(sortedFrequencies, [2, 3]):
            val = "5";
            break;
        // four of a kind
        case equalArrays(sortedFrequencies, [1, 4]):
            val = "6";
            break;
        // five of a kind
        case equalArrays(sortedFrequencies, [5]):
            val = "7";
            break;
    }

    chars.forEach((c) => {
        // Replace 'suit' letters with characters that WILL sort correctly
        switch (c) {
            case "A":
                val += "E";
                break;
            case "K":
                val += "D";
                break;
            case "Q":
                val += "C";
                break;
            case "J":
                val += jokers ? "0" : "B";
                break;
            case "T":
                val += "A";
                break;
            default:
                val += c;
                break;
        }
    });
    return val;
}

function one(input) {
    const input_lines = input.split("\n").map((line) => line.split(" "));
    const result = input_lines.map(([hand, bid]) => ({hand, bid: Number(bid), value: getHandValue(hand, false)}));
    return result.sort(sortASC("value")).reduce((acc, {hand, bid, value}, i) => {
        return acc + bid * (i + 1);
    }, 0);
}

function two(input) {
    const input_lines = input.split("\n").map((line) => line.split(" "));
    const result = input_lines.map(([hand, bid]) => ({hand, bid: Number(bid), value: getHandValue(hand, true)}));
    return result.sort(sortASC("value")).reduce((acc, {hand, bid, value}, i) => {
        return acc + bid * (i + 1);
    }, 0);
}
