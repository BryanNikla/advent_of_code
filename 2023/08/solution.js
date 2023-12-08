module.exports = {one, two, solutions: [6, 6]};

function parseInput(inputString) {
    const [steps, _, ...mapLines] = inputString.split("\n");
    const map = new Map(mapLines.map((s) => s.split(" = ")).map(([k, str]) => [k, str.replaceAll("(", "").replaceAll(")", "").split(", ")]));
    return {steps, map};
}

function useSteps(instructions) {
    let [current, pos] = [-1, instructions.split("")];
    return {next: () => (pos[++current % instructions.length] === "L" ? 0 : 1)};
}

function one(input) {
    const {steps, map} = parseInput(input);
    const {next} = useSteps(steps);
    let [i, pos] = [-1, "AAA"];
    while (-1 < ++i && pos !== "ZZZ") {
        pos = map.get(pos)[next()];
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
                const {next} = useSteps(steps);
                let [i, pos] = [-1, key];
                while (-1 < ++i && !pos.endsWith("Z")) {
                    pos = map.get(pos)[next()];
                }
                return i;
            })
    );
}
