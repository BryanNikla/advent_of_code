module.exports = {one, two, solutions: [35, 46]};

const {forX} = require(global.UTILITIES_PATH);

function isInRange(value, min, max) {
    return value >= min && value <= max;
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

function findValueInMapping(value, mapping) {
    const map = mapping.find(({dest, source, length}) => isInRange(value, source, source + length));
    if (map) {
        const d = value - map.source;
        return map.dest + d;
    }
    return value; // no mapping found, return original value
}

function one(input) {
    const {seeds, ...maps} = parseInput(input);
    return Math.min(...seeds.map((seed) => Object.values(maps).reduce((acc, map) => findValueInMapping(acc, map), seed)));
}

function two() {}
// function two(input) {
//     const {seeds, ...maps} = parseInput(input);

//     console.time("RunTime");

//     let lowest = null;
//     let rangeStart = 0;
//     seeds.forEach((val, i) => {
//         if (!isEven(i)) {
//             forX(
//                 val,
//                 (x) => {
//                     const low = Math.min(Object.values(maps).reduce((acc, map) => findValueInMapping(acc, map), x));
//                     if (low < lowest || lowest == null) {
//                         console.log("new lowest:", low);
//                         lowest = low;
//                     }
//                 },
//                 rangeStart
//             );
//         } else {
//             rangeStart = val;
//         }
//     });

//     console.log();
//     console.timeEnd("RunTime");
//     return lowest;
// }
