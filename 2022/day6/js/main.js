"use strict";

const fs = require("fs");

async function part1(input) {
  let result = 0;

  const characters = input.split("");
  for (let i = 0; i < characters.length; i++) {
    const sequence = characters.slice(i, i + 4);
    const notIt = sequence.some((c) => {
      if (sequence.filter((j) => c == j).length > 1) return true;
    });
    if (!notIt) {
      result = i + 4;
      break;
    }
  }

  return result;
}

async function part2(input) {
  let result = 0;

  const characters = input.split("");
  for (let i = 0; i < characters.length; i++) {
    const sequence = characters.slice(i, i + 14);
    const notIt = sequence.some((c) => {
      if (sequence.filter((j) => c == j).length > 1) return true;
    });
    if (!notIt) {
      result = i + 14;
      break;
    }
  }

  return result;
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
