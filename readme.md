# Static Site Generator

## deploy
change github folder to .github

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

하위 경로 (github.io 가 아닌 다른 레포에서 배포) 사용할 때 처리
페이지 생성 로직 수정