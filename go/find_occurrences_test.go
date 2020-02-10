package main

import "testing"

func TestFindOccurrences(t *testing.T) {
    res := find_occurrences("Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", "peter")
    expectRes := "1,20,75"
    if res != expectRes {
       t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", "Peter")
    expectRes = "1,20,75"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", "Peter ")
    expectRes = "1,20"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", "pick")
    expectRes = "30,58"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", "pick ")
    expectRes = "<No Output>"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", "pi")
    expectRes = "30,37,43,51,58"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", "z")
    expectRes = "<No Output>"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!", "Peterz")
    expectRes = "<No Output>"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("", "")
    expectRes = "<No Output>"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("ab", "")
    expectRes = "<No Output>"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }

    res = find_occurrences("", "s")
    expectRes = "<No Output>"
    if res != expectRes {
        t.Errorf("find_occurrences was incorrect, got: %s want: %s.", res, expectRes)
    }
}