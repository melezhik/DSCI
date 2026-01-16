#!python

d = config()

print("===")

print("param foo=", d["foo"])
print("param foo2=", d["foo2"])

print("===")


hello = task_var("hello")
print("task var hello=", hello)
