#include "qobjectdefs.h"
#define WORKAROUND_INNER_CLASS_DEFINITION_QMetaObject__Connection 1
#include <QMetaObject>
#include <QObject>

#include "../QtCore/gen_qmetaobject.h"
#include "../QtCore/gen_qobject.h"

#include <private/qobject_p.h>

extern "C" {

typedef void (*QObject_connectRawSlot_callback)(intptr_t, void **a);
typedef void (*QObject_connectRawSlot_release)(intptr_t);

QMetaObject__Connection *
QObject_connectRawSlot(QObject *sender, const char *signal, QObject *receiver,
                       intptr_t slot, QObject_connectRawSlot_callback callback,
                       QObject_connectRawSlot_release release,
                       Qt::ConnectionType type,
                       const QMetaObject *senderMetaObject) {
  // Mix between QFunctorSlotObject and QSlotObject where the signal is found by
  // signature but the slot is a type-erased value provided by the user who also
  // must unpack the arguments - relies on the same private connection method
  // as QML - loosely based on qmetaobject-rs:
  // https://github.com/woboq/qmetaobject-rs/blob/81e9ad1af8ea413a03bf17afe085d0b2e66ff3c4/qmetaobject/src/connections.rs#L126
  // More info: https://woboq.com/blog/how-qt-signals-slots-work-part2-qt5.html
  class QCSlotObject : public QtPrivate::QSlotObjectBase {
    intptr_t slot;
    QObject_connectRawSlot_callback callback;
    QObject_connectRawSlot_release release;

    static void impl(int which, QSlotObjectBase *this_, QObject *r, void **a,
                     bool *ret) {
      QCSlotObject *self = static_cast<QCSlotObject *>(this_);
      switch (which) {
      case Destroy:
        self->release(self->slot);
        delete self;
        break;
      case Call:
        self->callback(self->slot, a);
        break;
      case Compare: // not implemented
      case NumOperations:
        Q_UNUSED(ret);
      }
    }

  public:
    explicit QCSlotObject(intptr_t slot,
                          QObject_connectRawSlot_callback callback,
                          QObject_connectRawSlot_release release)
        : QSlotObjectBase(&impl), slot(slot), callback(callback),
          release(release) {}
  };
  auto slotObj = new QCSlotObject(slot, callback, release);
  auto signal_index = senderMetaObject->indexOfSignal(signal + 1);

#if QT_VERSION < QT_VERSION_CHECK(5, 15, 3)
  // https://github.com/qt/qtbase/commit/32d4e43f926c30bd1a8dd6e6744385d731908d06
  return new QMetaObject::Connection(
      QObjectPrivate::connect(sender, signal_index, slotObj, type));
#else
  return new QMetaObject::Connection(
      QObjectPrivate::connect(sender, signal_index, receiver, slotObj, type));
#endif
}
}
