const {arraySum} = require(global.UTILITIES_PATH);

module.exports = {
    solutions: [24000, 45000],
    one: (input) => {
        let elf_items = input.split("\n\n").map(r => r.split("\n"));
        const totals = new Int32Array(elf_items.map(arraySum));
        return Math.max(...totals);
    },
    two: (input) => {
        let elf_items = input.split("\n\n").map(r => r.split("\n"));
        const totals = new Int32Array(elf_items.map(arraySum));
        const top_three = totals.sort().slice(-3);
        return arraySum(top_three);
    },
};