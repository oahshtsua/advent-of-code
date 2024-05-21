# Day 2: I Was Told There Would Be No Math

import heapq


def part_one(dimensions: list[tuple[int, int, int]]) -> int:
    def calculate_area(l: int, w: int, h: int) -> int:
        surface_area = 2 * (l * w + w * h + h * l)
        smallest_sides = heapq.nsmallest(2, (l, w, h))
        smallest_side_area = smallest_sides[0] * smallest_sides[1]
        return surface_area + smallest_side_area

    total = 0
    for l, w, h in dimensions:
        total += calculate_area(l, w, h)
    return total


def part_two(dimensions: list[tuple[int, int, int]]) -> int:
    def calculate_ribbon_length(l: int, w: int, h: int) -> int:
        smallest_sides = heapq.nsmallest(2, (l, w, h))
        perimeter = 2 * sum(smallest_sides)
        volume = l * w * h
        return perimeter + volume

    total = 0
    for l, w, h in dimensions:
        total += calculate_ribbon_length(l, w, h)
    return total


if __name__ == "__main__":
    with open("inputs/day02.txt", "r") as f:
        dimensions = []
        for line in f:
            l, w, h = line.strip().split("x")
            dimensions.append((int(l), int(w), int(h)))

    print("Part one: ", part_one(dimensions))
    print("Part two: ", part_two(dimensions))
