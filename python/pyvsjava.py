import sys


class Hello():
   
    def hello(self, name):
        self.name = name
        return f"Hello {self.name}!"

if __name__ == "__main__":
    h = Hello()
    arg1 = sys.argv[1]
    print(h.hello(arg1))
    