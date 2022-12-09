"use strict";

const fs = require("fs");

async function part1(input) {
  const status = new HTStatus(2);
  input.forEach((m) => {
    const movement = m.split(" ");
    status.move(movement[0], Number(movement[1]));
  });

  return status.getVisitedCount();
}

async function part2(input) {
  const status = new HTStatus(10);
  input.forEach((m) => {
    const movement = m.split(" ");
    status.move(movement[0], Number(movement[1]));
  });

  return status.getVisitedCount();
}

class HTStatus {
  constructor(k) {
    this.knots = [];
    for (let i = 0; i < k; i++) {
      this.knots.push([0, 0]);
    }
    this.TVisited = {};
    this.setTVisited(this.knots[this.knots.length - 1]);
  }

  getVisitedCount() {
    return Object.keys(this.TVisited).length;
  }

  move(d, s) {
    const head = this.knots[0];

    while (s > 0) {
      switch (d) {
        case "U":
          head[1] += 1;
          break;
        case "D":
          head[1] -= 1;
          break;
        case "R":
          head[0] += 1;
          break;
        case "L":
          head[0] -= 1;
      }

      for (let i = 1; i < this.knots.length; i++) {
        const k1 = this.knots[i - 1];
        const k2 = this.knots[i];
        const [hx, hy] = k1;
        const [tx, ty] = k2;

        if (!this.touching(k1, k2)) {
          /*console.log(
            `moving k${i} [${k2.join(",")}] because it's not touching k${
              i - 1
            } [${k1.join(",")}]`
          );*/
          if (hx !== tx && hy !== ty) {
            // must move diagonally
            this.knots[i][0] = hx - tx < 0 ? tx - 1 : tx + 1;
            this.knots[i][1] = hy - ty < 0 ? ty - 1 : ty + 1;
            /*console.log(
              `moving k${i} to [${this.knots[i].join(",")}] diagonally`
            );*/
          } else {
            // must move vertically or horizontally
            if (Math.abs(hx - tx) > 1) {
              // horizontally
              this.knots[i][0] = hx - tx < 0 ? tx - 1 : tx + 1;
              /*console.log(
                `moving k${i} to [${this.knots[i].join(",")}] horizontally`
              );*/
            } else if (Math.abs(hy - ty) > 1) {
              // vertically
              this.knots[i][1] = hy - ty < 0 ? ty - 1 : ty + 1;
              /*console.log(
                `moving k${i} to [${this.knots[i].join(",")}] vertically`
              );*/
            }
          }
          if (i === this.knots.length - 1) this.setTVisited(this.knots[i]);
        } else {
          /*console.log(
            `no movement required as k${i - 1} [${k1.join(
              ","
            )}] and k${i} [${k2.join(",")}] are touching`
          );*/
        }
      }
      s--;
    }
  }

  touching(k1, k2) {
    const [hx, hy] = k1;
    const [tx, ty] = k2;
    return Math.abs(hx - tx) <= 1 && Math.abs(hy - ty) <= 1;
  }

  setTVisited(k) {
    const loc = k.join(",");
    //console.log(`setting visited [${k.join(",")}]`);
    if (!this.TVisited[loc]) this.TVisited[loc] = 1;
  }
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
