// Array filter function -- returns array of distinct values
function distinctValues(value, index, array) {
    return array.indexOf(value) === index;
}

function notNull(val) {
    return val != null;
}

///////////////////////////////
///////////////////////////////
module.exports = {
    distinctValues,
    notNull,
};
