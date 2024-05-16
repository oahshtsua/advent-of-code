# Day 1: No Time for a Taxicab

from collections import namedtuple

LEFT = {
    "North": "West",
    "South": "East",
    "West": "South",
    "East": "North",
}

RIGHT = {
    "North": "East",
    "South": "West",
    "West": "North",
    "East": "South",
}

DELTA = {
    "North": (0, 1),
    "South": (0, -1),
    "East": (1, 0),
    "West": (-1, 0),
}

Instruction = namedtuple("Instruction", ["dir", "dist"])


def part_one(instructions: list[Instruction]) -> int:
    x, y = 0, 0
    dir = "North"

    for entry in instructions:
        dir = RIGHT[dir] if entry.dir == "R" else LEFT[dir]
        dx, dy = DELTA[dir]
        x, y = x + (dx * entry.dist), y + (dy * entry.dist)

    return abs(x) + abs(y)


def part_two(instructions: list[Instruction]) -> int:
    x, y = 0, 0
    dir = "North"

    visited = set()

    for entry in instructions:
        dir = RIGHT[dir] if entry.dir == "R" else LEFT[dir]
        dx, dy = DELTA[dir]

        for _ in range(entry.dist):
            x, y = x + dx, y + dy
            if (x, y) in visited:
                return abs(x) + abs(y)
            visited.add((x, y))
    return -1


if __name__ == "__main__":
    with open("inputs/day01.txt", "r") as f:
        instructions = [
            Instruction(dir=entry[0], dist=int(entry[1:]))
            for entry in f.read().strip().split(", ")
        ]

    print("Part one: ", part_one(instructions))
    print("Part two: ", part_two(instructions))
