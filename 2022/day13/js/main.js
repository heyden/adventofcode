"use strict";

const fs = require("fs");

async function part1(input) {
  const pairs = [[]];
  input.forEach((i) => {
    if (i.trim() === "") pairs.push([]);
    else pairs[pairs.length - 1].push(JSON.parse(i));
  });

  let sum = 0;
  pairs.forEach((p, i) => {
    const p1 = p[0];
    const p2 = p[1];
    if (compareLists(p1, p2) <= 0) sum += i + 1;
  });
  return sum;
}

function compareLists(l1, l2) {
  const max = Math.min(l1.length, l2.length);
  for (let i = 0; i < max; i++) {
    const d = compareItems(l1[i], l2[i]);
    if (d !== 0) return d;
  }
  return Math.sign(l1.length - l2.length);
}

function compareItems(i1, i2) {
  if (typeof i1 === "number" && typeof i2 === "number") {
    return Math.sign(i1 - i2);
  } else if (Array.isArray(i1) && typeof i2 === "number") {
    return compareLists(i1, [i2]);
  } else if (typeof i1 === "number" && Array.isArray(i2)) {
    return compareLists([i1], i2);
  } else if (Array.isArray(i1) && Array.isArray(i2))
    return compareLists(i1, i2);
}

async function part2(input) {
  const d1 = [[2]];
  const d2 = [[6]];
  const pairs = [];
  input.forEach((i) => {
    if (i.trim() !== "") pairs.push(JSON.parse(i));
  });
  pairs.push(d1);
  pairs.push(d2);

  pairs.sort((a, b) => compareLists(a, b));

  const jd1 = JSON.stringify(d1);
  const jd2 = JSON.stringify(d2);
  const decoderKeys = [];
  pairs.forEach((p, i) => {
    const jp = JSON.stringify(p);
    if (jp === jd1 || jp === jd2) decoderKeys.push(i + 1);
  });
  return decoderKeys.reduce((a, c) => a * c, 1);
}

async function main() {
  const test_data = fs.readFileSync("../input.txt", { encoding: "utf8" });
  const input = test_data.split("\n");

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
