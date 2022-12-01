"use strict";

const assert = require("assert");

const { part1, part2 } = require("./main");

const input = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`;

async function test_part1() {
  const expected = 24000;
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2() {
  const expected = 45000;
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
