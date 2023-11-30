function execute(year, day) {
    const path = require('path');
    const zeroPad = (num, places) => String(num).padStart(places, '0')

    const output = {};
    if (day === -1) {
        for (let i = 1; i < 26; i++) {
            output[parseFloat(year + '.' + zeroPad(i, 2)).toFixed(2)] = _solve(year, i);
        }
    } else {
        output[parseFloat(year + '.' + zeroPad(day, 2)).toFixed(2)] = _solve(year, day);
    }

    console.table(output);

    function _solve(year, day) {
        const filePath = path.resolve(__dirname, String(year), zeroPad(day, 2), 'solution.js');
        try {
            const {solution, test} = require(filePath);
            const [answer_1, answer_2] = solution();
            const rsp = {'Solution 1': answer_1 ?? '', 'Solution 2': answer_2 ?? ''};
            if (typeof test === 'function') {
                const [test_1, test_2] = test();
                Object.assign(rsp, {'Test 1': test_1, 'Test 2': test_2});
            }
            return rsp;
        } catch (err) {
            switch (true) {
                case String(err).includes(filePath):
                    return {};
                default:
                    console.error(err);
                    process.exit();
            }
        }
    }
}

console.clear();

const yearIndex = process.argv.indexOf('-year');
const dayIndex = process.argv.indexOf('-day');

if (yearIndex === -1 || dayIndex === -1) {
    const readline = require('readline').createInterface({input: process.stdin, output: process.stdout});
    readline.question('What year? ', (year) => {
        readline.question("What day? (answer 'all' for all days in this year) ", (day) => {
            readline.close();
            year = parseInt(year);
            day = day === 'all' ? -1 : parseInt(day);

            if (!year || isNaN(year) || !day || isNaN(day)) {
                console.error('year & day arguements must be valid');
                process.exit();
            }
            execute(year, day);
        });
    });
} else {
    const year = parseInt(process.argv[yearIndex + 1]);
    const dayArg = process.argv[dayIndex + 1];
    const day = dayArg === 'all' ? -1 : parseInt(process.argv[dayIndex + 1]);
    if (!year || isNaN(year) || !day || isNaN(day)) {
        console.error('year & day arguements must be valid');
        process.exit();
    }
    execute(year, day);
}

