const {arraySum} = require(global.UTILITIES_PATH);

module.exports.solution = function() {
    
    const fs = require('fs');
    const path = require('path');
    const input = fs.readFileSync(path.resolve(__dirname, 'input.txt'), 'utf8');
    return [one(input), two(input)];
}

module.exports.test = function() {
    const fs = require('fs');
    const path = require('path');
    const input = fs.readFileSync(path.resolve(__dirname, 'input_text.txt'), 'utf8');
    return [
        one(input) === 24000, 
        two(input) === 45000,
    ];
}


function one(input) {
    let elf_items = input.split("\n\n").map(r => r.split("\n"));
    const totals = new Int32Array(elf_items.map(arraySum));
    return Math.max(...totals);
}

function two(input) {
    let elf_items = input.split("\n\n").map(r => r.split("\n"));
    const totals = new Int32Array(elf_items.map(arraySum));
    const top_three = totals.sort().slice(-3);
    return arraySum(top_three);
}