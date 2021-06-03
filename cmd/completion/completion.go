/*
Copyright 2020-2021 The OpenEBS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package completion

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/klog"
)

// NewCmdCompletion creates the completion command
func NewCmdCompletion(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:       "completion SHELL",
		ValidArgs: []string{"bash", "zsh"},
		Args:      cobra.ExactValidArgs(1),
		Short:     "Outputs shell completion code for the specified shell (bash or zsh)",
		Long: `
Outputs shell completion code for the specified shell (bash or zsh)

To load completion to current bash shell,
. <(openebs completion bash)

To configure your bash shell to load completions for each session add to your bashrc
# ~/.bashrc or ~/.profile
. <(kubectl openebs completion bash)

To load completion to current zsh shell,
. <(kubectl openebs completion zsh)

To configure your zsh shell to load completions for each session add to your zshrc
# ~/.zshrc
. <(kubectl openebs completion zsh)
		`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			RunCompletion(os.Stdout, rootCmd, args)
		},
	}

	return cmd
}

// RunCompletion is used to run the completion of the cobra commad
func RunCompletion(out io.Writer, cmd *cobra.Command, args []string) {
	if args[0] == "bash" {
		RunCompletionBash(out, cmd)
		return
	}
	if args[0] == "zsh" {
		RunCompletionZsh(out, cmd)
		return
	}
	fmt.Printf("Unsupported shell type %q.\n", args[0])
}

// RunCompletionBash is used for the bash shell
func RunCompletionBash(out io.Writer, cmd *cobra.Command) {
	err := cmd.GenBashCompletion(out)
	if err != nil {
		klog.Error(err)
	}
}

// RunCompletionZsh is used for the zsh shell
func RunCompletionZsh(out io.Writer, cmd *cobra.Command) {
	// TODO: Why not just use cmd.GenZshCompletion(out)
	zshInitialization := `
__openebs_bash_source() {
	alias shopt=':'
	alias _expand=_bash_expand
	alias _complete=_bash_comp
	emulate -L sh
	setopt kshglob noshglob braceexpand

	source "$@"
}

__openebs_type() {
	# -t is not supported by zsh
	if [ "$1" == "-t" ]; then
		shift

		# fake Bash 4 to disable "complete -o nospace". Instead
		# "compopt +-o nospace" is used in the code to toggle trailing
		# spaces. We don't support that, but leave trailing spaces on
		# all the time
		if [ "$1" = "__openebs_compopt" ]; then
			echo builtin
			return 0
		fi
	fi
	type "$@"
}

__openebs_compgen() {
	local completions w
	completions=( $(compgen "$@") ) || return $?

	# filter by given word as prefix
	while [[ "$1" = -* && "$1" != -- ]]; do
		shift
		shift
	done
	if [[ "$1" == -- ]]; then
		shift
	fi
	for w in "${completions[@]}"; do
		if [[ "${w}" = "$1"* ]]; then
			echo "${w}"
		fi
	done
}

__openebs_compopt() {
	true # don't do anything. Not supported by bashcompinit in zsh
}

__openebs_ltrim_colon_completions()
{
	if [[ "$1" == *:* && "$COMP_WORDBREAKS" == *:* ]]; then
		# Remove colon-word prefix from COMPREPLY items
		local colon_word=${1%${1##*:}}
		local i=${#COMPREPLY[*]}
		while [[ $((--i)) -ge 0 ]]; do
			COMPREPLY[$i]=${COMPREPLY[$i]#"$colon_word"}
		done
	fi
}

__openebs_get_comp_words_by_ref() {
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[${COMP_CWORD}-1]}"
	words=("${COMP_WORDS[@]}")
	cword=("${COMP_CWORD[@]}")
}

autoload -U +X bashcompinit && bashcompinit

# use word boundary patterns for BSD or GNU sed
LWORD='[[:<:]]'
RWORD='[[:>:]]'
if sed --help 2>&1 | grep -q GNU; then
	LWORD='\<'
	RWORD='\>'
fi

__openebs_convert_bash_to_zsh() {
	sed \
	-e 's/declare -F/whence -w/' \
	-e 's/_get_comp_words_by_ref "\$@"/_get_comp_words_by_ref "\$*"/' \
	-e 's/local \([a-zA-Z0-9_]*\)=/local \1; \1=/' \
	-e 's/flags+=("\(--.*\)=")/flags+=("\1"); two_word_flags+=("\1")/' \
	-e 's/must_have_one_flag+=("\(--.*\)=")/must_have_one_flag+=("\1")/' \
	-e "s/${LWORD}_filedir${RWORD}/__openebs_filedir/g" \
	-e "s/${LWORD}_get_comp_words_by_ref${RWORD}/__openebs_get_comp_words_by_ref/g" \
	-e "s/${LWORD}__ltrim_colon_completions${RWORD}/__openebs_ltrim_colon_completions/g" \
	-e "s/${LWORD}compgen${RWORD}/__openebs_compgen/g" \
	-e "s/${LWORD}compopt${RWORD}/__openebs_compopt/g" \
	-e "s/${LWORD}declare${RWORD}/builtin declare/g" \
	-e "s/\\\$(type${RWORD}/\$(__openebs_type/g" \
	<<'BASH_COMPLETION_EOF'
`
	out.Write([]byte(zshInitialization))

	buf := new(bytes.Buffer)
	err := cmd.GenBashCompletion(buf)
	if err != nil {
		klog.Error(err)
	}
	out.Write(buf.Bytes())

	zshTail := `
BASH_COMPLETION_EOF
}

__openebs_bash_source <(__openebs_convert_bash_to_zsh)
_complete openebs 2>/dev/null
`
	out.Write([]byte(zshTail))
}
