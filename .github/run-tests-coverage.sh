#!/bin/bash
# =============================================================================
#  Test Script to Run Go Test and Check 100% Coverage
# =============================================================================
#  This script WILL FAIL if the coverage rate is NOT "100%". To show where to
#  fix/cover then specify '-v' or '--verbose' option.
#
#  Requirements:
#    - go-carpet: https://github.com/msoap/go-carpet

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
PATH_DIR_PARENT="$(dirname "$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)")"
NAME_FILE_COVERAGE='coverage.out'
SUCCESS=0
FAILURE=1
TRUE=0
FALSE=1

# -----------------------------------------------------------------------------
#  Functions
# -----------------------------------------------------------------------------

# getNameFilesUncovered echoes file names in comma-separated list which didn't
# cover 100%. This will be used for “go-carpet` -file option.
function getNameFilesUncovered() {
    # Ommit lines that includes "100.0%" and "total" then awk the fist field
    list=$(go tool cover -func="$NAME_FILE_COVERAGE" | grep -v '100.0%\|total' | awk '{print $1}')

    # Trim the line number and get only the file path and join them with a comma
    echo "$list" | while IFS= read -r line; do
        basename "${line}" | sed -e 's/\.[^\.]*$/\.go/'
    done | uniq | tr '\n' ','
}

# indentStdIn indents the STDIN given to the function
function indentStdIn() {
    indent="\t"
    while IFS= read -r line; do
        echo -e "${indent}${line}"
    done
    echo
}

# isModeVerbose just returns a bool whether it's in verbose
# mode or not.
function isModeVerbose() {
    if [ "$mode_verbose" -eq 0 ]; then
        return $TRUE
    fi

    return $FALSE
}

# runGoCarpet displays details of the coverage
function runGoCarpet() {
    if ! which go-carpet >/dev/null; then
        echo '* aborted'
        echo >&2 '  * Command "go-carpet" not found.'
        echo >&2 '    "go-carpet" is needed to view the test coverage area in the terminal.'
        echo >&2 '    To install see: https://github.com/msoap/go-carpet'

        exit $FAILURE
    fi

    list="$(getNameFilesUncovered)"

    go-carpet -file "$list"

    echo "Files to be covered: ${list}"
}

# runGoVet runs Go vet for static analysis.
function runGoVet() {
    description="${1:?'Test description missing.'}"
    path_dir="${2:?'Path is missing'}"

    echo
    echo "- Static analysys: ${description}"
    if isModeVerbose; then
        go vet -v "$path_dir" 2>&1 | indentStdIn
        result=$?
    else
        go vet "$path_dir" 2>&1 | indentStdIn
        result=$?
    fi

    if [ "$result" -ne 0 ]; then
        echo >&2 "  ERROR: Static analysis failed."
        exit $FAILURE
    fi
    echo '  Success! All Go vet static analysis passed.'
    return $SUCCESS
}

# runTests runs unit tests.
# If verbose option is provided then it will display the details. If the
# coverage was lower than 100% then it will fail and show the cover area
# as well.
function runTests() {
    description="${1:?'Test description missing.'}"
    path_dir="${2:?'Path is missing'}"

    echo
    echo "- Unit test: ${description}"
    # Run tests
    if isModeVerbose; then
        go test -timeout 30s -cover -v -coverprofile "$NAME_FILE_COVERAGE" "$path_dir" | indentStdIn
    else
        go test -timeout 30s -cover -coverprofile "$NAME_FILE_COVERAGE" "$path_dir" | indentStdIn
    fi

    # Get coverage details
    cover=$(go tool cover -func="$NAME_FILE_COVERAGE")

    if isModeVerbose; then
        echo '- Coverage details'
        echo "$cover" | indentStdIn
    fi

    # Get coverage rate
    coverage=$(echo "$cover" | grep total | awk '{print $3}')

    # go test -cover ./... | grep -v "100.0%" | awk '{print $2}'

    if [ "$coverage" = "100.0%" ]; then
        echo '  Success! Coverage: 100%'
        return $SUCCESS
    else
        # Displays where to cover, if the total coverage wasn't 100%
        if isModeVerbose; then
            echo '- Missing Cover Area'
            runGoCarpet | indentStdIn
            echo >&2 "  ERROR: Coverage failed. Did not cover 100% of the statements."
        else
            echo >&2 "  ERROR: Coverage failed. Did not cover 100% of the statements."
            echo >&2 "         Use '--verbose' option to see where to cover."
        fi
        echo >&2 "         Coverage: ${coverage}"

        exit $FAILURE
    fi
}

# -----------------------------------------------------------------------------
#  Setup
# -----------------------------------------------------------------------------
# Detect verbose option
mode_verbose=$FALSE
echo "${@}" | grep -e "-v" -e "--verbose" >/dev/null && {
    mode_verbose=$TRUE
}

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------
set -eu
set -o pipefail

if isModeVerbose; then
    echo '* Running in verbose mode.'
else
    echo '* Running in regular mode. Use "-v" or "--verbose" option for verbose output.'
fi
echo "* Moving current path to: ${PATH_DIR_PARENT}"
cd "$PATH_DIR_PARENT"
echo "* Current path is: $(pwd)"

runGoVet "Scanning all the packages" "./..."
runTests "Testing all the packages" "./..."
