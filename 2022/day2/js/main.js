"use strict";

const fs = require("fs");

const SCORES = {
  X: 1,
  Y: 2,
  Z: 3,
};

const PAIRS = {
  "A X": 3,
  "A Y": 6,
  "A Z": 0,
  "B X": 0,
  "B Y": 3,
  "B Z": 6,
  "C X": 6,
  "C Y": 0,
  "C Z": 3,
};

async function part1(input) {
  let totalScore = 0;

  input.forEach((round) => {
    const choice2 = round.split(" ")[1];
    const shapeScore = SCORES[choice2];
    const roundScore = PAIRS[round];
    totalScore += shapeScore + roundScore;
  });

  return totalScore;
}

async function part2(input) {
  let totalScore = 0;
  input.forEach((round) => {
    const choices = round.split(" ");
    const choice1 = choices[0];
    const choice2 = choices[1];
    let roundScore = 0;
    if (choice2 == "X") {
      // round score is 0
    } else if (choice2 == "Y") {
      roundScore = 3;
    } else if (choice2 == "Z") {
      roundScore = 6;
    }
    const shapeScore =
      SCORES[
        Object.keys(PAIRS)
          .filter(
            (p) => p.split(" ")[0] == choice1 && PAIRS[p] == roundScore
          )[0]
          .split(" ")[1]
      ];
    totalScore += roundScore + shapeScore;
  });
  return totalScore;
}

async function main() {
  const data = fs.readFileSync("../input.txt", { encoding: "utf8" });
  // @TODO - transform
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
