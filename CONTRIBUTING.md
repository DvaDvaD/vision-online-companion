# Contribute

## Contents

- [Contribute](#contribute)
  - [Contents](#contents)
  - [About This Project](#about-this-project)
  - [Working on Issue](#working-on-issue)
  - [Writting Commit Messages](#writting-commit-messages)
  - [Submitting Pull Request](#submitting-pull-request)
  - [Code Review](#code-review)


## About This Project

This project is private, and contributions are made within the company’s internal team. Below are the guidelines for contributing effectively to maintain the codebase and ensure smooth collaboration.

## Working on Issue

1. **Issue Assignment**: Issues are assigned by the lead or yourself. Once an issue is assigned to you, ensure you fully understand the requirements before starting. Hint: ask.

2. **Create a branch from main**: Use the following naming convention for branches:

```
<type>/<issue-number>-<short-description>
```

For example:

```
feat/1234-add-user-table
fix/5678-typo-in-profile
fix/5678-grid-in-home-scren
refactor/6021-api-user-payload
```

3. You may create several branches and PRs for a single issue if needed. E.g. When the implementation is bigger than expected in the initial discussion
4. Branch types should follow **conventional commit categories** such as:

    * `feat`: Introduce a new feature
    * `fix`: Fixing a bug
    * `refactor`: Changing the structure of the code **without altering its external behavior** (no breaking changes)
    * `chore`: Routine tasks or maintenance changes that **don't affect the logic or functionality of the application**
    * There are other categories, but we usually use these 4.

5. **Trunk-Based Development**: Our project uses a [trunk-based development](https://trunkbaseddevelopment.com/) strategy. Before creating a PR, ensure your branch is always rebased from the latest main. This helps avoid merge conflicts and keeps history clean. **Set Git to Pull with Rebase** before working on a project (only one time).

```
# Setup git pull to use rebase globally
$ git config --global pull.rebase true

# Go to the project folder
$ cd project

# In the main branch, you create a new branch
(main)
$ git pull origin main
(main)
$ git checkout -b feat/1234-add-user-table
Switched to a new branch 'feat/1234-add-user-table'

(feat/1234-add-user-table)
$ git add . && git commit -m "feat: create a migration for user table"

(feat/1234-add-user-table)
$ git pull origin main
$ git push origin feat/1234-add-user-table
```

Ps: More advanced usage is separating `git pull` to `git fetch origin` and `git rebase origin/main`, using `git diff` in between.  

6. You may need to solve a conflict after pull from main. By selecting the correct code, you should be fine. Please discuss with your peer that has last commit in main. Tricks to get less conflict:
    * Working on small changes or small number of files. Break the task to smaller action.
    * Make sure you pull before create a new branch
    * Make sure you are not stalling a branch for too long (several days)
    * Sometimes you better to stash your changes before pull, then apply it

## Writting Commit Messages

**Before you commit**: 

* Please check what you want to commit.
* Make sure every file that will be committed already saved and in the staging area.
* Make sure there is no hard coded credential in code or in a file that should be ignored.
* Well, you are welcome to commit often actually.

We are following [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/), please read it. Scope that we use:

* Type: `feat`, `fix`, `refactor` and `chore`. To simplify, other type like `ci`, `docs`, `test` will be in the `chore`.
* If there is implementation that not finished yet, please write a `TODO: your msg` in your code and tell it in the body.
* If there is a breaking changes, write it in the body what and why it change

> About Breaking Changes:<br /><br />
> When you write something that make your user to change their behaviour, then it's a Breaking Changes.<br />
> example 1: A backend change the API field from `user` to `customer`, it's a breaking since FE need to change immediately.<br />
> example 2: A backend add a field `customer` that have same content with `user` and mark `user` as deprecated, it's not breaking.<br />

For one liner commit, you can use:

```
git commit -m "feat: create a migration for user table"
```

For commit with description (e.g. a breaking changes), you can run `git commit` and it will open your code editor. If your default is vim and you are not familiar with it:

- Press `i` to enter `insert mode`.
- After you finish writing your message, press `esc` or `ctrl+c` to enter `normal mode`
- In normal mode, press `:wq` to write and quit or `:q!` to exit without saving.


## Submitting Pull Request

| Step                  | Description                                                                                                                           |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| Branch Naming         | Use descriptive branch names (e.g., feat/1234-add-login-feature, fix/5678-correct-typo-in-header).                                    |
| Small, Focused PRs    | Address a single feature or issue to keep PRs easy to review.                                                                         |
| Rebase from main      | Keep your branch up to date by rebasing before pushing.                                                                               |
| PR Title              | Write a concise title that gives context (e.g., feat: add login feature).                                                             |
| Detailed Description  | Explain What, Why, and How. Link related issues (e.g., Closes #1234). Make sure all commit messages also automatically included here. |
| Commit Messages       | Use descriptive, convention-based commit messages (e.g., feat: add login component).                                                  |
| Request Reviewers     | Tag appropriate reviewers and add relevant labels.                                                                                    |
| Screenshots/Demos     | For UI changes, include screenshots, GIFs, or demos to help reviewers visualize changes.                                              |
| Tests and Linting     | Ensure all tests pass and linting is clean before submitting.                                                                         |
| Documentation Updates | Update or add documentation as needed to reflect changes in functionality.                                                            |

## Code Review

1. **Prepare the PR for Review**

* Rebase from main: Ensure your branch is rebased with the latest main branch to prevent conflicts.
* Verify All Checks Pass: Make sure all tests and automated checks pass before requesting a review.

1. **Notify Reviewers Manually**

Don’t rely solely on GitHub notifications. Mention the peer reviewers in your team chat and share the PR title and URL for direct access.

1. **Collaborative Benefits of Code Review**

Skill Development:
* **Senior-to-Junior**: Senior developers can guide juniors, ensuring code is functional and aligns with standards.
* **Junior-to-Senior**: Junior reviewers can learn best practices and observe how complex features are built.
* **Peer-to-Peer**: Peers can engage in discussions that enhance skills and broaden their approach to problem-solving.

1. **PR Size and Readability**

* **Small, Focused PRs**: Aim to submit small, focused PRs for easier review. Large PRs that are difficult to read **should be rejected**, except in special cases.
* **Constructive Feedback Only**: Reviews should be objective, focusing on code improvements, not personal judgments.

1. **Handling Long Discussions**

Move Extended Discussions Elsewhere: If a discussion is likely to be lengthy, consider moving it to another forum (e.g., online meeting) and documenting a summary of the discussion on the PR page to provide context for future reference.

1. **Approvals and Merging**

* Reviewers indicate approval by labeling the PR as approved.
* The author is then responsible for ensuring a final rebase, confirming a squash and merge for clean commit history, and completing the merge (press the merge button and celebrate 🎉).