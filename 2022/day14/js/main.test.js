"use strict";

const assert = require("assert");
const { part1, part2 } = require("./main");

const test_data = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`;

async function test_part1(input) {
  const expected = 24;
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2(input) {
  const expected = 93;
  const actual = await part2(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test() {
  const input = test_data.split("\n");
  await test_part1(input);
  await test_part2(input);
}

test();
