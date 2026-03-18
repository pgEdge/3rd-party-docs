#!/usr/bin/env bash
#-------------------------------------------------------------------------
#
# pgEdge PostgreSQL Docs
#
# Copyright (c) 2026, pgEdge, Inc.
# This software is released under The PostgreSQL License
#
#-------------------------------------------------------------------------
#
# build-all.sh — Build documentation for all (or selected) branches.
#
# Reads branches.yml, clones/fetches upstream repos, runs the
# converter in git worktrees, commits changes, and optionally
# pushes updated branches to the remote.
#
# Usage:
#   ./build-all.sh                          # build all branches
#   ./build-all.sh --branches pg17,pg18     # build specific branches
#   ./build-all.sh --branches "pg*"         # glob matching
#   ./build-all.sh --branches "pg*,post*"   # multiple patterns
#   ./build-all.sh --dry-run                # show what would be built
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CONFIG="${SCRIPT_DIR}/branches.yml"
WORKTREE_BASE="${SCRIPT_DIR}/.worktrees"
BINARY="${SCRIPT_DIR}/bin/pgdoc-converter"

BRANCH_FILTER=""
DRY_RUN=false

# Colours (disabled if not a terminal)
if [[ -t 1 ]]; then
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[0;33m'
    CYAN='\033[0;36m'
    BOLD='\033[1m'
    RESET='\033[0m'
else
    RED='' GREEN='' YELLOW='' CYAN='' BOLD='' RESET=''
fi

# ── Argument parsing ──────────────────────────────────────────────

usage() {
    cat <<EOF
Usage: $(basename "$0") [OPTIONS]

Build documentation for branches defined in branches.yml.

Options:
  --branches PATTERNS  Comma-separated branch names or glob patterns
                       (e.g. "pg17", "pg*", "pg17,postgrest145")
  --config PATH        Path to config file (default: branches.yml)
  --dry-run            Show what would be built without doing it
  -h, --help           Show this help message
EOF
}

while [[ $# -gt 0 ]]; do
    case "$1" in
        --branches)
            BRANCH_FILTER="$2"
            shift 2
            ;;
        --config)
            CONFIG="$2"
            shift 2
            ;;
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        -h|--help)
            usage
            exit 0
            ;;
        *)
            echo "Unknown option: $1" >&2
            usage >&2
            exit 1
            ;;
    esac
done

# ── Dependency checks ─────────────────────────────────────────────

