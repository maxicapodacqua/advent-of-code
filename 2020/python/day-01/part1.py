import sys
import os

def process(input: str):
    lines = input.split()
    linesSet = set(lines)
    print(linesSet)
    for line in lines:
        elementToFind = 2020 - int(line)
        if str(elementToFind) in linesSet:
            print("found element!", elementToFind, str(elementToFind) in linesSet)
            return elementToFind * int(line)
    return False

if __name__ == "__main__":
    path = os.path.dirname(os.path.realpath(__file__)) + "/input.txt"
    input = open(path, "r")
    solution = process(input.read())
    print("Solution", solution)
