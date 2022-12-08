"use strict";

const assert = require("assert");
const { part1, part2 } = require("./main");

const input = ["30373", "25512", "65332", "33549", "35390"];

async function test_part1(input) {
  const expected = 21;
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2(input) {
  const expected = 8;
  const actual = await part2(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test() {
  const data = input.map((r) => r.split(""));
  await test_part1(data);
  await test_part2(data);
}

test();
