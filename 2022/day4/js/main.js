"use strict";

const fs = require("fs");

async function part1(input) {
  let overlap = 0;

  input.forEach((i) => {
    const [pair1, pair2] = i.split(",");
    let [p1n1, p1n2] = pair1.split("-");
    p1n1 = Number(p1n1);
    p1n2 = Number(p1n2);
    let [p2n1, p2n2] = pair2.split("-");
    p2n1 = Number(p2n1);
    p2n2 = Number(p2n2);
    const range1 = Array.from(
      new Array(p1n2 - p1n1 + 1),
      (x, i) => i + p1n1 + 0
    );
    const range2 = Array.from(new Array(p2n2 - p2n1 + 1), (x, i) => i + p2n1);
    const hasOverlap =
      range1.every((i) => range2.indexOf(i) >= 0) ||
      range2.every((i) => range1.indexOf(i) >= 0);
    if (hasOverlap) overlap++;
  });
  return overlap;
}

async function part2(input) {
  let overlap = 0;

  input.forEach((i) => {
    const [pair1, pair2] = i.split(",");
    let [p1n1, p1n2] = pair1.split("-");
    p1n1 = Number(p1n1);
    p1n2 = Number(p1n2);
    let [p2n1, p2n2] = pair2.split("-");
    p2n1 = Number(p2n1);
    p2n2 = Number(p2n2);
    const range1 = Array.from(
      new Array(p1n2 - p1n1 + 1),
      (x, i) => i + p1n1 + 0
    );
    const range2 = Array.from(new Array(p2n2 - p2n1 + 1), (x, i) => i + p2n1);
    const hasOverlap =
      range1.some((i) => range2.indexOf(i) >= 0) ||
      range2.some((i) => range1.indexOf(i) >= 0);
    if (hasOverlap) overlap++;
  });
  return overlap;
}

async function main() {
  const data = fs.readFileSync("../input.txt", { encoding: "utf8" });
  const input = data.split("\n");

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
