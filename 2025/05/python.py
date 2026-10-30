def parse_input():
    f = open("input.txt")
    contents = f.read()
    ranges: list[list[int]] = []
    ids: list[int] = []
    for line in contents.split():
        if "-" in line:
            split_line = [int(num) for num in line.split('-')]
            ranges.append(split_line)
        else:
            ids.append(int(line))

    return (ranges, ids)

def main():
    ranges, ids = parse_input()
    for id in ids:
        if id >= 

main()
                    
                