check_deps() {
    local missing=()

    if ! command -v yq &>/dev/null; then
        missing+=("yq (https://github.com/mikefarah/yq)")
    fi
    if ! command -v git &>/dev/null; then
        missing+=("git")
    fi
    if ! command -v go &>/dev/null; then
        missing+=("go")
    fi

    if [[ ${#missing[@]} -gt 0 ]]; then
        echo -e "${RED}Missing required tools:${RESET}" >&2
        for dep in "${missing[@]}"; do
            echo "  - $dep" >&2
        done
        exit 1
    fi

    if [[ ! -f "$CONFIG" ]]; then
        echo -e "${RED}Config not found: ${CONFIG}${RESET}" >&2
        exit 1
    fi
}

# ── Filter matching ───────────────────────────────────────────────

matches_filter() {
    local branch="$1"

    # No filter means match everything
    if [[ -z "$BRANCH_FILTER" ]]; then
        return 0
    fi

    IFS=',' read -ra patterns <<< "$BRANCH_FILTER"
    for pattern in "${patterns[@]}"; do
        pattern="$(echo "$pattern" | xargs)" # trim whitespace
        # shellcheck disable=SC2053
        if [[ "$branch" == $pattern ]]; then
            return 0
        fi
    done
    return 1
}

# ── Upstream repo management ──────────────────────────────────────

ensure_upstream() {
    local url="$1"
    local ref="$2"
    local cache_dir="$3"

    local repo_name
    repo_name="$(basename "$url" .git)"
    local repo_path="${cache_dir}/${repo_name}"

    if [[ ! -d "$repo_path" ]]; then
        echo -e "  ${CYAN}Cloning ${repo_name}...${RESET}" >&2
        git clone --quiet "$url" "$repo_path"
    else
        echo -e "  ${CYAN}Fetching ${repo_name}...${RESET}" >&2
        git -C "$repo_path" fetch --quiet origin --prune --tags
    fi

    # Checkout the requested ref (tag, branch, or commit)
    if ! git -C "$repo_path" checkout --quiet --detach "$ref" 2>/dev/null; then
        git -C "$repo_path" checkout --quiet --detach "origin/$ref" 2>/dev/null
    fi

    echo "$repo_path"
}

# ── Branch worktree management ────────────────────────────────────

setup_worktree() {
    local branch="$1"
    local worktree="${WORKTREE_BASE}/${branch}"

    # Clean up any stale worktree
    if [[ -d "$worktree" ]]; then
        git -C "$SCRIPT_DIR" worktree remove --force \
            "$worktree" 2>/dev/null || rm -rf "$worktree"
    fi

    # Branch exists locally
    if git -C "$SCRIPT_DIR" show-ref --verify --quiet \
        "refs/heads/${branch}" 2>/dev/null; then
        git -C "$SCRIPT_DIR" worktree add --quiet \
            "$worktree" "$branch"

    # Branch exists on remote only
    elif git -C "$SCRIPT_DIR" show-ref --verify --quiet \
        "refs/remotes/origin/${branch}" 2>/dev/null; then
        git -C "$SCRIPT_DIR" worktree add --quiet \
            "$worktree" -b "$branch" "origin/${branch}"

    # New branch — create from main
    else
        echo -e "  ${YELLOW}Creating new branch ${branch} " \
            "from main${RESET}" >&2
        git -C "$SCRIPT_DIR" worktree add --quiet \
            "$worktree" -b "$branch" main
    fi

    echo "$worktree"
}

remove_worktree() {
    local branch="$1"
    local worktree="${WORKTREE_BASE}/${branch}"
    git -C "$SCRIPT_DIR" worktree remove --force \
        "$worktree" 2>/dev/null || rm -rf "$worktree"
}

# ── Clean generated docs ─────────────────────────────────────────

clean_docs() {
    local worktree="$1"
    local docs_dir="${worktree}/docs"

    if [[ ! -d "$docs_dir" ]]; then
        return
    fi

    # Remove everything except the MkDocs support directories
    find "$docs_dir" -mindepth 1 -maxdepth 1 \
        ! -name 'img' \
        ! -name 'stylesheets' \
        ! -name 'overrides' \
        -exec rm -rf {} +
}

# ── Run the converter ─────────────────────────────────────────────

run_converter() {
    local worktree="$1"
    local mode="$2"
    local src_path="$3"
    local version="$4"
    local copyright="$5"
    local pgadmin_src="$6"
    local skip_sections="$7"

    local cmd=("$BINARY"
        -mode "$mode"
        -src "$src_path"
        -out "${worktree}/docs"
        -mkdocs "${worktree}/mkdocs.yml"
        -version "$version"
    )

    if [[ "$mode" == "rst" ]]; then
        [[ -n "$copyright" ]] && \
            cmd+=(-copyright "$copyright")
        [[ -n "$pgadmin_src" ]] && \
            cmd+=(-pgadmin-src "$pgadmin_src")
        [[ -n "$skip_sections" ]] && \
            cmd+=(-skip-sections "$skip_sections")
    fi

    "${cmd[@]}"
}

# ── Commit changes ────────────────────────────────────────────────

commit_if_changed() {
    local worktree="$1"
    local branch="$2"
    local version="$3"

    cd "$worktree"
    git add -A

    if git diff --cached --quiet; then
        return 1 # no changes
    fi

    # Read site_name from mkdocs.yml for the commit message
    local site_name
    site_name="$(yq '.site_name' "${worktree}/mkdocs.yml" \
        2>/dev/null || echo "$branch $version")"

    git commit --quiet \
        -m "Regenerate ${site_name} documentation"
    return 0
}

# ── Cleanup trap ──────────────────────────────────────────────────

cleanup() {
    if [[ -d "$WORKTREE_BASE" ]]; then
        git -C "$SCRIPT_DIR" worktree prune 2>/dev/null || true
        rm -rf "$WORKTREE_BASE"
    fi
    cd "$SCRIPT_DIR"
}

trap cleanup EXIT

# ── Main ──────────────────────────────────────────────────────────

main() {
    check_deps

    # Ensure we start from the repo root on main
    cd "$SCRIPT_DIR"
    local current_branch
    current_branch="$(git rev-parse --abbrev-ref HEAD)"
    if [[ "$current_branch" != "main" ]]; then
        echo -e "${RED}Must be run from the main branch" \
            "(currently on ${current_branch})${RESET}" >&2
        exit 1
    fi

    # Read config
    local cache_dir
    cache_dir="$(yq -r '.cache_dir' "$CONFIG")"
    if [[ "$cache_dir" != /* ]]; then
        cache_dir="${SCRIPT_DIR}/${cache_dir}"
    fi
    mkdir -p "$cache_dir"

    local branch_count
    branch_count="$(yq '.branches | length' "$CONFIG")"

    # Collect matching branches for display
    local selected=()
    for ((i = 0; i < branch_count; i++)); do
        local b
        b="$(yq -r ".branches[$i].branch" "$CONFIG")"
        if matches_filter "$b"; then
            selected+=("$b")
        fi
    done

    if [[ ${#selected[@]} -eq 0 ]]; then
        echo -e "${YELLOW}No branches matched filter:" \
            "'${BRANCH_FILTER}'${RESET}"
        exit 0
    fi

    echo -e "${BOLD}Building ${#selected[@]} branch(es):" \
        "${selected[*]}${RESET}"
    echo ""

    # Build converter binary
    echo -e "${BOLD}Building converter...${RESET}"
    make -C "$SCRIPT_DIR" build --quiet
    echo ""

    if $DRY_RUN; then
        echo -e "${YELLOW}Dry run — listing branches " \
            "that would be built:${RESET}"
        for ((i = 0; i < branch_count; i++)); do
            local b mode version upstream ref src_subdir
            b="$(yq -r ".branches[$i].branch" "$CONFIG")"
            matches_filter "$b" || continue
            mode="$(yq -r ".branches[$i].mode" "$CONFIG")"
            version="$(yq -r ".branches[$i].version" "$CONFIG")"
            upstream="$(yq -r ".branches[$i].upstream" "$CONFIG")"
            ref="$(yq -r ".branches[$i].ref" "$CONFIG")"
            src_subdir="$(yq -r ".branches[$i].src_subdir" "$CONFIG")"
            printf "  %-20s %-5s %-10s %s@%s:%s\n" \
                "$b" "$mode" "$version" \
                "$(basename "$upstream" .git)" "$ref" "$src_subdir"
        done
        exit 0
    fi

    # Create worktree base
    mkdir -p "$WORKTREE_BASE"

    # Track results
    declare -A RESULTS
    local updated_branches=()

    # Process each branch
    for ((i = 0; i < branch_count; i++)); do
        local branch mode version upstream ref src_subdir
        local copyright pgadmin_src_flag skip_sections

        branch="$(yq -r ".branches[$i].branch" "$CONFIG")"
        matches_filter "$branch" || continue

        mode="$(yq -r ".branches[$i].mode" "$CONFIG")"
        version="$(yq -r ".branches[$i].version" "$CONFIG")"
        upstream="$(yq -r ".branches[$i].upstream" "$CONFIG")"
        ref="$(yq -r ".branches[$i].ref" "$CONFIG")"
        src_subdir="$(yq -r ".branches[$i].src_subdir" "$CONFIG")"
        copyright="$(yq -r \
            ".branches[$i].copyright // \"\"" "$CONFIG")"
        skip_sections="$(yq -r \
            ".branches[$i].skip_sections // \"\"" "$CONFIG")"

        local use_pgadmin_src
        use_pgadmin_src="$(yq -r \
            ".branches[$i].pgadmin_src // \"\"" "$CONFIG")"

        echo -e "${BOLD}=== ${branch} " \
            "(${mode}, ${version}) ===${RESET}"

        # Fetch upstream source
        local repo_path
        if ! repo_path="$(ensure_upstream \
                "$upstream" "$ref" "$cache_dir")"; then
            echo -e "  ${RED}Failed to fetch upstream${RESET}"
            RESULTS["$branch"]="FAILED (fetch)"
            continue
        fi
        local src_path="${repo_path}/${src_subdir}"

        # Resolve pgadmin_src if needed
        pgadmin_src_flag=""
        if [[ "$use_pgadmin_src" == "yes" ]]; then
            pgadmin_src_flag="$repo_path"
        fi

        # Setup branch worktree
        local worktree
        if ! worktree="$(setup_worktree "$branch")"; then
            echo -e "  ${RED}Failed to create worktree${RESET}"
            RESULTS["$branch"]="FAILED (worktree)"
            continue
        fi

        # Reset mkdocs.yml to skeleton from main
        git -C "$SCRIPT_DIR" show main:mkdocs.yml \
            > "${worktree}/mkdocs.yml"

        # Clean generated docs
        clean_docs "$worktree"

        # Run converter
        echo "  Converting..."
        if ! run_converter "$worktree" "$mode" "$src_path" \
                "$version" "$copyright" \
                "$pgadmin_src_flag" "$skip_sections"; then
            echo -e "  ${RED}Conversion failed${RESET}"
            RESULTS["$branch"]="FAILED (convert)"
            remove_worktree "$branch"
            cd "$SCRIPT_DIR"
            continue
        fi

        # Commit
        if commit_if_changed "$worktree" "$branch" "$version"
        then
            RESULTS["$branch"]="updated"
            updated_branches+=("$branch")
            echo -e "  ${GREEN}Committed${RESET}"
        else
            RESULTS["$branch"]="no changes"
            echo -e "  ${CYAN}No changes${RESET}"
        fi

        cd "$SCRIPT_DIR"
        remove_worktree "$branch"
        echo ""
    done

    cd "$SCRIPT_DIR"

    # ── Summary ───────────────────────────────────────────────────

    echo -e "${BOLD}===============================${RESET}"
    echo -e "${BOLD}Build Summary${RESET}"
    echo -e "${BOLD}===============================${RESET}"
    printf "  %-22s %s\n" "Branch" "Status"
    printf "  %-22s %s\n" "------" "------"
    for branch in "${selected[@]}"; do
        local status="${RESULTS[$branch]:-skipped}"
        local colour="$RESET"
        case "$status" in
            updated)     colour="$GREEN" ;;
            "no changes") colour="$CYAN" ;;
            FAILED*)     colour="$RED" ;;
        esac
        printf "  %-22s ${colour}%s${RESET}\n" \
            "$branch" "$status"
    done
    echo -e "${BOLD}===============================${RESET}"

    # ── Push prompt ───────────────────────────────────────────────

    if [[ ${#updated_branches[@]} -eq 0 ]]; then
        echo ""
        echo "No branches were updated."
        return
    fi

    echo ""
    echo "The following branches have new commits:"
    for b in "${updated_branches[@]}"; do
        # Show the commit that was made
        local msg
        msg="$(git log -1 --oneline "$b" 2>/dev/null)"
        echo "  $msg"
    done
    echo ""
    if [[ ! -t 0 ]]; then
        echo "Non-interactive mode — skipping push. Run:"
        echo "  git push origin ${updated_branches[*]}"
        return
    fi
    read -rp "Push all updated branches to origin? [y/N] " confirm
    if [[ "$confirm" =~ ^[Yy]$ ]]; then
        for b in "${updated_branches[@]}"; do
            echo -e "  Pushing ${CYAN}${b}${RESET}..."
            git push origin "$b"
        done
        echo -e "${GREEN}Done.${RESET}"
    else
        echo "Skipped. Push manually with:"
        echo "  git push origin ${updated_branches[*]}"
    fi
}

main "$@"
