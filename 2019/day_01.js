// Day 1: The Tyranny of the Rocket Equation

import fs from "node:fs";

const calculateFuel = (mass) => Math.floor(mass / 3) - 2;

const calculateTotalFuel = (mass) => {
  let fuel = calculateFuel(mass);
  return fuel <= 0 ? 0 : fuel + calculateTotalFuel(fuel);
};

const partOne = (masses) => {
  return masses.reduce((acc, curr) => acc + calculateFuel(curr), 0);
};

const partTwo = (masses) => {
  return masses.reduce((acc, curr) => acc + calculateTotalFuel(curr), 0);
};

const masses = fs.readFileSync("inputs/day01.txt", "utf8").trim().split("\n");
console.log("Part one: ", partOne(masses));
console.log("Part two: ", partTwo(masses));
