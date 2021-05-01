#!/usr/bin/env zsh

cat <<'EOF' >|tree.md
# gorepotemplate

> This is the initial directory tree for gorepotemplate. Use the make_tree_md.sh script (GNU-tree required) to update it if you wish. It is safe to delete this file.

### Directory Structure

```sh
EOF

tree -a -I '.git|bak' >> tree.md
echo '```' >> tree.md
