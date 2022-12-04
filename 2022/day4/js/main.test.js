"use strict";

const assert = require("assert");
const { part1, part2 } = require("./main");

const input = [
  "2-4,6-8",
  "2-3,4-5",
  "5-7,7-9",
  "2-8,3-7",
  "6-6,4-6",
  "2-6,4-8",
];

async function test_part1() {
  const expected = 2;
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2() {
  const expected = 0;
  const actual = await part2(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test() {
  await test_part1(input);
  await test_part2(input);
}

test();
