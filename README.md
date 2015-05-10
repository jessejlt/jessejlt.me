About
---

Implementation for [jessejlt.me](http://jessejlt.me)

Hack
---

* To update resume, modify resume.json and then render the template via `go run template.go`
* Build docker image `docker build -t nginx .`
* Run docker image `docker run -p 80:80 -v /<absolute path>/site:/var/www/home -d nginx` then visit `boot2docker ip`