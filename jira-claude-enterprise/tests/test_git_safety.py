from app.git_service import GitService

def test_safe_branch_regex(tmp_path):
    git = GitService(tmp_path, "main")
    assert git.SAFE_BRANCH.fullmatch("feature/PROJ-123-add-payment")
    assert not git.SAFE_BRANCH.fullmatch("feature/../main")
