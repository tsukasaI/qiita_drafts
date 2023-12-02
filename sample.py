def main():
    # test and
    print("======and test: both true========")
    result = true_print() and true_print()
    print("result: ", result)

    print("======and test: lest false========")
    result = false_print() and true_print()
    print("result: ", result)

    print("======and test: right false========")
    result = true_print() and false_print()
    print("result: ", result)

    print("======and test: both false========")
    result = false_print() and false_print()
    print("result: ", result)

    # test or
    print("======or test: both true========")
    result = true_print() or true_print()
    print("result: ", result)

    print("======or test: lest false========")
    result = false_print() or true_print()
    print("result: ", result)

    print("======or test: right false========")
    result = true_print() or false_print()
    print("result: ", result)

    print("======or test: both false========")
    result = false_print() or false_print()
    print("result: ", result)



def true_print() -> bool:
    print("true_print called")
    return True

def false_print() -> bool:
    print("false_print called")
    return False
main()
