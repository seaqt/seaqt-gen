import seaqt/[qapplication, qprintdialog, qpushbutton]

proc main() =
  let
    _ = QApplication.create()
    btn = QPushButton.create("QPrintSupport sample")

  btn.setFixedWidth(320)

  btn.onPressed(
    proc() =
      let dlg = QPrintDialog.create()
      dlg.show()
  )

  btn.show()

  discard QApplication.exec()

when isMainModule:
  main()
