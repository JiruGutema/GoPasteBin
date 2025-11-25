## A small web service that lets you: paste text , get a link, view raw text and auto-expire

**Folder Structere:**
```

pastebin/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   │   └── paste.go
│   ├── db/
│   │   └── sqlite.go
│   ├── model/
│   │   └── paste.go
│   ├── services/
│   │   └── paste_service.go
│   ├── utils/
│   │   └── id.go
│   └── cleanup/
│       └── cleanup.go
│
├── templates/
│   ├── view.html
│   └── create.html
│
└── go.mod
```
