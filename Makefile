push:
	git add .
	git commit -m "update"
	git tag v1.0.0
	git push origin v1.0.0
	git checkout master
	git merge v1.0.0
	git push origin master

.PHONY: push