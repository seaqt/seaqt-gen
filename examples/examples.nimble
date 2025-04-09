version = "0.1.0"
author = "The seaqt authors"
description = "seaqt examples"
license = "MIT"

requires "nim >= 2.0", "unittest2"

task test_5_15, "Run tests":
  # These require interaction
  for e in ["helloworld"]:
    exec "nim c -f nim-seaqt-5.15/" & e & "/main"

task test_6_4, "Run tests":
  # These require interaction
  for e in ["helloworld"]:
    exec "nim c -f nim-seaqt-6.4/" & e & "/main"
