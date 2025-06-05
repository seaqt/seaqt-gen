import seaqt/[qapplication, qwebengineview, qurl]

proc main() =
  let
    _ = QApplication.create()
    w = QWebEngineView.create()

  w.load(QUrl.create("https://www.eff.org/"))

  w.show()

  discard QApplication.exec()

when isMainModule:
  main()
