module.exports = {one, two, solutions: [35, 46]};

const {forX, arrOfLength} = require(global.UTILITIES_PATH);

function isInRange(n, min, max) {
    return n >= min && n <= max;
}

function isEven(value) {
    return value % 2 === 0;
}

function parseInput(input) {
    let parts = input.split("\n\n");
    function parseMapping(line) {
        const v = line.split(" ").map(Number);
        return {dest: v[0], source: v[1], length: v[2]};
    }
    function prepSection(section) {
        return section.split("\n").filter((_, i) => i > 0);
    }
    return {
        seeds: parts[0].split(": ")[1].split(" ").map(Number),
        maps: arrOfLength(7).map((_, i) => prepSection(parts[i + 1]).map(parseMapping)),
    };
}

function findValueInMaps(value, maps) {
    const map = maps.find(({source, length}) => {
        return isInRange(value, source, source + length - 1);
    });
    return map ? map.dest + (value - map.source) : value;
}

function one(input) {
    const {seeds, maps} = parseInput(input);
    return Math.min(...seeds.map((seed) => maps.reduce((acc, map) => findValueInMaps(acc, map), seed)));
}

function two(input, {isTest}) {
    if (!isTest) {
        // My solution is incredibly inefficent, but succesful.
        // This is here to speed up a full run and skip this portion.
        return "skip";
    }

    const {seeds: seedData, maps} = parseInput(input);

    const ranges = seedData.map((n, i, arr) => [n, arr[i + 1]]).filter((_, i) => i % 2 === 0);

    let lowest = null;
    ranges.forEach(([start, length]) => {
        forX(
            length,
            (seed) => {
                const location = maps.reduce((acc, map) => findValueInMaps(acc, map), seed);
                if (location < lowest || lowest == null) {
                    lowest = location;
                }
            },
            start
        );
    });

    return lowest;
}
