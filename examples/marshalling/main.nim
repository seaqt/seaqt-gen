import
  std/[sequtils, tables],
  unittest2,
  seaqt/[
    qaction, qapplication, qcheckbox, qfile, qinputdialog, qjsonobject, qkeysequence,
    qsize, qvariant, qversionnumber,
  ]

suite "marshalling":
  setup:
    let _ = QApplication.create()

  test "bool":
    let b = QCheckBox.create()
    b.setChecked(true)

    assert b.isChecked()

  test "int":
    let s = QSize.create()
    s.setWidth(128)
    assert s.width == 128

  test "string":
    let w = QWidget.create()
    w.setToolTip("Sample text")
    assert w.toolTip() == "Sample text"

  test "intlist":
    let expect = @[cint 10, 20, 30, 40, 55]
    let s = QVersionNumber.create(expect)
    assert s.segments() == expect

  test "stringlist":
    let c = QInputDialog.create()
    let expect = @["foo", "bar", "baz", "quux"]
    c.setComboBoxItems(expect)
    assert c.comboBoxItems() == expect

  test "seq of q":
    let expect = ["F1", "F2", "F3"].mapIt(QKeySequence.fromString(it))
    let c = QAction.create()
    c.setShortcuts(expect)

    assert expect.mapIt(it.toString()) == c.shortcuts().mapIt(it.toString())

  test "string":
    let input = "foo bar baz"
    let ba = QFile.encodeName(input)
    assert QFile.decodeName(ba) == input

  test "empty string":
    let input = ""
    let ba = QFile.encodeName(input)
    assert QFile.decodeNAme(ba) == input

  test "Table":
    var input: Table[string, QVariant]
    input["foo"] = QVariant.create("FOO")
    input["bar"] = QVariant.create("BAR")
    input["baz"] = QVariant.create("BAZ")

    let qtobj = QJsonObject.fromVariantMap(input)
    let got = qtobj.toVariantMap()

    assert got.len == input.len(), $got.len()

    for k in input.keys():
      assert got[k].toString() == input[k].toString()
