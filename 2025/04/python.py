def parse_input() -> list[str]:
    f = open("test_input.txt")
    content = f.read()
    return content.split()


def check_rolls():
    content = parse_input()
    window_len = 8

    adjacent = []
    for line in content:
        for char in line:


check_rolls()
