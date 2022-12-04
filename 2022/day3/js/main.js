"use strict";

const fs = require("fs");

const alphabet = "abcdefghijklmnopqrstuvwxyz";
const priorities = [...alphabet.split(""), ...alphabet.toUpperCase().split("")];

async function part1(input) {
  let priority = 0;

  input.forEach((rucksack) => {
    const c1 = rucksack.substring(0, rucksack.length / 2).split("");
    const c2 = rucksack
      .substring(rucksack.length / 2, rucksack.length)
      .split("");

    let foundItem = null;
    c1.some((i) => {
      const found = c2.some((j) => i == j);
      if (found) foundItem = i;
      return found;
    });
    const itemPriority = priorities.indexOf(foundItem) + 1;
    priority += itemPriority;
  });

  return priority;
}

async function part2(input) {
  let priority = 0;

  for (let i = 0; i < input.length; i += 3) {
    let foundItem = null;
    const c1 = input[i].split("");
    const c2 = input[i + 1];
    const c3 = input[i + 2];
    c1.some((i) => {
      let found = false;
      if (c2.indexOf(i) >= 0 && c3.indexOf(i) >= 0) {
        foundItem = i;
        found = true;
      }
      return found;
    });
    const itemPriority = priorities.indexOf(foundItem) + 1;
    priority += itemPriority;
  }
  return priority;
}

async function main() {
  const data = fs.readFileSync("../input.txt", { encoding: "utf8" });
  const rucksacks = data.split("\n");

  const result1 = await part1(rucksacks);
  console.log(`Part 1 = ${result1}`);
  const result2 = await part2(rucksacks);
  console.log(`Part 2 = ${result2}`);
}

if (require.main === module) {
  main();
}

module.exports = {
  part1,
  part2,
};
