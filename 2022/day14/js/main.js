"use strict";

const fs = require("fs");

async function part1(input) {
  let result = 0;
  const grid = buildGrid(input);
  const start = "500,0";

  const lowestPoint = Math.max(
    ...Object.keys(grid).map((k) => Number(k.split(",")[1]))
  );

  let done = false;
  while (!done) {
    let cp = asPoint(start);
    let falling = true;
    while (falling) {
      if (cp[1] > lowestPoint) {
        done = true;
        break;
      } else {
        const pl = grid[point([cp[0] - 1, cp[1] + 1])];
        const pr = grid[point([cp[0] + 1, cp[1] + 1])];
        const b = grid[point([cp[0], cp[1] + 1])];
        if (!b) cp[1] += 1;
        else if (!pl) cp = [cp[0] - 1, cp[1] + 1];
        else if (!pr) cp = [cp[0] + 1, cp[1] + 1];
        else {
          grid[point(cp)] = "o";
          falling = false;
        }
      }
    }
    if (!done) result++;
  }
  return result;
}

async function part2(input) {
  let result = 0;
  const grid = buildGrid(input);
  const start = "500,0";

  const floor =
    Math.max(...Object.keys(grid).map((k) => Number(k.split(",")[1]))) + 2;

  let done = false;
  while (!done) {
    let cp = asPoint(start);
    let falling = true;
    while (falling) {
      if (cp[1] + 1 === floor) {
        grid[point(cp)] = "o";
        falling = false;
      } else {
        const pl = grid[point([cp[0] - 1, cp[1] + 1])];
        const pr = grid[point([cp[0] + 1, cp[1] + 1])];
        const b = grid[point([cp[0], cp[1] + 1])];
        if (!b) cp[1] += 1;
        else if (!pl) cp = [cp[0] - 1, cp[1] + 1];
        else if (!pr) cp = [cp[0] + 1, cp[1] + 1];
        else if (cp[1] === 0 && grid[point([cp[0], cp[1] + 1])]) {
          done = true;
          falling = false;
        } else {
          grid[point(cp)] = "o";
          falling = false;
        }
      }
    }
    result++;
  }
  draw(grid);
  return result;
}

function buildGrid(input) {
  const grid = {};
  input.forEach((d) => {
    const points = d.split(" -> ");
    for (let i = 0; i < points.length - 1; i++) {
      const start = points[i].split(",").map((n) => Number(n));
      const dest = points[i + 1].split(",").map((n) => Number(n));
      const move = [0, 0];
      let diff = 0;
      if (start[0] === dest[0]) {
        diff = dest[1] - start[1];
        move[1] = Math.sign(dest[1] - start[1]);
        diff = Math.abs(diff);
      } else {
        diff = dest[0] - start[0];
        move[0] = Math.sign(dest[0] - start[0]);
        diff = Math.abs(diff);
      }
      const cp = [...start];
      grid[point(cp)] = "#";
      for (let m = 1; m <= diff; m++) {
        cp[0] += move[0];
        cp[1] += move[1];
        grid[point(cp)] = "#";
      }
    }
  });

  return grid;
}

function draw(grid) {
  const o = Object.keys(grid).sort((a, b) => {
    const [ax, ay] = a.split(",");
    const [bx, by] = b.split(",");

    if (ax < bx) return -1;
    else if (ax > bx) return 1;
    else if (ax === bx) {
      if (ay < by) return -1;
      else if (ay > by) return 1;
      else return 0;
    }
  });

  let left = Number(o[0].split(",")[0]);
  let right = Number(o[o.length - 1].split(",")[0]);
  const lowestPoint = Math.max(
    ...Object.keys(grid).map((k) => Number(k.split(",")[1]))
  );

  for (let i = 0; i <= lowestPoint; i++) {
    let s = "";
    for (let j = left; j <= right; j++) {
      s += grid[`${j},${i}`] ? grid[`${j},${i}`] : ".";
    }
    console.log(s);
  }
}

function point(s) {
  return `${s[0]},${s[1]}`;
}

function asPoint(s) {
  return s.split(",").map((n) => Number(n));
}

async function main() {
  const data = fs.readFileSync("../input.txt", { encoding: "utf8" });
  const input = data.split("\n");

  const result1 = await part1(input);
  console.log(`Part 1 = ${result1}`);
  const result2 = await part2(input);
  console.log(`Part 2 = ${result2}`);
}

if (require.main === module) {
  main();
}

module.exports = {
  part1,
  part2,
};
