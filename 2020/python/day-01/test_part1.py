import unittest
from part1 import process

class TestPart1(unittest.TestCase):
    def test_basic(self):
        input = '''1721
979
366
299
675
1456'''
        self.assertEqual(514579, process(input))

if __name__ == '__main__':
    unittest.main()
