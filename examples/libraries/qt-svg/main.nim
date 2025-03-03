import seaqt/[qapplication, qsvgwidget]

proc main() =
  let
    _ = QApplication.create()
    w = QSvgWidget.create("../../../doc/logo.svg")
  w.show()

  discard QApplication.exec()

when isMainModule:
  main()
