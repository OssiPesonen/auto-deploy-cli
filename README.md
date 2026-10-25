# Auto deployment CLI tool

Third iteration of the same idea to build an auto deployment tool to help me install apps on on-premise servers without any manual steps. Instead of previous versions where my genius idea was to expose a public API to push to (with auth of course) I'm now building this with a pull model, where the app is triggered by systemd in set intervals to check if a release has happened in a repository before pulling the ready-for-use artifact onto the server.

I'm working in small pieces and using Neovim as my new IDE for this, so I'm learning that as I go.
