#!raku

my $msg = prompt("message: ");

if ".git-commit-prefix".IO ~~ :e {
  my $prefix = ".git-commit-prefix".IO.lines.head;
  $msg =  "{$prefix} {$msg}";
} 

task-run "commit my changes", "git-commit", %( message => $msg );
