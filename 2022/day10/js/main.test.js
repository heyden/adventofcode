"use strict";

const assert = require("assert");
const fs = require("fs");
const { part1, part2 } = require("./main");

async function test_part1(input) {
  const expected = 13140;
  const actual = await part1(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test_part2(input) {
  const expected = 0;
  const actual = await part2(input);
  assert.equal(
    actual,
    expected,
    `actual: [${actual}] != expected [${expected}]`
  );
}

async function test() {
  const data = fs.readFileSync("../input.test.txt", { encoding: "utf8" });
  const input = data.split("\n");
  input.reverse();
  await test_part1(input);
  await test_part2(input);
}

test();
