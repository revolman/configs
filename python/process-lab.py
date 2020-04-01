from multiprocessing import Queue, Process
import time


def f(q):
    x = q.get()
    print(f"Process number {x}, sleeps for {x} seconds")
    time.sleep(x)
    print(f"Process number {x} finished")


q = Queue()

for i in range(10):
    q.put(i)
    i = Process(target=f, args=[q])
    i.start()

print("Main process joins on queue")
i.join()
time.sleep(2)
print("Main Process finished")
