version = "0.1.0"
requires "nim >= 2.0", "unittest2"

task test515, "Run tests":
  for e in ["marshalling"]:
    exec "nim c --skipProjCfg -p:../gen/nim-seaqt-5.15 -r " & e & "/main"

  # These require interaction
  for e in ["libraries/qt-multimedia", "libraries/qt-printsupport", "libraries/qt-svg"]:
    exec "nim c --skipProjCfg -p:../gen/nim-seaqt-5.15 " & e & "/main"


task test64, "Run tests":
  for e in ["helloworld", "marshalling"]:
    exec "nim c --skipProjCfg -p:../gen/nim-seaqt-6.4 -r " & e & "/main"

  # These require interaction
  for e in ["libraries/qt-multimedia", "libraries/qt-printsupport", "libraries/qt-svg", "libraries/qt-webengine"]:
    exec "nim c --skipProjCfg -p:../gen/nim-seaqt-6.4 " & e & "/main"
