module.exports.solution = function() {

    const fs = require('fs');
    const path = require('path');
    const input = fs.readFileSync(path.resolve(__dirname, 'input.txt'), 'utf8');

    return ['part 1 answer', 'part 2 answer'];
}