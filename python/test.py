#!/usr/bin/env python3

import unittest
from find_occurrences import find_occurrence


class TestFindOccurence(unittest.TestCase):
    def test_search_all_lower_peter(self):
        textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
        subtext = "peter"
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "1,20,75")

    def test_search_with_upper_peter(self):
        textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
        subtext = "Peter"
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "1,20,75")

    def test_search_with_upper_peter_and_space(self):
        textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
        subtext = "Peter "
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "1,20")

    def test_search_with_all_lower_pick(self):
        textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
        subtext = "pick"
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "30,58")

    def test_search_with_all_lower_pi(self):
        textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
        subtext = "pi"
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "30,37,43,51,58")

    def test_search_no_output1(self):
        textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
        subtext = "z"
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "<No Output>")

    def test_search_no_output2(self):
        textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
        subtext = "Peterz"
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "<No Output>")

    def test_search_with_all_string_len_zero(self):
        textToSearch = ""
        subtext = ""
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "<No Output>")

    def test_search_with_search_string_len_zero(self):
        textToSearch = ""
        subtext = "z"
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "<No Output>")

    def test_search_with_sub_string_len_zero(self):
        textToSearch = "abc"
        subtext = ""
        result = find_occurrence(textToSearch, subtext)
        self.assertEqual(result, "<No Output>")


if __name__ == '__main__':
    unittest.main()
