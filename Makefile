push:
	git add .
	git commit -m "update"
	git tag v1.0.1
	git push origin v1.0.1
	git checkout master
	git merge v1.0.1
	git push origin master

.PHONY: push