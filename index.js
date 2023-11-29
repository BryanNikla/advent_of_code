function execute(year, day) {
    try {
        const {solution} = require(`./${year}/${day < 10 ? `0${day}` : day}/solution`);
        const [answer_1, answer_2] = solution();
        console.table({Puzzle: parseFloat(year + '.' + day), 'Answer 1': answer_1, 'Answer 2': answer_2});
    } catch (err) {
        switch (true) {
            case String(err).includes('Cannot find module'):
                console.error('Cannot find this problem');
                break;
            default:
                console.error('Unexepected error execution solution', err);
        }
    }
}

const yearIndex = process.argv.indexOf('--year');
const dayIndex = process.argv.indexOf('--day');

if (yearIndex === -1 || dayIndex === -1) {
    const readline = require('readline').createInterface({input: process.stdin, output: process.stdout});
    readline.question('What year? ', (year) => {
        readline.question("What day? ", (day) => {
            readline.close();
            year = parseInt(year);
            day = parseInt(day);
            if (!year || isNaN(year) || !day || isNaN(day)) {
                console.error('year & day arguements must be valid');
                return;
            }
            execute(year, day);
        });
    });
} else {
    const year = parseInt(process.argv[yearIndex + 1]);
    const day = parseInt(process.argv[dayIndex + 1]);
    if (!year || isNaN(year) || !day || isNaN(day)) {
        console.error('year & day arguements must be valid');
        return;
    }
    execute(year, day);
}

