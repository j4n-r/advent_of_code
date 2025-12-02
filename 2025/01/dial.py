

def parse_input():
    f = open("test_input.txt")
    contents = f.read().split()
    ops: list[tuple[str,int]] = [] 
    for op in contents:
        ops.append((op[0], int(op[1:])))
    return ops

def part1():
    ops = parse_input()
    dial = 50
    counter = 0

    for op in ops:
        if op[0] == 'L':
            dial -= op[1]
            dial = dial % 100 
        if op[0] == 'R':
            dial += op[1]
            dial = dial % 100 
        if dial == 0:
            counter += 1
    print(counter)

def part2():
    def modulo():

        pass
    ops = parse_input()
    dial = 50
    dial2 = 50
    counter = 0

    for op in ops:
        if op[0] == 'L':
            dial -= op[1]
            dial2 += op[1]
            dial = dial % 100 
        if op[0] == 'R':
            dial += op[1]
            dial2 += op[1]
            dial = dial % 100 
        if dial == 0:
            counter += 1
    counter += dial2 // 100
    print(counter)
    print(dial2)


part2()

