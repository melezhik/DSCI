#!/bin/python

print(task_var("foo"))
print(task_var("bar"))

print("=== done ====")

update_state({
  'out1' : 'out1 value',
  'out2' : 'out2 value'
})
