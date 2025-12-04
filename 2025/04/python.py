def parse_input() -> list[str]:
    f = open("input.txt")
    content = f.read()
    lines = content.split()
    out = []
    for line in lines:
        out.append( list(line))
    return out

def print_grid(grid):
    for line in grid:
        print("".join(line))
    

def check_rolls():
    content = parse_input()
    window_len = 8

    # ..@@.@@@@.
    # @@@.@.@.@@
    # @@@@@.@.@@
    # @.@@@@..@.

    sum = 0

    grid = []
    for i in range(len(content)):
        for j in range(len(content[0])):
            adjacent: list[str]= []
            # check row on top
            if content[i][j] == '@':
                if i != 0:
                    if j != 0:
                        adjacent.append(content[i-1][j-1])
                    adjacent.append(content[i-1][j])
                    if j != len(content[0]) - 1:
                        adjacent.append(content[i-1][j+1])

                # check same row

                if j != 0:
                    adjacent.append(content[i][j-1])

                if j != len(content[0]) - 1:
                    adjacent.append(content[i][j+1])

                # check last row
                if i != len(content)-1:
                    if j != 0:
                        adjacent.append(content[i+1][j-1])
                    adjacent.append(content[i+1][j])
                    if j != len(content[0]) - 1:
                        adjacent.append(content[i+1][j+1])

                if adjacent.count('@') + adjacent.count('X') < 4:
                    content[i][j] = 'X'
                    print(adjacent)
                    sum += 1

    print_grid(content)
    print(sum)


check_rolls()

#print(parse_input())

