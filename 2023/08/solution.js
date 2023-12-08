////////////////////////////////////////////////////////////////
module.exports = {one, two, solutions: [6, 6]};
////////////////////////////////////////////////////////////////

function parseInput(inputString) {
    const [steps, _, ...mapLines] = inputString.split("\n");
    const map = new Map(mapLines.map((s) => s.split(" = ")).map(([k, str]) => [k, str.replaceAll("(", "").replaceAll(")", "").split(", ")]));
    return {steps, map};
}

function useSteps(instructions) {
    let [current, pos] = [-1, instructions.split("")];
    return {
        getNextStep() {
            if (++current > instructions.length - 1) {
                current = 0;
            }
            return pos[current] === "L" ? 0 : 1;
        },
    };
}

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////

function one(input) {
    const {steps, map} = parseInput(input);
    const {getNextStep} = useSteps(steps);
    let [i, pos] = [-1, "AAA"];
    while (-1 < ++i && pos !== "ZZZ") {
        pos = map.get(pos)[getNextStep()];
    }
    return i;
}

function two(input) {
    const {leastCommonMultiple} = require(global.UTILITIES_PATH);
    const {steps, map} = parseInput(input);
    return leastCommonMultiple(
        Array.from(map.keys())
            .filter((str) => str.endsWith("A"))
            .map((key) => {
                const {getNextStep} = useSteps(steps);
                let [i, pos] = [-1, key];
                while (-1 < ++i && !pos.endsWith("Z")) {
                    pos = map.get(pos)[getNextStep()];
                }
                return i;
            })
    );
}
