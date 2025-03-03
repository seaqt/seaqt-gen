import seaqt/[qapplication, qabstractlistmodel, qlistview, qvariant, QtCore/gen_qnamespace]

type Model = ref object of VirtualQAbstractListModel

method rowCount(self: Model, idx: QModelIndex): cint =
  1000

method data(self: Model, idx: QModelIndex, role: cint): QVariant =
  if not idx.isValid():
    return QVariant.create()

  if role == ItemDataRoleEnum.DisplayRole:
    QVariant.create("this is row " & $idx.row())
  else:
    QVariant.create()

proc main() =
  let
    _ = QApplication.create()

    model = Model()

  QAbstractListModel.create(model)

  let v = QListView.create()
  v.setModel(model[])
  v.show()

  discard QApplication.exec()

when isMainModule:
  main()
