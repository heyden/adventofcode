"use strict";

const fs = require("fs");

async function part1(input) {
  let result = 0;

  for (let r = 0; r < input.length; r++) {
    for (let c = 0; c < input[r].length; c++) {
      if (
        r === 0 ||
        r === input.length - 1 ||
        c === 0 ||
        c === input[r].length - 1
      ) {
        result++;
        continue;
      }
      const left = input[r].slice(0, c);
      const right = input[r].slice(c + 1, input[r].length);
      const column = [];
      for (let row = 0; row < input.length; row++) {
        column.push(input[row][c]);
      }

      const top = column.slice(0, r);
      const bottom = column.slice(r + 1, column.length);

      const t = input[r][c];
      let visible = false;
      if (
        left.every((tree) => tree < t) ||
        right.every((tree) => tree < t) ||
        top.every((tree) => tree < t) ||
        bottom.every((tree) => tree < t)
      ) {
        result++;
        continue;
      }
    }
  }

  return result;
}

async function part2(input) {
  let result = 0;
  const length = input[0].length;
  for (let r = 0; r < input.length; r++) {
    for (let c = 0; c < input[r].length; c++) {
      const left = input[r].slice(0, c);
      const right = input[r].slice(c + 1, input[r].length);
      const column = [];
      for (let row = 0; row < input.length; row++) {
        column.push(input[row][c]);
      }
      const top = column.slice(0, r);
      const bottom = column.slice(r + 1, column.length);
      const t = input[r][c];

      const topScore = visibility(top.reverse(), t);
      const bottomScore = visibility(bottom, t);
      const leftScore = visibility(left.reverse(), t);
      const rightScore = visibility(right, t);
      const scenicScore = topScore * bottomScore * leftScore * rightScore;
      if (scenicScore > result) result = scenicScore;
    }
  }
  return result;
}

function visibility(view, t) {
  let count = 0;
  for (let i = 0; i < view.length; i++) {
    count++;
    if (view[i] >= t) {
      break;
    }
  }
  return count;
}

async function main() {
  const input = fs.readFileSync("../input.txt", { encoding: "utf8" });

  const data = input.split("\n").map((row) => {
    return row.split("").map((n) => Number(n));
  });

  const result1 = await part1(data);
  console.log(`Part 1 = ${result1}`);
  const result2 = await part2(data);
  console.log(`Part 2 = ${result2}`);
}

if (require.main === module) {
  main();
}

module.exports = {
  part1,
  part2,
};
