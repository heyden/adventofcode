"use strict";

const { ifError } = require("assert");
const fs = require("fs");
const { isMainThread } = require("worker_threads");

async function part1(input) {
  let register = 1;
  let c = 1;
  let v = undefined;
  let strengths = [];
  const instructions = [...input];

  while (true) {
    if (c === 20 || (c - 20) % 40 == 0) strengths.push(register * c);
    if (v) {
      register += v;
      v = undefined;
    } else {
      const inst = instructions.pop();
      if (inst === "noop") v = undefined;
      else if (inst.includes("addx")) v = Number(inst.split(" ")[1]);
    }
    if (v === undefined && instructions.length === 0) break;
    c++;
  }
  return strengths.reduce((a, c) => a + c, 0);
}

async function part2(input) {
  let register = 1;
  let c = 1;
  let v = undefined;
  let screen = [[], [], [], [], [], []];
  let r = 0;
  const instructions = [...input];

  while (true) {
    let i = screen[r].length;
    let pixel = ".";
    if (screen[r].length === 40) r += 1;
    if (i >= register - 1 && i <= register + 1) pixel = "#";
    screen[r].push(pixel);

    if (v) {
      register += v;
      v = undefined;
    } else {
      const inst = instructions.pop();
      if (inst === "noop") v = undefined;
      else if (inst.includes("addx")) v = Number(inst.split(" ")[1]);
    }
    if (v === undefined && instructions.length === 0) break;
    c++;
  }
  screen.forEach((r) => console.log(r.join("")));
  return 0;
}

async function main() {
  const data = fs.readFileSync("../input.txt", { encoding: "utf8" });
  let input = data.split("\n");
  input.reverse();
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
