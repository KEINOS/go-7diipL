#!/bin/bash

# Remove go-build* files in temp
rm -rf /tmp/go-build* 2>/dev/null 1>/dev/null

# Shoe welcome message
echo "========================================================================="
echo " Welcome To QiiTrans Development Container"
echo "========================================================================="
echo "- OS: $(head -n1 </etc/issue)"
echo "- $(go version)"
echo "- Current path: $(pwd)"
echo "- Current user: $(whoami)"
echo "- Time: $(date)"
echo "* Useful command and scripts:"
echo "  - To run unit and coverage tests : $ run-coverage"
echo "  - To run lint and static analysis: $ run-lint"
echo "  - To run full tests before merge : $ run-all"
echo "  - To update README.md documents  : $ update-docs"
echo "  - To build binary                : $ build-app --help"
echo "  - To display this message again  : $ welcome"
echo "  - To search a string in a file   : $ rg foobar ."
echo "  - To run the app before build    : $ go run ./main list"
echo "========================================================================="
