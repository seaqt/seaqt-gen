#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDesktopWidget>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QRect>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <qdesktopwidget.h>
#include "gen_qdesktopwidget.h"

#ifdef __cplusplus
extern "C" {
#endif

void miqt_exec_callback_QDesktopWidget_Resized(intptr_t, int);
void miqt_exec_callback_QDesktopWidget_WorkAreaResized(intptr_t, int);
void miqt_exec_callback_QDesktopWidget_ScreenCountChanged(intptr_t, int);
void miqt_exec_callback_QDesktopWidget_PrimaryScreenChanged(intptr_t);
void miqt_exec_callback_QDesktopWidget_ResizeEvent(QDesktopWidget*, intptr_t, QResizeEvent*);
int miqt_exec_callback_QDesktopWidget_DevType(const QDesktopWidget*, intptr_t);
void miqt_exec_callback_QDesktopWidget_SetVisible(QDesktopWidget*, intptr_t, bool);
QSize* miqt_exec_callback_QDesktopWidget_SizeHint(const QDesktopWidget*, intptr_t);
QSize* miqt_exec_callback_QDesktopWidget_MinimumSizeHint(const QDesktopWidget*, intptr_t);
int miqt_exec_callback_QDesktopWidget_HeightForWidth(const QDesktopWidget*, intptr_t, int);
bool miqt_exec_callback_QDesktopWidget_HasHeightForWidth(const QDesktopWidget*, intptr_t);
QPaintEngine* miqt_exec_callback_QDesktopWidget_PaintEngine(const QDesktopWidget*, intptr_t);
bool miqt_exec_callback_QDesktopWidget_Event(QDesktopWidget*, intptr_t, QEvent*);
void miqt_exec_callback_QDesktopWidget_MousePressEvent(QDesktopWidget*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QDesktopWidget_MouseReleaseEvent(QDesktopWidget*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QDesktopWidget_MouseDoubleClickEvent(QDesktopWidget*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QDesktopWidget_MouseMoveEvent(QDesktopWidget*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QDesktopWidget_WheelEvent(QDesktopWidget*, intptr_t, QWheelEvent*);
void miqt_exec_callback_QDesktopWidget_KeyPressEvent(QDesktopWidget*, intptr_t, QKeyEvent*);
void miqt_exec_callback_QDesktopWidget_KeyReleaseEvent(QDesktopWidget*, intptr_t, QKeyEvent*);
void miqt_exec_callback_QDesktopWidget_FocusInEvent(QDesktopWidget*, intptr_t, QFocusEvent*);
void miqt_exec_callback_QDesktopWidget_FocusOutEvent(QDesktopWidget*, intptr_t, QFocusEvent*);
void miqt_exec_callback_QDesktopWidget_EnterEvent(QDesktopWidget*, intptr_t, QEvent*);
void miqt_exec_callback_QDesktopWidget_LeaveEvent(QDesktopWidget*, intptr_t, QEvent*);
void miqt_exec_callback_QDesktopWidget_PaintEvent(QDesktopWidget*, intptr_t, QPaintEvent*);
void miqt_exec_callback_QDesktopWidget_MoveEvent(QDesktopWidget*, intptr_t, QMoveEvent*);
void miqt_exec_callback_QDesktopWidget_CloseEvent(QDesktopWidget*, intptr_t, QCloseEvent*);
void miqt_exec_callback_QDesktopWidget_ContextMenuEvent(QDesktopWidget*, intptr_t, QContextMenuEvent*);
void miqt_exec_callback_QDesktopWidget_TabletEvent(QDesktopWidget*, intptr_t, QTabletEvent*);
void miqt_exec_callback_QDesktopWidget_ActionEvent(QDesktopWidget*, intptr_t, QActionEvent*);
void miqt_exec_callback_QDesktopWidget_DragEnterEvent(QDesktopWidget*, intptr_t, QDragEnterEvent*);
void miqt_exec_callback_QDesktopWidget_DragMoveEvent(QDesktopWidget*, intptr_t, QDragMoveEvent*);
void miqt_exec_callback_QDesktopWidget_DragLeaveEvent(QDesktopWidget*, intptr_t, QDragLeaveEvent*);
void miqt_exec_callback_QDesktopWidget_DropEvent(QDesktopWidget*, intptr_t, QDropEvent*);
void miqt_exec_callback_QDesktopWidget_ShowEvent(QDesktopWidget*, intptr_t, QShowEvent*);
void miqt_exec_callback_QDesktopWidget_HideEvent(QDesktopWidget*, intptr_t, QHideEvent*);
bool miqt_exec_callback_QDesktopWidget_NativeEvent(QDesktopWidget*, intptr_t, struct miqt_string, void*, long*);
void miqt_exec_callback_QDesktopWidget_ChangeEvent(QDesktopWidget*, intptr_t, QEvent*);
int miqt_exec_callback_QDesktopWidget_Metric(const QDesktopWidget*, intptr_t, int);
void miqt_exec_callback_QDesktopWidget_InitPainter(const QDesktopWidget*, intptr_t, QPainter*);
QPaintDevice* miqt_exec_callback_QDesktopWidget_Redirected(const QDesktopWidget*, intptr_t, QPoint*);
QPainter* miqt_exec_callback_QDesktopWidget_SharedPainter(const QDesktopWidget*, intptr_t);
void miqt_exec_callback_QDesktopWidget_InputMethodEvent(QDesktopWidget*, intptr_t, QInputMethodEvent*);
QVariant* miqt_exec_callback_QDesktopWidget_InputMethodQuery(const QDesktopWidget*, intptr_t, int);
bool miqt_exec_callback_QDesktopWidget_FocusNextPrevChild(QDesktopWidget*, intptr_t, bool);
bool miqt_exec_callback_QDesktopWidget_EventFilter(QDesktopWidget*, intptr_t, QObject*, QEvent*);
void miqt_exec_callback_QDesktopWidget_TimerEvent(QDesktopWidget*, intptr_t, QTimerEvent*);
void miqt_exec_callback_QDesktopWidget_ChildEvent(QDesktopWidget*, intptr_t, QChildEvent*);
void miqt_exec_callback_QDesktopWidget_CustomEvent(QDesktopWidget*, intptr_t, QEvent*);
void miqt_exec_callback_QDesktopWidget_ConnectNotify(QDesktopWidget*, intptr_t, QMetaMethod*);
void miqt_exec_callback_QDesktopWidget_DisconnectNotify(QDesktopWidget*, intptr_t, QMetaMethod*);
#ifdef __cplusplus
} /* extern C */
#endif

class MiqtVirtualQDesktopWidget final : public QDesktopWidget {
public:

	MiqtVirtualQDesktopWidget(): QDesktopWidget() {};

	virtual ~MiqtVirtualQDesktopWidget() override = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ResizeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void resizeEvent(QResizeEvent* e) override {
		if (handle__ResizeEvent == 0) {
			QDesktopWidget::resizeEvent(e);
			return;
		}
		
		QResizeEvent* sigval1 = e;

		miqt_exec_callback_QDesktopWidget_ResizeEvent(this, handle__ResizeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ResizeEvent(QResizeEvent* e) {

		QDesktopWidget::resizeEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DevType = 0;

	// Subclass to allow providing a Go implementation
	virtual int devType() const override {
		if (handle__DevType == 0) {
			return QDesktopWidget::devType();
		}
		

		int callback_return_value = miqt_exec_callback_QDesktopWidget_DevType(this, handle__DevType);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_DevType() const {

		return QDesktopWidget::devType();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetVisible = 0;

	// Subclass to allow providing a Go implementation
	virtual void setVisible(bool visible) override {
		if (handle__SetVisible == 0) {
			QDesktopWidget::setVisible(visible);
			return;
		}
		
		bool sigval1 = visible;

		miqt_exec_callback_QDesktopWidget_SetVisible(this, handle__SetVisible, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetVisible(bool visible) {

		QDesktopWidget::setVisible(visible);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__SizeHint == 0) {
			return QDesktopWidget::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QDesktopWidget_SizeHint(this, handle__SizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeHint() const {

		return new QSize(QDesktopWidget::sizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumSizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSizeHint() const override {
		if (handle__MinimumSizeHint == 0) {
			return QDesktopWidget::minimumSizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QDesktopWidget_MinimumSizeHint(this, handle__MinimumSizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MinimumSizeHint() const {

		return new QSize(QDesktopWidget::minimumSizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int heightForWidth(int param1) const override {
		if (handle__HeightForWidth == 0) {
			return QDesktopWidget::heightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QDesktopWidget_HeightForWidth(this, handle__HeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_HeightForWidth(int param1) const {

		return QDesktopWidget::heightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HasHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasHeightForWidth() const override {
		if (handle__HasHeightForWidth == 0) {
			return QDesktopWidget::hasHeightForWidth();
		}
		

		bool callback_return_value = miqt_exec_callback_QDesktopWidget_HasHeightForWidth(this, handle__HasHeightForWidth);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HasHeightForWidth() const {

		return QDesktopWidget::hasHeightForWidth();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEngine = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintEngine* paintEngine() const override {
		if (handle__PaintEngine == 0) {
			return QDesktopWidget::paintEngine();
		}
		

		QPaintEngine* callback_return_value = miqt_exec_callback_QDesktopWidget_PaintEngine(this, handle__PaintEngine);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintEngine* virtualbase_PaintEngine() const {

		return QDesktopWidget::paintEngine();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__Event == 0) {
			return QDesktopWidget::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QDesktopWidget_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* event) {

		return QDesktopWidget::event(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MousePressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mousePressEvent(QMouseEvent* event) override {
		if (handle__MousePressEvent == 0) {
			QDesktopWidget::mousePressEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_MousePressEvent(this, handle__MousePressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MousePressEvent(QMouseEvent* event) {

		QDesktopWidget::mousePressEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseReleaseEvent(QMouseEvent* event) override {
		if (handle__MouseReleaseEvent == 0) {
			QDesktopWidget::mouseReleaseEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_MouseReleaseEvent(this, handle__MouseReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseReleaseEvent(QMouseEvent* event) {

		QDesktopWidget::mouseReleaseEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseDoubleClickEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseDoubleClickEvent(QMouseEvent* event) override {
		if (handle__MouseDoubleClickEvent == 0) {
			QDesktopWidget::mouseDoubleClickEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_MouseDoubleClickEvent(this, handle__MouseDoubleClickEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseDoubleClickEvent(QMouseEvent* event) {

		QDesktopWidget::mouseDoubleClickEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseMoveEvent(QMouseEvent* event) override {
		if (handle__MouseMoveEvent == 0) {
			QDesktopWidget::mouseMoveEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_MouseMoveEvent(this, handle__MouseMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseMoveEvent(QMouseEvent* event) {

		QDesktopWidget::mouseMoveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WheelEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void wheelEvent(QWheelEvent* event) override {
		if (handle__WheelEvent == 0) {
			QDesktopWidget::wheelEvent(event);
			return;
		}
		
		QWheelEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_WheelEvent(this, handle__WheelEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_WheelEvent(QWheelEvent* event) {

		QDesktopWidget::wheelEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyPressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyPressEvent(QKeyEvent* event) override {
		if (handle__KeyPressEvent == 0) {
			QDesktopWidget::keyPressEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_KeyPressEvent(this, handle__KeyPressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyPressEvent(QKeyEvent* event) {

		QDesktopWidget::keyPressEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyReleaseEvent(QKeyEvent* event) override {
		if (handle__KeyReleaseEvent == 0) {
			QDesktopWidget::keyReleaseEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_KeyReleaseEvent(this, handle__KeyReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyReleaseEvent(QKeyEvent* event) {

		QDesktopWidget::keyReleaseEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusInEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusInEvent(QFocusEvent* event) override {
		if (handle__FocusInEvent == 0) {
			QDesktopWidget::focusInEvent(event);
			return;
		}
		
		QFocusEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_FocusInEvent(this, handle__FocusInEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusInEvent(QFocusEvent* event) {

		QDesktopWidget::focusInEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusOutEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusOutEvent(QFocusEvent* event) override {
		if (handle__FocusOutEvent == 0) {
			QDesktopWidget::focusOutEvent(event);
			return;
		}
		
		QFocusEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_FocusOutEvent(this, handle__FocusOutEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusOutEvent(QFocusEvent* event) {

		QDesktopWidget::focusOutEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void enterEvent(QEvent* event) override {
		if (handle__EnterEvent == 0) {
			QDesktopWidget::enterEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_EnterEvent(this, handle__EnterEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_EnterEvent(QEvent* event) {

		QDesktopWidget::enterEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__LeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void leaveEvent(QEvent* event) override {
		if (handle__LeaveEvent == 0) {
			QDesktopWidget::leaveEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_LeaveEvent(this, handle__LeaveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_LeaveEvent(QEvent* event) {

		QDesktopWidget::leaveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void paintEvent(QPaintEvent* event) override {
		if (handle__PaintEvent == 0) {
			QDesktopWidget::paintEvent(event);
			return;
		}
		
		QPaintEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_PaintEvent(this, handle__PaintEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_PaintEvent(QPaintEvent* event) {

		QDesktopWidget::paintEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void moveEvent(QMoveEvent* event) override {
		if (handle__MoveEvent == 0) {
			QDesktopWidget::moveEvent(event);
			return;
		}
		
		QMoveEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_MoveEvent(this, handle__MoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MoveEvent(QMoveEvent* event) {

		QDesktopWidget::moveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CloseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void closeEvent(QCloseEvent* event) override {
		if (handle__CloseEvent == 0) {
			QDesktopWidget::closeEvent(event);
			return;
		}
		
		QCloseEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_CloseEvent(this, handle__CloseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CloseEvent(QCloseEvent* event) {

		QDesktopWidget::closeEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ContextMenuEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void contextMenuEvent(QContextMenuEvent* event) override {
		if (handle__ContextMenuEvent == 0) {
			QDesktopWidget::contextMenuEvent(event);
			return;
		}
		
		QContextMenuEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_ContextMenuEvent(this, handle__ContextMenuEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ContextMenuEvent(QContextMenuEvent* event) {

		QDesktopWidget::contextMenuEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TabletEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void tabletEvent(QTabletEvent* event) override {
		if (handle__TabletEvent == 0) {
			QDesktopWidget::tabletEvent(event);
			return;
		}
		
		QTabletEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_TabletEvent(this, handle__TabletEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TabletEvent(QTabletEvent* event) {

		QDesktopWidget::tabletEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ActionEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void actionEvent(QActionEvent* event) override {
		if (handle__ActionEvent == 0) {
			QDesktopWidget::actionEvent(event);
			return;
		}
		
		QActionEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_ActionEvent(this, handle__ActionEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ActionEvent(QActionEvent* event) {

		QDesktopWidget::actionEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragEnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragEnterEvent(QDragEnterEvent* event) override {
		if (handle__DragEnterEvent == 0) {
			QDesktopWidget::dragEnterEvent(event);
			return;
		}
		
		QDragEnterEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_DragEnterEvent(this, handle__DragEnterEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragEnterEvent(QDragEnterEvent* event) {

		QDesktopWidget::dragEnterEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragMoveEvent(QDragMoveEvent* event) override {
		if (handle__DragMoveEvent == 0) {
			QDesktopWidget::dragMoveEvent(event);
			return;
		}
		
		QDragMoveEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_DragMoveEvent(this, handle__DragMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragMoveEvent(QDragMoveEvent* event) {

		QDesktopWidget::dragMoveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragLeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragLeaveEvent(QDragLeaveEvent* event) override {
		if (handle__DragLeaveEvent == 0) {
			QDesktopWidget::dragLeaveEvent(event);
			return;
		}
		
		QDragLeaveEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_DragLeaveEvent(this, handle__DragLeaveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragLeaveEvent(QDragLeaveEvent* event) {

		QDesktopWidget::dragLeaveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DropEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dropEvent(QDropEvent* event) override {
		if (handle__DropEvent == 0) {
			QDesktopWidget::dropEvent(event);
			return;
		}
		
		QDropEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_DropEvent(this, handle__DropEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DropEvent(QDropEvent* event) {

		QDesktopWidget::dropEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ShowEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void showEvent(QShowEvent* event) override {
		if (handle__ShowEvent == 0) {
			QDesktopWidget::showEvent(event);
			return;
		}
		
		QShowEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_ShowEvent(this, handle__ShowEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ShowEvent(QShowEvent* event) {

		QDesktopWidget::showEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HideEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void hideEvent(QHideEvent* event) override {
		if (handle__HideEvent == 0) {
			QDesktopWidget::hideEvent(event);
			return;
		}
		
		QHideEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_HideEvent(this, handle__HideEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_HideEvent(QHideEvent* event) {

		QDesktopWidget::hideEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__NativeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual bool nativeEvent(const QByteArray& eventType, void* message, long* result) override {
		if (handle__NativeEvent == 0) {
			return QDesktopWidget::nativeEvent(eventType, message, result);
		}
		
		const QByteArray eventType_qb = eventType;
		struct miqt_string eventType_ms;
		eventType_ms.len = eventType_qb.length();
		eventType_ms.data = static_cast<char*>(malloc(eventType_ms.len));
		memcpy(eventType_ms.data, eventType_qb.data(), eventType_ms.len);
		struct miqt_string sigval1 = eventType_ms;
		void* sigval2 = message;
		long* sigval3 = result;

		bool callback_return_value = miqt_exec_callback_QDesktopWidget_NativeEvent(this, handle__NativeEvent, sigval1, sigval2, sigval3);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_NativeEvent(struct miqt_string eventType, void* message, long* result) {
		QByteArray eventType_QByteArray(eventType.data, eventType.len);

		return QDesktopWidget::nativeEvent(eventType_QByteArray, message, static_cast<long*>(result));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChangeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void changeEvent(QEvent* param1) override {
		if (handle__ChangeEvent == 0) {
			QDesktopWidget::changeEvent(param1);
			return;
		}
		
		QEvent* sigval1 = param1;

		miqt_exec_callback_QDesktopWidget_ChangeEvent(this, handle__ChangeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChangeEvent(QEvent* param1) {

		QDesktopWidget::changeEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metric = 0;

	// Subclass to allow providing a Go implementation
	virtual int metric(QPaintDevice::PaintDeviceMetric param1) const override {
		if (handle__Metric == 0) {
			return QDesktopWidget::metric(param1);
		}
		
		QPaintDevice::PaintDeviceMetric param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		int callback_return_value = miqt_exec_callback_QDesktopWidget_Metric(this, handle__Metric, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Metric(int param1) const {

		return QDesktopWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InitPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual void initPainter(QPainter* painter) const override {
		if (handle__InitPainter == 0) {
			QDesktopWidget::initPainter(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QDesktopWidget_InitPainter(this, handle__InitPainter, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_InitPainter(QPainter* painter) const {

		QDesktopWidget::initPainter(painter);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Redirected = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintDevice* redirected(QPoint* offset) const override {
		if (handle__Redirected == 0) {
			return QDesktopWidget::redirected(offset);
		}
		
		QPoint* sigval1 = offset;

		QPaintDevice* callback_return_value = miqt_exec_callback_QDesktopWidget_Redirected(this, handle__Redirected, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintDevice* virtualbase_Redirected(QPoint* offset) const {

		return QDesktopWidget::redirected(offset);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SharedPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual QPainter* sharedPainter() const override {
		if (handle__SharedPainter == 0) {
			return QDesktopWidget::sharedPainter();
		}
		

		QPainter* callback_return_value = miqt_exec_callback_QDesktopWidget_SharedPainter(this, handle__SharedPainter);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPainter* virtualbase_SharedPainter() const {

		return QDesktopWidget::sharedPainter();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InputMethodEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void inputMethodEvent(QInputMethodEvent* param1) override {
		if (handle__InputMethodEvent == 0) {
			QDesktopWidget::inputMethodEvent(param1);
			return;
		}
		
		QInputMethodEvent* sigval1 = param1;

		miqt_exec_callback_QDesktopWidget_InputMethodEvent(this, handle__InputMethodEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_InputMethodEvent(QInputMethodEvent* param1) {

		QDesktopWidget::inputMethodEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InputMethodQuery = 0;

	// Subclass to allow providing a Go implementation
	virtual QVariant inputMethodQuery(Qt::InputMethodQuery param1) const override {
		if (handle__InputMethodQuery == 0) {
			return QDesktopWidget::inputMethodQuery(param1);
		}
		
		Qt::InputMethodQuery param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		QVariant* callback_return_value = miqt_exec_callback_QDesktopWidget_InputMethodQuery(this, handle__InputMethodQuery, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QVariant* virtualbase_InputMethodQuery(int param1) const {

		return new QVariant(QDesktopWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(param1)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusNextPrevChild = 0;

	// Subclass to allow providing a Go implementation
	virtual bool focusNextPrevChild(bool next) override {
		if (handle__FocusNextPrevChild == 0) {
			return QDesktopWidget::focusNextPrevChild(next);
		}
		
		bool sigval1 = next;

		bool callback_return_value = miqt_exec_callback_QDesktopWidget_FocusNextPrevChild(this, handle__FocusNextPrevChild, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_FocusNextPrevChild(bool next) {

		return QDesktopWidget::focusNextPrevChild(next);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__EventFilter == 0) {
			return QDesktopWidget::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QDesktopWidget_EventFilter(this, handle__EventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EventFilter(QObject* watched, QEvent* event) {

		return QDesktopWidget::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__TimerEvent == 0) {
			QDesktopWidget::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* event) {

		QDesktopWidget::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChildEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__ChildEvent == 0) {
			QDesktopWidget::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_ChildEvent(this, handle__ChildEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChildEvent(QChildEvent* event) {

		QDesktopWidget::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CustomEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__CustomEvent == 0) {
			QDesktopWidget::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QDesktopWidget_CustomEvent(this, handle__CustomEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CustomEvent(QEvent* event) {

		QDesktopWidget::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__ConnectNotify == 0) {
			QDesktopWidget::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QDesktopWidget_ConnectNotify(this, handle__ConnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ConnectNotify(QMetaMethod* signal) {

		QDesktopWidget::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DisconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__DisconnectNotify == 0) {
			QDesktopWidget::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QDesktopWidget_DisconnectNotify(this, handle__DisconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DisconnectNotify(QMetaMethod* signal) {

		QDesktopWidget::disconnectNotify(*signal);

	}

};

QDesktopWidget* QDesktopWidget_new() {
	return new MiqtVirtualQDesktopWidget();
}

void QDesktopWidget_virtbase(QDesktopWidget* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QDesktopWidget_MetaObject(const QDesktopWidget* self) {
	return (QMetaObject*) self->metaObject();
}

void* QDesktopWidget_Metacast(QDesktopWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QDesktopWidget_Tr(const char* s) {
	QString _ret = QDesktopWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDesktopWidget_TrUtf8(const char* s) {
	QString _ret = QDesktopWidget::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QDesktopWidget_ScreenNumber(const QDesktopWidget* self) {
	return self->screenNumber();
}

QRect* QDesktopWidget_ScreenGeometry(const QDesktopWidget* self, QWidget* widget) {
	return new QRect(self->screenGeometry(widget));
}

QRect* QDesktopWidget_AvailableGeometry(const QDesktopWidget* self, QWidget* widget) {
	return new QRect(self->availableGeometry(widget));
}

bool QDesktopWidget_IsVirtualDesktop(const QDesktopWidget* self) {
	return self->isVirtualDesktop();
}

int QDesktopWidget_NumScreens(const QDesktopWidget* self) {
	return self->numScreens();
}

int QDesktopWidget_ScreenCount(const QDesktopWidget* self) {
	return self->screenCount();
}

int QDesktopWidget_PrimaryScreen(const QDesktopWidget* self) {
	return self->primaryScreen();
}

int QDesktopWidget_ScreenNumberWithQPoint(const QDesktopWidget* self, QPoint* param1) {
	return self->screenNumber(*param1);
}

QWidget* QDesktopWidget_Screen(QDesktopWidget* self) {
	return self->screen();
}

QRect* QDesktopWidget_ScreenGeometry2(const QDesktopWidget* self) {
	return new QRect(self->screenGeometry());
}

QRect* QDesktopWidget_ScreenGeometryWithPoint(const QDesktopWidget* self, QPoint* point) {
	return new QRect(self->screenGeometry(*point));
}

QRect* QDesktopWidget_AvailableGeometry2(const QDesktopWidget* self) {
	return new QRect(self->availableGeometry());
}

QRect* QDesktopWidget_AvailableGeometryWithPoint(const QDesktopWidget* self, QPoint* point) {
	return new QRect(self->availableGeometry(*point));
}

void QDesktopWidget_Resized(QDesktopWidget* self, int param1) {
	self->resized(static_cast<int>(param1));
}

void QDesktopWidget_connect_Resized(QDesktopWidget* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget::connect(self, static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::resized), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QDesktopWidget_Resized(slot, sigval1);
	});
}

void QDesktopWidget_WorkAreaResized(QDesktopWidget* self, int param1) {
	self->workAreaResized(static_cast<int>(param1));
}

void QDesktopWidget_connect_WorkAreaResized(QDesktopWidget* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget::connect(self, static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::workAreaResized), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QDesktopWidget_WorkAreaResized(slot, sigval1);
	});
}

void QDesktopWidget_ScreenCountChanged(QDesktopWidget* self, int param1) {
	self->screenCountChanged(static_cast<int>(param1));
}

void QDesktopWidget_connect_ScreenCountChanged(QDesktopWidget* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget::connect(self, static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::screenCountChanged), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QDesktopWidget_ScreenCountChanged(slot, sigval1);
	});
}

void QDesktopWidget_PrimaryScreenChanged(QDesktopWidget* self) {
	self->primaryScreenChanged();
}

void QDesktopWidget_connect_PrimaryScreenChanged(QDesktopWidget* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget::connect(self, static_cast<void (QDesktopWidget::*)()>(&QDesktopWidget::primaryScreenChanged), self, [=]() {
		miqt_exec_callback_QDesktopWidget_PrimaryScreenChanged(slot);
	});
}

struct miqt_string QDesktopWidget_Tr2(const char* s, const char* c) {
	QString _ret = QDesktopWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDesktopWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDesktopWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDesktopWidget_TrUtf82(const char* s, const char* c) {
	QString _ret = QDesktopWidget::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDesktopWidget_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QDesktopWidget::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QDesktopWidget_ScreenNumber1(const QDesktopWidget* self, QWidget* widget) {
	return self->screenNumber(widget);
}

QWidget* QDesktopWidget_Screen1(QDesktopWidget* self, int screen) {
	return self->screen(static_cast<int>(screen));
}

QRect* QDesktopWidget_ScreenGeometry1(const QDesktopWidget* self, int screen) {
	return new QRect(self->screenGeometry(static_cast<int>(screen)));
}

QRect* QDesktopWidget_AvailableGeometry1(const QDesktopWidget* self, int screen) {
	return new QRect(self->availableGeometry(static_cast<int>(screen)));
}

bool QDesktopWidget_override_virtual_ResizeEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ResizeEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_ResizeEvent(void* self, QResizeEvent* e) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_ResizeEvent(e);
}

bool QDesktopWidget_override_virtual_DevType(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DevType = slot;
	return true;
}

int QDesktopWidget_virtualbase_DevType(const void* self) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_DevType();
}

bool QDesktopWidget_override_virtual_SetVisible(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__SetVisible = slot;
	return true;
}

void QDesktopWidget_virtualbase_SetVisible(void* self, bool visible) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_SetVisible(visible);
}

bool QDesktopWidget_override_virtual_SizeHint(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__SizeHint = slot;
	return true;
}

QSize* QDesktopWidget_virtualbase_SizeHint(const void* self) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_SizeHint();
}

bool QDesktopWidget_override_virtual_MinimumSizeHint(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MinimumSizeHint = slot;
	return true;
}

QSize* QDesktopWidget_virtualbase_MinimumSizeHint(const void* self) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_MinimumSizeHint();
}

bool QDesktopWidget_override_virtual_HeightForWidth(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__HeightForWidth = slot;
	return true;
}

int QDesktopWidget_virtualbase_HeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_HeightForWidth(param1);
}

bool QDesktopWidget_override_virtual_HasHeightForWidth(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__HasHeightForWidth = slot;
	return true;
}

bool QDesktopWidget_virtualbase_HasHeightForWidth(const void* self) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_HasHeightForWidth();
}

bool QDesktopWidget_override_virtual_PaintEngine(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__PaintEngine = slot;
	return true;
}

QPaintEngine* QDesktopWidget_virtualbase_PaintEngine(const void* self) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_PaintEngine();
}

bool QDesktopWidget_override_virtual_Event(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Event = slot;
	return true;
}

bool QDesktopWidget_virtualbase_Event(void* self, QEvent* event) {
	return ( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_Event(event);
}

bool QDesktopWidget_override_virtual_MousePressEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MousePressEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_MousePressEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_MousePressEvent(event);
}

bool QDesktopWidget_override_virtual_MouseReleaseEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MouseReleaseEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_MouseReleaseEvent(event);
}

bool QDesktopWidget_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MouseDoubleClickEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_MouseDoubleClickEvent(event);
}

bool QDesktopWidget_override_virtual_MouseMoveEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MouseMoveEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_MouseMoveEvent(event);
}

bool QDesktopWidget_override_virtual_WheelEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__WheelEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_WheelEvent(void* self, QWheelEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_WheelEvent(event);
}

bool QDesktopWidget_override_virtual_KeyPressEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__KeyPressEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_KeyPressEvent(void* self, QKeyEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_KeyPressEvent(event);
}

bool QDesktopWidget_override_virtual_KeyReleaseEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__KeyReleaseEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_KeyReleaseEvent(event);
}

bool QDesktopWidget_override_virtual_FocusInEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__FocusInEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_FocusInEvent(void* self, QFocusEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_FocusInEvent(event);
}

bool QDesktopWidget_override_virtual_FocusOutEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__FocusOutEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_FocusOutEvent(void* self, QFocusEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_FocusOutEvent(event);
}

bool QDesktopWidget_override_virtual_EnterEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__EnterEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_EnterEvent(void* self, QEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_EnterEvent(event);
}

bool QDesktopWidget_override_virtual_LeaveEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__LeaveEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_LeaveEvent(void* self, QEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_LeaveEvent(event);
}

bool QDesktopWidget_override_virtual_PaintEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__PaintEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_PaintEvent(void* self, QPaintEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_PaintEvent(event);
}

bool QDesktopWidget_override_virtual_MoveEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__MoveEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_MoveEvent(void* self, QMoveEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_MoveEvent(event);
}

bool QDesktopWidget_override_virtual_CloseEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__CloseEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_CloseEvent(void* self, QCloseEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_CloseEvent(event);
}

bool QDesktopWidget_override_virtual_ContextMenuEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ContextMenuEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_ContextMenuEvent(event);
}

bool QDesktopWidget_override_virtual_TabletEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__TabletEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_TabletEvent(void* self, QTabletEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_TabletEvent(event);
}

bool QDesktopWidget_override_virtual_ActionEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ActionEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_ActionEvent(void* self, QActionEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_ActionEvent(event);
}

bool QDesktopWidget_override_virtual_DragEnterEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DragEnterEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_DragEnterEvent(event);
}

bool QDesktopWidget_override_virtual_DragMoveEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DragMoveEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_DragMoveEvent(event);
}

bool QDesktopWidget_override_virtual_DragLeaveEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DragLeaveEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_DragLeaveEvent(event);
}

bool QDesktopWidget_override_virtual_DropEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DropEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_DropEvent(void* self, QDropEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_DropEvent(event);
}

bool QDesktopWidget_override_virtual_ShowEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ShowEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_ShowEvent(void* self, QShowEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_ShowEvent(event);
}

bool QDesktopWidget_override_virtual_HideEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__HideEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_HideEvent(void* self, QHideEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_HideEvent(event);
}

bool QDesktopWidget_override_virtual_NativeEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__NativeEvent = slot;
	return true;
}

bool QDesktopWidget_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, long* result) {
	return ( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_NativeEvent(eventType, message, result);
}

bool QDesktopWidget_override_virtual_ChangeEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ChangeEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_ChangeEvent(void* self, QEvent* param1) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_ChangeEvent(param1);
}

bool QDesktopWidget_override_virtual_Metric(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Metric = slot;
	return true;
}

int QDesktopWidget_virtualbase_Metric(const void* self, int param1) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_Metric(param1);
}

bool QDesktopWidget_override_virtual_InitPainter(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__InitPainter = slot;
	return true;
}

void QDesktopWidget_virtualbase_InitPainter(const void* self, QPainter* painter) {
	( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_InitPainter(painter);
}

bool QDesktopWidget_override_virtual_Redirected(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__Redirected = slot;
	return true;
}

QPaintDevice* QDesktopWidget_virtualbase_Redirected(const void* self, QPoint* offset) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_Redirected(offset);
}

bool QDesktopWidget_override_virtual_SharedPainter(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__SharedPainter = slot;
	return true;
}

QPainter* QDesktopWidget_virtualbase_SharedPainter(const void* self) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_SharedPainter();
}

bool QDesktopWidget_override_virtual_InputMethodEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__InputMethodEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_InputMethodEvent(param1);
}

bool QDesktopWidget_override_virtual_InputMethodQuery(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__InputMethodQuery = slot;
	return true;
}

QVariant* QDesktopWidget_virtualbase_InputMethodQuery(const void* self, int param1) {
	return ( (const MiqtVirtualQDesktopWidget*)(self) )->virtualbase_InputMethodQuery(param1);
}

bool QDesktopWidget_override_virtual_FocusNextPrevChild(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__FocusNextPrevChild = slot;
	return true;
}

bool QDesktopWidget_virtualbase_FocusNextPrevChild(void* self, bool next) {
	return ( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_FocusNextPrevChild(next);
}

bool QDesktopWidget_override_virtual_EventFilter(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__EventFilter = slot;
	return true;
}

bool QDesktopWidget_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_EventFilter(watched, event);
}

bool QDesktopWidget_override_virtual_TimerEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__TimerEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_TimerEvent(event);
}

bool QDesktopWidget_override_virtual_ChildEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ChildEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_ChildEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_ChildEvent(event);
}

bool QDesktopWidget_override_virtual_CustomEvent(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__CustomEvent = slot;
	return true;
}

void QDesktopWidget_virtualbase_CustomEvent(void* self, QEvent* event) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_CustomEvent(event);
}

bool QDesktopWidget_override_virtual_ConnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__ConnectNotify = slot;
	return true;
}

void QDesktopWidget_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_ConnectNotify(signal);
}

bool QDesktopWidget_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQDesktopWidget* self_cast = dynamic_cast<MiqtVirtualQDesktopWidget*>( (QDesktopWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__DisconnectNotify = slot;
	return true;
}

void QDesktopWidget_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQDesktopWidget*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QDesktopWidget_Delete(QDesktopWidget* self) {
	delete self;
}

