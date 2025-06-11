push:
	git add .
	git commit -m "update"
	git tag v1.0.2
	git push origin v1.0.2
	git checkout master
	git merge v1.0.2
	git push origin master

.PHONY: push