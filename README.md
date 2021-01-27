# ficheSyntaxHighlight

[![DeepSource](https://deepsource.io/gh/pczajkowski/ficheSyntaxHighlight.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/pczajkowski/ficheSyntaxHighlight/?ref=repository-badge)

I'm using the great [solusipse/fiche](https://github.com/solusipse/fiche). It generates text files, but I want syntax highlighting and line numbers. So I've created this server.

It's using [alecthomas/chroma](https://github.com/alecthomas/chroma) for generating HTML files with syntax highlighting, but language detection is my own creation as chroma's one doesn't detect what I want.

By default you'll be served with HTML version of your paste, if you want TXT then simply add /t to the link, like `http://localhost/xxxx/t`.

fiche's author provides his own [solution](https://github.com/solusipse/fiche/tree/master/extras/lines). If you're not performance freak (it generates HTML with every request, not once like mine, plus it's Python) and are not afraid of Python and Flask on your server it could be better solution for you.

This project was created mainly for fun, so don't expect much;)
