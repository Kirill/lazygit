// hmm auto-generated for testing purposes. To re-generate, do: <ifacemaker --file="pkg/commands/*.go" --struct=Git --iface=IGit --pkg=commands -o pkg/commands/igit.go --doc false --comment="$(cat pkg/commands/auto-generation-message.txt)"> from the root directory of the repo and fix up any missing imports

package commands

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"github.com/jesseduffield/lazygit/pkg/commands/patch"
	. "github.com/jesseduffield/lazygit/pkg/commands/types"
	"github.com/sirupsen/logrus"
)

//counterfeiter:generate . IGit
type IGit interface {
	ICommander
	IGitConfigMgr

	Branches() IBranchesMgr
	Commits() ICommitsMgr
	Worktree() IWorktreeMgr
	Submodules() ISubmodulesMgr
	Status() IStatusMgr
	Stash() IStashMgr
	Tags() ITagsMgr
	Remotes() IRemotesMgr
	Reflog() IReflogMgr
	Sync() ISyncMgr
	Flow() IFlowMgr
	Rebasing() IRebasingMgr

	// diffing
	WorktreeFileDiff(file *models.File, plain bool, cached bool) string
	WorktreeFileDiffCmdObj(node models.IFile, plain bool, cached bool) ICmdObj
	ShowFileDiff(from string, to string, reverse bool, fileName string, plain bool) (string, error)
	ShowFileDiffCmdObj(from string, to string, reverse bool, path string, plain bool, showRenames bool) ICmdObj
	DiffEndArgs(from string, to string, reverse bool, path string) string
	GetFilesInDiff(from string, to string, reverse bool) ([]*models.CommitFile, error)

	// common
	GetLog() *logrus.Entry
	WithSpan(span string) IGit
	GetOS() oscommands.IOS

	// patch
	NewPatchManager() *patch.PatchManager
	DeletePatchesFromCommit(commits []*models.Commit, commitIndex int, p *patch.PatchManager) error
	MovePatchToSelectedCommit(commits []*models.Commit, sourceCommitIdx int, destinationCommitIdx int, p *patch.PatchManager) error
	MovePatchIntoIndex(commits []*models.Commit, commitIdx int, p *patch.PatchManager, stash bool) error
	PullPatchIntoNewCommit(commits []*models.Commit, commitIdx int, p *patch.PatchManager) error
}
