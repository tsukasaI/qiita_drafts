def append_abc(s: str) -> str:
    a = append_a(s)
    b = append_b(a)
    c = append_c(b)
    return c

def append_a(s: str) -> str:
    return s + 'a'

def append_b(s: str) -> str:
    return s + 'b'

def append_c(s: str) -> str:
    return s + 'c'


def main():
    s = 'start'
    a = append_a(s)
    b = append_b(s)
    c = append_c(s)
    abc = append_abc(s)

if __name__ == '__main__':
    main()
