def parse_input() -> list[list[int]]:
    f = open("input.txt")
    content = f.read()
    # content = "3475855376536553445244645449948436834663673457444335454425154463544745554434526565475244677445573464"
    banks: list[list[int]] = []
    for line in content.split():
        banks.append([int(num) for num in str(line)])

    return banks

def part1():
    input = parse_input()
    def parse_bank(bank: list[int]):
        high1: list[int] = [0, 0]
        high2: list[int] = [0, 1]

        for i in range(len(bank) -1):
            if bank[i] > high1[0]:
                high1[0] = bank[i]
                high1[1] = i

        for i in range(high1[1] + 1, len(bank)):
            if bank[i] > high2[0]:
                high2[0] = bank[i]

        return int(str(high1[0]) + str(high2[0]))

    sum = 0
    for bank in input:
        sum += parse_bank(bank)

    print(sum)

part1()

