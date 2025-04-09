import std/strformat, seaqt/[qapplication, qpushbutton]

proc main() =
  let
    _ = QApplication.create() # Initialize Qt
    btn = QPushButton.create("Hello world!")

  btn.setFixedWidth(320)

  var counter = 0

  btn.onPressed(
    proc() =
      counter += 1
      btn.setText(&"You have clicked the button {counter} time(s)")
  )

  btn.show()

  quit QApplication.exec().int

main()
