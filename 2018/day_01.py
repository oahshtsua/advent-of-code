# Day 1: Chronal Calibration


def part_one(sequence: list[int]) -> int:
    return sum(sequence)


def part_two(sequence: list[int]) -> int:
    idx = 0
    sum = 0
    cache = set()

    while True:
        sum += sequence[idx % len(sequence)]
        if sum in cache:
            break
        cache.add(sum)

        idx += 1
    return sum


if __name__ == "__main__":
    with open("inputs/day01.txt", "r") as f:
        sequence = [int(line.strip()) for line in f]

    print("Part one: ", part_one(sequence))
    print("Part two: ", part_two(sequence))
