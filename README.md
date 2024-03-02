# GoGitLabEnum 🕵️‍♂️

**GoGitLabEnum** is a powerful Go-based command-line tool designed for security enthusiasts and professionals to perform user enumeration on GitLab instances. By utilizing custom wordlists, it systematically checks for the existence of usernames, providing immediate feedback with an intuitive, color-coded output. This tool is essential for penetration testing and reconnaissance, offering a straightforward and effective method to identify potential targets within a GitLab environment.

## Features 🚀

- **Efficient Enumeration**: Leverages Go's concurrency for fast username checking.
- **Color-Coded Output**: Easily distinguish between found (green) and not found (red) usernames.
- **Custom Wordlists Support**: Tailor your search with any wordlist to fit the scope of your testing.
- **Simple CLI Usage**: Designed with simplicity in mind, just two flags are needed to start enumerating.

## Quick Start 🏁

Ensure you have Go installed on your system, then clone this repository and build the tool:

```bash
git clone https://github.com/yourusername/GoGitLabEnum.git
cd GoGitLabEnum
go build
```

To run the tool, simply provide the GitLab URL and the path to your wordlist:

```bash
./GoGitLabEnum --url https://gitlab.example.com --wordlist /path/to/wordlist.txt
```

## Contributing 🤝

Contributions are welcome! If you have improvements or bug fixes, please fork the repository, make your changes, and submit a pull request.

## License 📄

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

Let's make **GoGitLabEnum** the go-to tool for GitLab user enumeration! 🌟
