import re

with open("data.txt") as raw_file:
    data = raw_file.read()
    # the idea was probably to build a tokenizer, but screw it its 11pm regex it is
    regex = r"mul\([0-9]*,[0-9]*\)"
    matches = re.findall(regex, data)

    total = 0
    for expression in matches:
        # turns (for example) mul(1,2) into 1,2
        comma_separated_nums = expression.replace("mul(", "")[:-1]
        nums_as_strings = comma_separated_nums.split(",")
        num1, num2 = nums_as_strings
        product = int(num1) * int(num2)
        total += product
        print(f"Multiplying {num1} by {num2} resulting in {product}")

    print(f"End result: {total}")

