# Static Site Generator

## deploy

1. change github folder to .github
2. set config.toml url "https://<username>.github.io/"
3. if repo is not github.io, "https://<username>.github.io/<reponame>/"
4. push github repo
5. go and check
    setting -> action -> general -> Workflow permissions -> Read and write permissions
6. wait for build action
7. go to
    setting -> page -> Build and deployment -> branch -> select gh-pages
8. wait for deploy action

## rule

### file naming

Name the folder that will become group
- num.group_name
- ex) 1.Study
 Name of the md, html file that will be the page
- num.file_name
- ex) 1.CS.md

### page setting

The beginning of the PAGE file
```
+++
title = “title of page”
template = “template_name.html”
+++
```
If template is not present, default.html will be used as the default

## content dir

config menu tree as content dir tree

```
Folder Tree ->

content
|-- 1.about
|   `-- _index.md
|-- 2.study
|   |-- 1.language
|   |   `-- _index.md
|   |-- 2.cs
|   |   `-- _index.md
|   `-- _index.md
|-- 3.contact
|   `-- _index.md
`-- _index.md

Menu ->

Home About Study        Contact
           `-- Language
           `-- Cs
```

## static dir

public file

## template

go template html file

# TODO