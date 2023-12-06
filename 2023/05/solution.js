module.exports = {one, two, solutions: [35, 46]};

const {forX} = require(global.UTILITIES_PATH);

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
        seed_soil: prepSection(parts[1]).map(parseMapping),
        soil_fertilizer: prepSection(parts[2]).map(parseMapping),
        fertilizer_water: prepSection(parts[3]).map(parseMapping),
        water_light: prepSection(parts[4]).map(parseMapping),
        light_temp: prepSection(parts[5]).map(parseMapping),
        temp_humidity: prepSection(parts[6]).map(parseMapping),
        humidity_location: prepSection(parts[7]).map(parseMapping),
    };
}

function findValueInMaps(value, maps) {
    const map = maps.find(({source, length}) => {
        return isInRange(value, source, source + length - 1);
    });
    return map ? map.dest + (value - map.source) : value;
}

function one(input) {
    const {seeds, ...maps} = parseInput(input);
    return Math.min(...seeds.map((seed) => Object.values(maps).reduce((acc, map) => findValueInMaps(acc, map), seed)));
}

function two(input) {
    const {seeds: seedData, ...maps} = parseInput(input);

    const ranges = [];
    while (seedData.length > 0) {
        ranges.push([seedData.shift(), seedData.shift()]);
    }

    let lowest = null;
    ranges.forEach(([start, length]) => {
        forX(
            length,
            (seed) => {
                const location = Object.values(maps).reduce((acc, map) => findValueInMaps(acc, map), seed);
                if (location < lowest || lowest == null) {
                    console.log("new lowest location:", location);
                    lowest = location;
                }
            },
            start
        );
    });

    return lowest;
}
