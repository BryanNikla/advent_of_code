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
function* stepGenerator(stepArray = []) {
    let current = -1;
    while (true) {
        yield stepArray[++current % stepArray.length] === "L" ? 0 : 1;
    }
}

function one(input) {
    const {stepsArray, map} = parseInput(input);
    let [steps, pos, gen] = [-1, "AAA", stepGenerator(stepsArray)];
    while (-1 < ++steps && pos !== "ZZZ") {
        pos = map.get(pos)[gen.next().value];
    }
    return steps;
}

function two(input) {
    const {leastCommonMultiple} = require(global.UTILITIES_PATH);
    const {stepsArray, map} = parseInput(input);
    return leastCommonMultiple(
        Array.from(map.keys())
            .filter((str) => str.endsWith("A"))
            .map((pos) => {
                let [steps, gen] = [-1, stepGenerator(stepsArray)];
                while (-1 < ++steps && !pos.endsWith("Z")) {
                    pos = map.get(pos)[gen.next().value];
                }
                return steps;
            })
    );
}
