"use strict";

const assert = require("assert");

const input = null;

async function test_part1() {
  const expected = 0;
  const actual = await test_part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2() {
  const expected = 0;
  const actual = await test_part2(input);
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
