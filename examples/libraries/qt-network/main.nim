import
  seaqt/
    [qapplication, qcoreapplication, qdnslookup, qdnshostaddressrecord, qhostaddress]

proc main() =
  let _ = QApplication.create()

  echo "Looking up DNS info, please wait..."

  let dns = QDnsLookup.create(QDnsLookupTypeEnum.A, "google.com")

  dns.onFinished(
    proc() =
      if dns.error() != QDnsLookupErrorEnum.NoError:
        echo "An error occurred: ", dns.errorString()
        return

      let results = dns.hostAddressRecords()
        # CanonicalNameRecords, TextRecords, ServiceRecords, ...
      echo "Found ", len(results), " result(s)."

      for record in results:
        echo "- ", record.value().toString()

      QCoreApplication.exit()
  )
  dns.lookup()

  discard QApplication.exec()

when isMainModule:
  main()
