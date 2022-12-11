"use strict";

const fs = require("fs");

async function part1(input) {
  const monkeys = buildMonkeys(input);

  for (let r = 1; r <= 20; r++) {
    Object.keys(monkeys).forEach((m) => {
      const monkey = monkeys[m];
      const items = [...monkey.items];
      items.forEach((item, i) => {
        monkey.inspected = monkey.inspected + 1;
        let operation = monkey.operation.replaceAll("old", item);
        let w = doMath(operation);
        w = Math.floor(w / 3);
        if (w % monkey.test.o === 0) monkeys[monkey.test.t].items.push(w);
        else monkeys[monkey.test.f].items.push(w);
      });
      monkey.items = [];
    });
  }

  const mw = Object.keys(monkeys)
    .map((m) => monkeys[m].inspected)
    .sort((a, b) => b - a);

  return mw[0] * mw[1];
}

async function part2(input) {
  const monkeys = buildMonkeys(input);
  const multiple = Object.keys(monkeys).reduce(
    (a, c) => a * monkeys[c].test.o,
    1
  );

  for (let r = 1; r <= 10000; r++) {
    Object.keys(monkeys).forEach((m) => {
      const monkey = monkeys[m];
      const items = [...monkey.items];
      items.forEach((item) => {
        monkey.inspected = monkey.inspected + 1;
        let operation = monkey.operation.replaceAll("old", item);
        let w = doMath(operation);
        w = w % multiple;
        if (w % monkey.test.o === 0) monkeys[monkey.test.t].items.push(w);
        else monkeys[monkey.test.f].items.push(w);
      });
      monkey.items = [];
    });
  }

  const mw = Object.keys(monkeys)
    .map((m) => monkeys[m].inspected)
    .sort((a, b) => b - a);

  return mw[0] * mw[1];
}

function buildMonkeys(input) {
  const i = input.split("\n");
  const monkeys = {};

  let cm = 0;
  i.forEach((line) => {
    if (line.includes("Monkey")) {
      line = line.replaceAll(":", "");
      const n = line.split(" ")[1];
      cm = Number(n);
      monkeys[cm] = {
        inspected: 0,
        items: [],
        operation: undefined,
        test: {},
      };
    } else if (line.includes("items"))
      monkeys[cm].items = line
        .split(": ")[1]
        .split(", ")
        .map((n) => Number(n));
    else if (line.includes("Operation"))
      monkeys[cm].operation = line.split(" = ")[1];
    else if (line.includes("Test"))
      monkeys[cm].test["o"] = Number(line.split("by ")[1]);
    else if (line.includes("true"))
      monkeys[cm].test["t"] = Number(line.split("monkey")[1]);
    else if (line.includes("false"))
      monkeys[cm].test["f"] = Number(line.split("monkey")[1]);
  });
  return monkeys;
}

function doMath(op) {
  return Function(`'use strict'; return (${op})`)();
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
