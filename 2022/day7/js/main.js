"use strict";

const fs = require("fs");

async function part1(input) {
  let result = 0;
  const fs = buildFs(input);
  Object.keys(fs).forEach((dir) => {
    const dirSize = getDirSize(fs, dir);
    console.log(`${dir} size = ${dirSize}`);
    if (dirSize <= 100000) result += dirSize;
  });
  return result;
}

async function part2(input) {
  const total = 70000000;
  const need = 30000000;
  const fs = buildFs(input);
  Object.keys(fs).forEach((dir) => {
    const dirSize = getDirSize(fs, dir);
    fs[dir].size = dirSize;
  });
  const used = fs["/"].size;
  const unused = total - used;
  let smallest = Infinity;
  Object.keys(fs).forEach((dir) => {
    if (unused + fs[dir].size >= need && fs[dir].size <= smallest)
      smallest = fs[dir].size;
  });
  return smallest;
}

async function main() {
  const data = fs.readFileSync("../input.txt", { encoding: "utf8" });

  const result1 = await part1(data);
  console.log(`Part 1 = ${result1}`);
  const result2 = await part2(data);
  console.log(`Part 2 = ${result2}`);
}

function getDirSize(fs, dir) {
  let dirSize = 0;
  const contents = fs[dir].contents;
  contents.forEach((item) => {
    if (item.type === "dir") {
      const subDir = `${dir}${item.name}/`;
      dirSize += getDirSize(fs, subDir);
    } else if (item.type === "file") dirSize += item.size;
  });
  return dirSize;
}

function buildFs(input) {
  let currentDirectory = [];
  const fs = {};

  const lines = input.split("\n");

  let LSing = false;
  lines.forEach((line) => {
    if (line.indexOf("$") === 0) {
      LSing = false;
      if (line.indexOf("cd") === 2) {
        const goingTo = line.split(" ")[2];
        if (goingTo === "..") currentDirectory.pop();
        else if (goingTo === "/") currentDirectory.push("/");
        else currentDirectory.push(goingTo + "/");

        const newDir = currentDirectory.join("");
        if (!fs[newDir]) fs[newDir] = { contents: [] };
      } else if (line.indexOf("ls") == 2) {
        LSing = true;
      }
    } else if (LSing) {
      const dir = currentDirectory.join("");
      if (line.indexOf("dir") === 0) {
        fs[dir].contents.push({
          name: line.split(" ")[1],
          type: "dir",
        });
      } else {
        const [size, file] = line.split(" ");
        fs[dir].contents.push({
          name: file,
          size: Number(size),
          type: "file",
        });
      }
    }
  });
  return fs;
}

if (require.main === module) {
  main();
}

module.exports = {
  part1,
  part2,
};
