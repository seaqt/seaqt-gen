version = "0.1.0"
requires "nim >= 2.0", "unittest2"

task test515, "Run tests":
  for e in ["marshalling"]:
    exec "nim c --skipProjCfg -p:../gen/qt-5.15 -r " & e & "/main"

task test64, "Run tests":
  for e in ["marshalling"]:
    exec "nim c --skipProjCfg -p:../gen/qt-6.4 -r " & e & "/main"
