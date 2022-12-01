"use strict";

const fs = require("fs");

async function part1(input) {
  let highest = 0;
  let sum = 0;
  const lines = input.split("\n");
  lines.forEach((line) => {
    if (!line || line.trim().length == 0) {
      if (sum > highest) highest = sum;
      sum = 0;
    }
    const calories = Number(line);
    sum += calories;
  });
  return highest;
}

async function part2(input) {
  let sum = 0;
  const sums = [];
  const lines = input.split("\n");
  lines.forEach((line) => {
    if (!line || line.trim().length == 0) {
      sums.push(sum);
      sum = 0;
    }
    const calories = Number(line);
    sum += calories;
  });
  return sums
    .sort((a, b) => b - a)
    .slice(0, 3)
    .reduce((a, b) => a + b, 0);
}

async function main() {
  const data = fs.readFileSync("../input.txt", { encoding: "utf8" });

  const result1 = await part1(data);
  console.log(`Part 1 = ${result1}`);
  const result2 = await part2(data);
  console.log(`Part 2 = ${result2}`);
}

if (require.main === module) {
  main();
}

module.exports = {
  part1,
  part2,
};
