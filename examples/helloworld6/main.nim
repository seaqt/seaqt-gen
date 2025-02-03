import seeqt/[qapplication, qpushbutton]

import strformat

proc main() =
  let app = QApplication.create()

  let btn = QPushButton.create("Hello world!")
  btn.setFixedWidth(320)

  var counter = 0

  btn.onPressed(proc =
    counter += 1
    btn.setText(&"You have clicked the button {counter} time(s)")
  )

  btn.show()

  echo QApplication.exec()


main()
