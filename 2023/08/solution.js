module.exports = {one, two, solutions: [6, 6]};

function parseInput(inputString) {
    const [steps, _, ...mapLines] = inputString.split("\n");
    return {
        stepsArray: steps.split(""),
        map: new Map(mapLines.map((s) => s.split(" = ")).map(([k, str]) => [k, str.replaceAll("(", "").replaceAll(")", "").split(", ")])),
    };
}

/**
 * @description Generator function that yields the next step indefinately
 * @param {Array} stepArray
 * @see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Generator
 */
function* steps(stepArray = []) {
    let current = -1;
    while (true) {
        yield stepArray[++current % stepArray.length] === "L" ? 0 : 1;
    }
}

function one(input) {
    const {stepsArray, map} = parseInput(input);
    const stepGenerator = steps(stepsArray);
    let [i, pos] = [-1, "AAA"];
    while (-1 < ++i && pos !== "ZZZ") {
        pos = map.get(pos)[stepGenerator.next().value];
    }
    return i;
}

function two(input) {
    const {leastCommonMultiple} = require(global.UTILITIES_PATH);
    const {stepsArray, map} = parseInput(input);
    return leastCommonMultiple(
        Array.from(map.keys())
            .filter((str) => str.endsWith("A"))
            .map((key) => {
                const stepGenerator = steps(stepsArray);
                let [i, pos] = [-1, key];
                while (-1 < ++i && !pos.endsWith("Z")) {
                    pos = map.get(pos)[stepGenerator.next().value];
                }
                return i;
            })
    );
}
