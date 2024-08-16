module.exports = {
    solutions: [8, 2286],
    one: (input) => {
        const red = 12;
        const green = 13;
        const blue = 14;

        const test_against = [red, green, blue];

        const lines = input.split("\n");
        const games = lines.map((line) => {
            const [g, x] = line.split(": ");
            const xx = x.split(";");

            const xxx = xx.map((str) => {
                const pulls = str.split(",");
                const values = [0, 0, 0];
                pulls.forEach((pull) => {
                    switch (true) {
                        case pull.includes("red"):
                            values[0] = parseInt(pull.replace(/[^0-9]/g, ""));
                            break;
                        case pull.includes("green"):
                            values[1] = parseInt(pull.replace(/[^0-9]/g, ""));
                            break;
                        case pull.includes("blue"):
                            values[2] = parseInt(pull.replace(/[^0-9]/g, ""));
                            break;
                    }
                });
                return values;
            });

            return [parseInt(g.replace(/[^0-9]/g, "")), xxx];
        });

        let sum = 0;

        games.forEach(([gameNumber, pulls]) => {
            if (
                pulls.every((pull) => {
                    if (pull[0] > test_against[0]) {
                        return false;
                    }
                    if (pull[1] > test_against[1]) {
                        return false;
                    }
                    if (pull[2] > test_against[2]) {
                        return false;
                    }
                    return true;
                })
            ) {
                sum += gameNumber;
            }
        });

        return sum;
    },
    two: (input) => {
        const lines = input.split("\n");
        const games = lines.map((line) => {
            const [g, x] = line.split(": ");
            const xx = x.split(";");

            const xxx = xx.map((str) => {
                const pulls = str.split(",");
                const values = [0, 0, 0];
                pulls.forEach((pull) => {
                    switch (true) {
                        case pull.includes("red"):
                            values[0] = parseInt(pull.replace(/[^0-9]/g, ""));
                            break;
                        case pull.includes("green"):
                            values[1] = parseInt(pull.replace(/[^0-9]/g, ""));
                            break;
                        case pull.includes("blue"):
                            values[2] = parseInt(pull.replace(/[^0-9]/g, ""));
                            break;
                    }
                });
                return values;
            });

            return [parseInt(g.replace(/[^0-9]/g, "")), xxx];
        });

        games2 = games.map(([game, pulls]) => {
            const greatest = [0, 0, 0];
            pulls.forEach((pull) => {
                if (pull[0] > greatest[0]) {
                    greatest[0] = pull[0];
                }
                if (pull[1] > greatest[1]) {
                    greatest[1] = pull[1];
                }
                if (pull[2] > greatest[2]) {
                    greatest[2] = pull[2];
                }
            });

            return greatest[0] * greatest[1] * greatest[2];
        });

        const sum = games2.reduce((acc, cur) => acc + cur, 0);
        return sum;
    },
};
