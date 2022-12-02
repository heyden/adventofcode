"use strict";

const assert = require("assert");
const { part1, part2 } = require("./main");

const input = ["A Y", "B X", "C Z"];

async function test_part1() {
  const expected = 15;
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2() {
  const expected = 12;
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
