#!/usr/bin/env python3


def find_occurrence(textToSearch, subText):
    if textToSearch == "" or subText == "":
        return "<No Output>"

    ans = ""
    textToSearchLower = textToSearch.lower()
    subTextLower = subText.lower()

    if subTextLower in textToSearchLower:
        ch = subTextLower[0]
        for i in range(len(textToSearchLower)):
            if textToSearchLower[i] == ch:
                # if the first character is same, compare the rest characters
                if textToSearchLower[i:i+len(subTextLower)] == subTextLower:
                    ans += str(i+1) + ","
                    i += len(subTextLower)
        ans = ans[:-1]
    else:
        ans = "<No Output>"

    return ans
