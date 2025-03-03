import std/os, seaqt/[qapplication, qcoreapplication, qmediacontent, qmediaplayer, qurl]

proc main() =
  let
    _ = QApplication.create()
    srcFile = expandFilename("pixabay-public-domain-strong-hit-36455.mp3")
    content = QMediaContent.create(QUrl.fromLocalFile(srcFile))
    player = QMediaPlayer.create()

  player.setMedia(content)
  player.setVolume(50)
  player.onStateChanged(
    proc(s: cint) =
      echo "- Playback state: ", s

      if s == QMediaPlayerStateEnum.StoppedState:
        echo "Playback complete."
        QCoreApplication.exit()
  )

  echo "Playback starting..."
  player.play()

  discard QApplication.exec()

when isMainModule:
  main()
