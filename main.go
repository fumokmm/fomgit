package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
	"bufio"
	"strconv"
)

func getAnswer(message string) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(message + " (Y/N)")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)  // Remove the newline at the end
	
	return answer == "Y" || answer == "y"
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	return strings.TrimSpace(answer)  // Remove the newline at the end
}

func getCurrentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}

func getBranchList() []string {
	cmd := exec.Command("git", "branch", "--format", "%(refname:short)")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	branches := []string{}
	for _, branch := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		branches = append(branches, branch)
	}
	return branches
}

func branchExists(branchName string) bool {
	branches := getBranchList()
	for _, branch := range branches {
		if branch == branchName {
			return true
		}
	}
	return false
}

func mainMerge(currentBranch string) {
	mergeBranch := ""

	// mainブランチにいる場合
	if currentBranch == "main" {
		branches := getBranchList()
		// mainブランチは省く
		branchesWithoutMain := []string{}
		for _, branch := range branches {
			if branch != "main" {
				branchesWithoutMain = append(branchesWithoutMain, branch)
			}
		}

		if len(branchesWithoutMain) == 0 {
			fmt.Println("There are no branch to merge.")
			os.Exit(0)
		}

		fmt.Println("Select a branch to merge into main:")
		for i, branch := range branchesWithoutMain {
			fmt.Printf("%d: %s\n", i+1, branch)
		}

		var selection = getInput()
		num, err := strconv.Atoi(selection)
		// 数値に変換できなかった場合、ブランチ名でチェックする
		if err != nil {
			for _, branch := range branchesWithoutMain {
				if branch == selection {
					mergeBranch = selection
					break
				}
			}
			// 一致するものがなければエラー
			if mergeBranch == "" {
				fmt.Println("Invalid branch name.")
				os.Exit(1)
			}

		// 数値に変換できた場合
		} else {
			// 範囲外ならエラー
			if num < 1 || num > len(branchesWithoutMain) {
				fmt.Println("Invalid selection number.")
				os.Exit(1)
			}
			mergeBranch = branchesWithoutMain[num - 1]
		}

	// mainブランチにいない場合
	} else {
		var result = getAnswer(fmt.Sprintf("Are you sure you want to merge the current branch [ %s ] into main?", currentBranch))
		if !result {
			fmt.Println("Aborting.")
			os.Exit(1)
		}
		mergeBranch = currentBranch
	}

	// 現在mainブランチにいなければmainブランチをチェックアウトする
	if currentBranch != "main" {
		if err := exec.Command("git", "checkout", "main").Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// マージを実行
	if err := exec.Command("git", "merge", mergeBranch, "--no-ff").Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Merge successful.")

	if getAnswer(fmt.Sprintf("Do you want to delete the merged branch [ %s ]?", mergeBranch)) {
		//fmt.Printf("git branch -d %s\n", mergeBranch)
		if err := exec.Command("git", "branch", "-d", mergeBranch).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// 元のブランチが残っていればチェックアウトする
	if branchExists(currentBranch) {
		//fmt.Printf("git checkout %s\n", currentBranch)
		if err := exec.Command("git", "checkout", currentBranch).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}


func mainFeature() {
	// フィーチャー名を取得
	featureName := "feature"
	fmt.Printf("Enter a feature name. Default is [ %s ]. (default when omitted):\n", featureName)
	inputtedFeatureName := getInput()
    if inputtedFeatureName != "" {
        featureName = inputtedFeatureName
	}

	fmt.Printf("Enter a feature branch name:\n")
	featureBranchName := getInput()
    if featureBranchName == "" {
		fmt.Println("Please enter feature branch name.")
		os.Exit(1)
	}

    // mainブランチを起点としてフィーチャーブランチを作成し、スイッチする
	branchName := fmt.Sprintf("%s/%s", featureName, featureBranchName)
 	fmt.Printf("git switch -c %s main\n", branchName)
	if err := exec.Command("git", "switch", "-c", branchName, "main").Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

    fmt.Printf("Created feature branch [ %s ] and switched to it.\n", branchName)
}

func usage() {
	fmt.Println("Usage: merge")
	fmt.Println("       feature")
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 || (args[0] != "merge" && args[0] != "feature") {
		usage()
		os.Exit(1)
	}

	// コマンドを取得
	command := args[0]

	if command == "merge" {
		currentBranch := getCurrentBranch()
		mainMerge(currentBranch)
		os.Exit(0)

	} else if command == "feature" {
		mainFeature()
		os.Exit(0)

	} else {
		usage()
		os.Exit(1)
	}
}
