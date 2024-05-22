# Day 2: Bathroom Security


def part_one(sequence: list[str]) -> str:
    KEYPAD = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9],
    ]

    r, c = 1, 1

    ROW_MIN, ROW_MAX = 0, len(KEYPAD[0]) - 1
    COL_MIN, COL_MAX = 0, len(KEYPAD) - 1

    code = ""
    for moves in sequence:
        for move in moves:
            if move == "U":
                r = max(ROW_MIN, r - 1)
            elif move == "D":
                r = min(ROW_MAX, r + 1)
            elif move == "L":
                c = max(COL_MIN, c - 1)
            else:
                c = min(COL_MAX, c + 1)
        code += str(KEYPAD[r][c])
    return code


def part_two(sequence: list[str]) -> str:
    KEYPAD = [
        [0, 0, 1, 0, 0],
        [0, 2, 3, 4, 0],
        [5, 6, 7, 8, 9],
        [0, "A", "B", "C", 0],
        [0, 0, "D", 0, 0],
    ]

    r, c = 2, 0

    ROW_MIN, ROW_MAX = 0, len(KEYPAD[0]) - 1
    COL_MIN, COL_MAX = 0, len(KEYPAD) - 1

    code = ""
    for moves in sequence:
        for move in moves:
            if move == "U":
                r_next = max(ROW_MIN, r - 1)
                if KEYPAD[r_next][c] != 0:
                    r = r_next
            elif move == "D":
                r_next = min(ROW_MAX, r + 1)
                if KEYPAD[r_next][c] != 0:
                    r = r_next
            elif move == "L":
                c_next = max(COL_MIN, c - 1)
                if KEYPAD[r][c_next] != 0:
                    c = c_next
            else:
                c_next = min(COL_MAX, c + 1)
                if KEYPAD[r][c_next] != 0:
                    c = c_next
        code += str(KEYPAD[r][c])
    return code


if __name__ == "__main__":
    with open("inputs/day02.txt", "r") as f:
        sequence = [line.strip() for line in f]

    print("Part one: ", part_one(sequence))
    print("Part two: ", part_two(sequence))
