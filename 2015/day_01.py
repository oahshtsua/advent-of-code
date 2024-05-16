# Day 1: Not Quite Lisp


def part_one(instructions: str) -> int:
    curr_floor = 0
    for instruction in instructions:
        if instruction == "(":
            curr_floor += 1
        else:
            curr_floor -= 1
    return curr_floor


def part_two(instructions: str) -> int:
    idx = 0
    curr_floor = 0
    while idx < len(instructions):
        if instructions[idx] == "(":
            curr_floor += 1
        else:
            curr_floor -= 1

        if curr_floor < 0:
            break
        idx += 1

    return idx + 1


if __name__ == "__main__":
    with open("inputs/day01.txt", "r") as f:
        instructions = f.read().strip()

    print("Part one: ", part_one(instructions))
    print("Part two: ", part_two(instructions))
