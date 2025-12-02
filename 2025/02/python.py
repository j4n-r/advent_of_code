def parse_input()-> list[list[str]]:
    # f = open("test_input.txt")
    f = open("input.txt")
    content = f.read()
    ranges = content.strip().split(",")

    parsed_id_ranges: list[list[str]] = []
    for id_range in ranges:
        parsed_id_ranges.append(id_range.split("-"))

    return parsed_id_ranges


def part1():
    def is_valid(string: str):
        length = len(string)
        if length % 2 != 0:
            return True
        return string[0:length//2] != string[length//2:length]
        
    parsed_id_ranges = parse_input()

    valid_ids: list[int] = []
    for id_range in parsed_id_ranges:
        for id in range(int(id_range[0]), int(id_range[1])+1):
            valid = is_valid(str(id))
            if not valid:
                valid_ids.append(id)

    print(sum(valid_ids))

def part2():
    def split_chunks(id_split: list[str], chunk_size: int):
        chunks: list[list[str]] = []
        for i in range(0, len(id_split), chunk_size):
            chunks.append(id_split[i:i + chunk_size])
        return chunks

    def is_valid(id: str):
        id_split = list(id)
        length = len(id)
        dividers: list[int] = []
        for i in range(1, length):
            if length % i == 0:
                dividers.append(i)

        chunks = []
        for divider in dividers:
            chunks.append(split_chunks(id_split, divider))


        for chunk in chunks:
            if  all(x == chunk[0] for x in chunk):
                # print(chunk)
                return False
        return True

            
    input  = parse_input()
    # input = [["122223","122224"]]
    not_valid_ids: list[int] = []
    # print(input)
    for id_range in input:
        for id in range(int(id_range[0]), int(id_range[1])+1):
            valid = is_valid(str(id))
            if not valid:
                not_valid_ids.append(id)

    print(not_valid_ids)

    print(sum(not_valid_ids))

part2()
