module.exports = {
    solutions: [142, 281],
    one: (input) => {
        const lines = input.split("\n");
        const digits = lines.map((line) => line.match(/\d/g));
        const numbers = digits.map((arr) => {
            if (Array.isArray(arr) && arr.length) {
                return `${arr[0]}${arr.slice(-1)}`;
            }
            return 0;
        });
        return numbers.reduce((acc, cur) => acc + Number(cur), 0);
    },
    two: (input) => {
        const lines = input.split("\n").filter((line) => line.length);
        // const values = lines.map((line) => line.match(/one|two|three|four|five|six|seven|eight|nine|\d/g));
        const values = lines.map((line) => {
            const pattern = /one|two|three|four|five|six|seven|eight|nine|\d/g;
            let matches = [];
            while ((found = pattern.exec(line))) {
                matches.push(found[0]);
                pattern.lastIndex = found.index + 1;
            }
            return matches;
        });

        const digits = values.map((arr) => {
            return arr.map((val) => {
                if (val === "one") return 1;
                if (val === "two") return 2;
                if (val === "three") return 3;
                if (val === "four") return 4;
                if (val === "five") return 5;
                if (val === "six") return 6;
                if (val === "seven") return 7;
                if (val === "eight") return 8;
                if (val === "nine") return 9;
                return Number(val);
            });
        });

        const numbers = digits.map((arr) => {
            if (Array.isArray(arr) && arr.length) {
                return `${arr[0]}${arr.slice(-1)}`;
            }
            return 0;
        });

        return numbers.reduce((acc, cur) => acc + Number(cur), 0);
    },
};
