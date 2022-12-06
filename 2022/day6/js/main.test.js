"use strict";

const assert = require("assert");
const { part1, part2 } = require("./main");

async function test_part1(input) {
  const expected = 7;
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2(input) {
  const expected = 19;
  const actual = await part2(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test() {
  let input = "mjqjpqmgbljsphdztnvjfqwrcgsmlb";
  await test_part1(input);
  input = "mjqjpqmgbljsphdztnvjfqwrcgsmlb";
  await test_part2(input);
}

test();
