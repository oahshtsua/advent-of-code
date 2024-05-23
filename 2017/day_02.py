# Day 2: Corruption Checksum

import itertools


def part_one(spreadsheet: list[list[int]]) -> int:
    checksum = 0
    for row in spreadsheet:
        min, max = row[0], row[0]
        for item in row:
            if item < min:
                min = item
            if item > max:
                max = item
        checksum += max - min
    return checksum


def part_two(spreadsheet: list[list[int]]) -> int:
    checksum = 0
    for row in spreadsheet:
        for x, y in itertools.combinations(row, 2):
            if x % y == 0 or y % x == 0:
                checksum += x // y if x > y else y // x
    return checksum


if __name__ == "__main__":
    with open("inputs/day02.txt", "r") as f:
        spreadsheet = [[int(val) for val in line.split()] for line in f]

    print("Part one: ", part_one(spreadsheet))
    print("Part two: ", part_two(spreadsheet))
