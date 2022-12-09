"use strict";

const assert = require("assert");
const { part1, part2 } = require("./main");

const input = ["R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"];
const input2 = ["R 5", "U 8", "L 8", "D 3", "R 17", "D 10", "L 25", "U 20"];

async function test_part1(expected, input) {
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2(expected, input) {
  const actual = await part2(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test() {
  //await test_part1(13, input);
  //await test_part2(1, input);
  await test_part2(36, input2);
}

test();
