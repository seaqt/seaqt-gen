version = "0.1.0"
author = "The seaqt authors"
description = "seaqt examples"
license = "MIT"

requires "nim >= 2.0", "unittest2"

task examples_5_15, "Run tests":
  # These require interaction
  for e in ["helloworld", "qt-multimedia", "qt-printsupport", "qt-svg"]:
    exec "nim c nim-seaqt-5.15/" & e & "/main"

task examples_6_4, "Run tests":
  # These require interaction
  for e in ["helloworld", "qt-printsupport", "qt-svg", "qt-webengine"]:
    exec "nim c nim-seaqt-6.4/" & e & "/main"

task test_5_15, "Run tests":
  for e in ["marshalling", "qt-script"]:
    exec "nim c -r nim-seaqt-5.15/" & e & "/main"

task test_6_4, "Run tests":
  for e in ["marshalling"]:
    exec "nim c -r nim-seaqt-6.4/" & e & "/main"
