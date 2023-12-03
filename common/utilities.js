// Exoport all functions from the utilities folder at once to avoid having to import each one individually in solution.js
module.exports = {
    ...require("./generalUtilities.js"),
    ...require("./arrayFilters.js"),
    ...require("./matrixUtilities.js"),
};
