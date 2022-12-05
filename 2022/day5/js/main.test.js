"use strict";

const assert = require("assert");
const { part1, part2 } = require("./main");

const input = `
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`;

async function test_part1() {
  const expected = "CMZ";
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2() {
  const expected = "MCD";
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
