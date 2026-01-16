run_task('task_zero'); # install Python3 

run_task(
  'task_one', {
      'foo' : 'foo value',
      'bar' : 'bar value'
    }
);

run_task(
  'task_two', {
      'foo' : 'foo value',
      'bar' : 'bar value'
    }
);
