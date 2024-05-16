# Day 1: Inverse Captcha


def matching_sum(sequence: str, step: int) -> int:
    sum = 0
    n = len(sequence)
    for i in range(n):
        if sequence[i] == sequence[(i + step) % n]:
            sum += int(sequence[i])
    return sum


def part_one(sequence: str) -> int:
    return matching_sum(sequence, 1)


def part_two(sequence: str) -> int:
    return matching_sum(sequence, len(sequence) // 2)


if __name__ == "__main__":
    with open("inputs/day01.txt", "r") as f:
        sequence = f.read().strip()

    print("Part one: ", part_one(sequence))
    print("Part two: ", part_two(sequence))
