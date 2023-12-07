module.exports = {one, two, solutions: [6440, 5905]};

const {sortASC, arrayFrequencies, equalArrays, maxObjectKey} = require(global.UTILITIES_PATH);

function getHandValue(hand, jokers = false) {
    const chars = hand.split("");
    let val = "";
    let freqencies = arrayFrequencies(chars);
    if (jokers) {
        if (!chars.every((c) => c == "J")) {
            const charsWithoutJokers = chars.filter((c) => c !== "J");
            freqencies = arrayFrequencies(charsWithoutJokers);
            const maxKey = maxObjectKey(freqencies);
            if (maxKey) {
                freqencies[maxKey] += chars.length - charsWithoutJokers.length;
            }
        }
    }

    const sortedFrequencies = Object.values(freqencies).sort();
    switch (true) {
        case equalArrays(sortedFrequencies, [1, 1, 1, 1, 1]): // high hand
            val = "1";
            break;
        case equalArrays(sortedFrequencies, [1, 1, 1, 2]): // pair
            val = "2";
            break;
        case equalArrays(sortedFrequencies, [1, 2, 2]): // two pair
            val = "3";
            break;
        case equalArrays(sortedFrequencies, [1, 1, 3]): // three of a kind
            val = "4";
            break;
        case equalArrays(sortedFrequencies, [2, 3]): // full house
            val = "5";
            break;
        case equalArrays(sortedFrequencies, [1, 4]): // four of a kind
            val = "6";
            break;
        case equalArrays(sortedFrequencies, [5]): // five of a kind
            val = "7";
            break;
    }

    // Substitute 'suit' letters with characters that WILL sort correctly
    chars.forEach((c) => {
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
    return result.sort(sortASC("value")).reduce((acc, {bid}, i) => {
        return acc + bid * (i + 1);
    }, 0);
}

function two(input) {
    const input_lines = input.split("\n").map((line) => line.split(" "));
    const result = input_lines.map(([hand, bid]) => ({hand, bid: Number(bid), value: getHandValue(hand, true)}));
    return result.sort(sortASC("value")).reduce((acc, {bid}, i) => {
        return acc + bid * (i + 1);
    }, 0);
}
