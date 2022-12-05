"use strict";

const fs = require("fs");

async function part1(input) {
  let result = "";
  const stacks = parseStacks(input);
  const moves = input.split("\n").forEach((move) => {
    if (move.indexOf("move") !== 0) return;
    const moveAmount = Number(
      move.substring(move.indexOf("move") + 5, move.indexOf("from") + -1)
    );
    const moveFrom = Number(
      move.substring(move.indexOf("from") + 5, move.indexOf("to") + -1)
    );
    const moveTo = Number(move.substring(move.indexOf("to") + 2));
    const fromStack = stacks[moveFrom];
    const toStack = stacks[moveTo];
    for (let i = moveAmount; i > 0; i--) {
      const value = fromStack.pop();
      toStack.push(value);
    }
  });
  const tops = [];
  Object.keys(stacks).forEach((stack) => {
    tops.push(stacks[stack].pop());
  });
  return tops.join("");
}

async function part2(input) {
  let result = "";
  const stacks = parseStacks(input);
  const moves = input.split("\n").forEach((move) => {
    if (move.indexOf("move") !== 0) return;
    const moveAmount = Number(
      move.substring(move.indexOf("move") + 5, move.indexOf("from") + -1)
    );
    const moveFrom = Number(
      move.substring(move.indexOf("from") + 5, move.indexOf("to") + -1)
    );
    const moveTo = Number(move.substring(move.indexOf("to") + 2));
    const fromStack = stacks[moveFrom];
    const toStack = stacks[moveTo];
    const moving = fromStack.splice(fromStack.length - moveAmount, moveAmount);
  });
  const tops = [];
  Object.keys(stacks).forEach((stack) => {
    tops.push(stacks[stack].pop());
  });
  return tops.join("");
}

function parseStacks(input) {
  const stacks = {};
  let height = 0;
  let done = false;
  const data = input.split("\n");
  while (!done) {
    const line = data[height];
    if (line.replaceAll(" ", "").match(/[0-9]+/g)) {
      done = true;
      break;
    }
    height++;
  }

  const stackLine = data[height];
  let stack = 0;
  for (let j = 0; j <= stackLine.length; j++) {
    const character = stackLine.substring(j, j + 1);
    if (!character.trim()) {
      continue;
    }
    for (let i = height - 1; i >= 0; i--) {
      const item = data[i].substring(j, j + 1).trim();
      if (item.trim()) {
        if (!stacks[character]) stacks[character] = [];
        stacks[character].push(item);
      }
    }
  }
  return stacks;
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
