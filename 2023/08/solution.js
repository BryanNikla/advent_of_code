module.exports = {one, two, solutions: [6, 0]};

function parseInput(inputString) {
    const [instructions, _, ...mapLines] = inputString.split("\n");
    const map = new Map();
    mapLines
        .map((s) => s.split(" = "))
        .forEach(([k, str]) => {
            [left, right] = str.replaceAll("(", "").replaceAll(")", "").split(", ");
            map.set(k, [left, right]);
        });
    return {instructions, map};
}

function useNextInstruction(instructions) {
    instructions = instructions.split("");
    let current = -1;
    return {
        getNextStep() {
            current++;
            if (current > instructions.length - 1) {
                current = 0;
            }
            return instructions[current];
        },
    };
}

function one(input) {
    const {instructions, map} = parseInput(input);
    const {getNextStep} = useNextInstruction(instructions);
    let steps = 0;
    let position = "AAA";
    while (position !== "ZZZ") {
        steps++;
        const [left, right] = map.get(position);
        position = getNextStep() === "L" ? left : right;
    }
    return steps;
}

function two(input) {
    return 0;
}
