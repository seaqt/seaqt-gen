import seaqt/[qapplication, qscriptengine, qscriptvalue]

proc main() =
  let
    _ = QApplication.create()
    inputProgram = "1 + 2"
    eng = QScriptEngine.create()
  echo inputProgram, " = ", eng.evaluate(inputProgram).toNumber()

when isMainModule:
  main()
