## hotpweb

This is a simple web app that demonstrates OATH-HOTP one-time
passwords. It uses the [hotp](https://github.com/gokyle/hotp) package
to provide the codes.

When you pull up the app for the first time, it will ask for your
name. Then, it will present you with a QR code that you should scan
in with the Google Authenticator app. Once it's loaded, you can
start trying to enter codes to see it in action.


### Running the app
```
go get github.com/robfig/revel
go get github.com/gokyle/hotpweb
revel run github.com/gokyle/hotpweb
```

By default, the app runs on 127.0.0.1:8080


### Notes

* I didn't want to require people to set up a database to use the
demo, so the key values are stored in the session. In practice,
this is an incredibly bad idea. The [hotp
godocs](https://godoc.org/github.com/gokyle/hotp/) have some notes
on storing the key values in practice.

* This only works with the Google Authenticator app. It wouldn't
be too hard to adapt it for YubiKeys programmed in OATH-HOTP mode,
but that was more work than I wanted to put into this. The `hotp`
package does have a method for
[handling YubiKeys](http://godoc.org/github.com/gokyle/hotp#HOTP.YubiKey),
though.


### License

Copyright (c) 2013 Kyle Isom <kyle@tyrfingr.is>

Permission to use, copy, modify, and distribute this software for any
purpose with or without fee is hereby granted, provided that the above 
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE. 
