"use strict";

const fs = require("fs");

async function part1(input) {
  let result = 0;
  return result;
}

async function part2(input) {
  let result = 0;
  return result;
}

async function main() {
  const input = fs.readFileSync("../input.txt", { encoding: "utf8" });
  // @TODO - transform

  const result1 = await part1(input);
  console.log(`Part 1 = ${result1}`);
  const result2 = await part2(input);
  console.log(`Part 2 = ${result2}`);
}

if (require.main === module) {
  main();
}

module.exports = {
  part1,
  part2,
};
